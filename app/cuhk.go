package app

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

type Cuhk struct {
	dt *DownloadTask
}
type ResponsePage struct {
	ImagePage []ImagePage `json:"pages"`
}

type ImagePage struct {
	Pid        string `json:"pid"`
	Page       string `json:"page"`
	Label      string `json:"label"`
	Width      string `json:"width"`
	Height     string `json:"height"`
	Dsid       string `json:"dsid"`
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
}

func (r Cuhk) Init(iTask int, sUrl string) (msg string, err error) {
	r.dt = new(DownloadTask)
	r.dt.UrlParsed, err = url.Parse(sUrl)
	r.dt.Url = sUrl
	r.dt.Index = iTask
	r.dt.BookId = r.getBookId(r.dt.Url)
	if r.dt.BookId == "" {
		return "requested URL was not found.", err
	}
	r.dt.Jar, _ = cookiejar.New(nil)
	return r.download()
}

func (r Cuhk) getBookId(sUrl string) (bookId string) {
	m := regexp.MustCompile(`item/cuhk-([A-Za-z0-9]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func (r Cuhk) download() (msg string, err error) {
	name := util.GenNumberSorted(r.dt.Index)
	log.Printf("Get %s  %s\n", name, r.dt.Url)

	respVolume, err := r.getVolumes(r.dt.Url, r.dt.Jar)
	if err != nil {
		fmt.Println(err)
		return "getVolumes", err
	}
	for i, vol := range respVolume {
		if config.Conf.Volume > 0 && config.Conf.Volume != i+1 {
			continue
		}
		vid := util.GenNumberSorted(i + 1)
		r.dt.VolumeId = r.dt.BookId + "_vol." + vid
		r.dt.SavePath = config.CreateDirectory(r.dt.Url, r.dt.VolumeId)
		canvases, err := r.getCanvases(vol, r.dt.Jar)
		if err != nil || canvases == nil {
			fmt.Println(err)
			continue
		}
		log.Printf(" %d/%d volume, %d pages \n", i+1, len(respVolume), len(canvases))
		r.do(canvases)
	}
	return "", nil
}

//func (r Cuhk) do(imgUrls []string) (msg string, err error) {
//	panic("implement me")
//}

func (r Cuhk) do(imgUrls []string) (msg string, err error) {
	if imgUrls == nil {
		return
	}
	fmt.Println()
	referer := url.QueryEscape(r.dt.Url)
	size := len(imgUrls)
	for i, uri := range imgUrls {
		if !config.PageRange(i, size) {
			continue
		}
		if uri == "" {
			continue
		}
		sortId := util.GenNumberSorted(i + 1)
		filename := sortId + config.Conf.FileExt
		dest := config.GetDestPath(r.dt.Url, r.dt.VolumeId, filename)
		if FileExist(dest) {
			continue
		}
		log.Printf("Get %d/%d page, URL: %s\n", i+1, size, uri)
		opts := gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: 1,
			CookieFile:  config.Conf.CookieFile,
			CookieJar:   r.dt.Jar,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
				"Referer":    referer,
				//"X-ISLANDORA-TOKEN": v.Token,
			},
		}
		_, err := gohttp.FastGet(uri, opts)
		if err != nil {
			fmt.Println(err)
			util.PrintSleepTime(60)
			continue
		}
		util.PrintSleepTime(config.Conf.Speed)
	}
	fmt.Println()
	return "", err
}

func (r Cuhk) getVolumes(sUrl string, jar *cookiejar.Jar) (volumes []string, err error) {
	bs, err := r.getBody(sUrl, jar)
	if err != nil {
		return
	}
	text := string(bs)
	subText := util.SubText(text, "id=\"block-islandora-compound-object-compound-navigation-select-list\"", "id=\"book-viewer\">")
	matches := regexp.MustCompile(`value=['"]([A-z\d:_-]+)['"]`).FindAllStringSubmatch(subText, -1)
	if matches == nil {
		volumes = append(volumes, sUrl)
		return
	}
	volumes = make([]string, 0, len(matches))
	for _, m := range matches {
		//value='ignore'
		if m[1] == "ignore" {
			continue
		}
		id := strings.Replace(m[1], ":", "-", 1)
		volumes = append(volumes, fmt.Sprintf("https://repository.lib.cuhk.edu.hk/sc/item/%s#page/1/mode/2up", id))
	}
	return volumes, nil
}

func (r Cuhk) getCanvases(sUrl string, jar *cookiejar.Jar) (canvases []string, err error) {
	bs, err := r.getBody(sUrl, jar)
	if err != nil {
		return
	}
	var resp ResponsePage
	matches := regexp.MustCompile(`"pages":([^]]+)]`).FindSubmatch(bs)
	if matches == nil {
		return nil, err
	}
	data := []byte("{\"pages\":" + string(matches[1]) + "]}")
	if err = json.Unmarshal(data, &resp); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
	}
	for _, page := range resp.ImagePage {
		var imgUrl string
		if config.Conf.FileExt == ".jpg" {
			imgUrl = fmt.Sprintf("https://repository.lib.cuhk.edu.hk/iiif/2/%s/full/full/0/default.jpg", page.Identifier)
		} else {
			imgUrl = fmt.Sprintf("https://repository.lib.cuhk.edu.hk/islandora/object/%s/datastream/JP2", page.Pid)
		}
		canvases = append(canvases, imgUrl)
	}
	return canvases, err
}

func (r Cuhk) getBody(apiUrl string, jar *cookiejar.Jar) ([]byte, error) {
	referer := url.QueryEscape(apiUrl)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
			"Referer":    referer,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	bs, _ := resp.GetBody()
	if resp.GetStatusCode() == 202 || bs == nil {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d, %s", resp.GetStatusCode(), resp.GetReasonPhrase()))
	}
	return bs, nil
}

func (r Cuhk) getCanvasesJPEG2000(sUrl string, jar *cookiejar.Jar) (imagePage []ImagePage) {
	bs, err := r.getBody(sUrl, jar)
	if err != nil {
		return
	}
	var resp ResponsePage
	matches := regexp.MustCompile(`"pages":([^]]+)]`).FindSubmatch(bs)
	if matches != nil {
		data := []byte("{\"pages\":" + string(matches[1]) + "]}")
		if err = json.Unmarshal(data, &resp); err != nil {
			log.Printf("json.Unmarshal failed: %s\n", err)
		}
		imagePage = make([]ImagePage, len(resp.ImagePage))
		copy(imagePage, resp.ImagePage)
	}
	return imagePage
}
