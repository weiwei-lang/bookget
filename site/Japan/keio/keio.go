package keio

import (
	"bookget/app"
	"bookget/config"
	curl "bookget/lib/curl"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	return Download(dt)
}

func getBookId(text string) string {
	sUrl := strings.ToLower(text)
	bookId := ""
	m := regexp.MustCompile(`id=([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "", err
	}

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)
	volumes := getBookVolumes(dt)
	log.Printf(" %d volumes(parts).\n", len(volumes))

	for k, vol := range volumes {
		if config.Conf.Volume > 0 && config.Conf.Volume != k+1 {
			continue
		}
		canvases, err := getCanvases(vol)
		log.Printf(" %d/%d volume, %d pages \n", k+1, len(volumes), canvases.Size)
		if err != nil {
			continue
		}
		volPath := fmt.Sprintf("%s_volume%d", dt.BookId, k+1)
		config.CreateDirectory(vol, volPath)
		if config.Conf.UseDziRs {
			app.DziDownload(vol, volPath, canvases.IiifUrls)
		} else {
			app.NormalDownload(vol, volPath, canvases.ImgUrls, nil)
		}
	}
	return "", err
}

func getBookVolumes(dt *DownloadTask) (volumeUrls []string) {
	bs, err := curl.Get(dt.Url, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`<p[^>]+data-cid=['|"]([a-zA-Z0-9]+)['|"]`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	size := len(matches)
	volumeUrls = make([]string, 0, size)
	for _, v := range matches {
		childId := makeId(v[1], dt.BookId, size)
		fmt.Sprintf("%s\n", childId)
		uri := fmt.Sprintf("https://db2.sido.keio.ac.jp/iiif/manifests/kanseki/%s/%s/manifest.json", dt.BookId, childId)
		volumeUrls = append(volumeUrls, uri)
	}
	return volumeUrls
}

func makeId(childId string, bookId string, iMax int) string {
	childIDfmt := ""
	//i, _ := strconv.Atoi(childId)
	iLen := 3
	if iMax > 999 {
		iLen = 4
	}
	for k := iLen - len(childId); k > 0; k-- {
		childIDfmt += "0"
	}
	childIDfmt += childId
	return bookId + "-" + childIDfmt
}

func getCanvases(uri string) (canvases Canvases, err error) {
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(uri)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	//delete '?'
	if bs[0] != 123 {
		for i := 0; i < len(bs); i++ {
			if bs[i] == 123 {
				bs = bs[i:]
				break
			}
		}
	}
	return parseXml(bs)
}

func parseXml(bs []byte) (canvases Canvases, err error) {
	var manifest = new(Manifest)
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,", config.Conf.FullImageWidth)
	}

	size := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, size)
	canvases.IiifUrls = make([]string, 0, size)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//iifUrl, _ := url.QueryUnescape(image.Resource.Service.Id)
			//dezoomify-rs URL
			iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			canvases.IiifUrls = append(canvases.IiifUrls, iiiInfo)

			//JPEG URL
			imgUrl := fmt.Sprintf("%s/%s/0/default.jpg", image.Resource.Service.Id, newWidth)
			canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
		}
	}
	canvases.Size = size
	return canvases, nil
}
