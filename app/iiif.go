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
	"os"
	"regexp"
	"strconv"
	"strings"
)

type IIIF struct {
	dt         *DownloadTask
	xmlContent []byte
	BookId     string
}

// ResponseManifest  by view-source:https://iiif.lib.harvard.edu/manifests/drs:53262215
type ResponseManifest struct {
	Label     interface{} `json:"label"`
	Sequences []struct {
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Height int    `json:"height"`
			Images []struct {
				Id       string `json:"@id"`
				Type     string `json:"@type"`
				On       string `json:"on"`
				Resource struct {
					Id      string `json:"@id"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Height  int    `json:"height"`
					Service struct {
						Id string `json:"@id"`
					} `json:"service"`
					Width int `json:"width"`
				} `json:"resource"`
			} `json:"images"`
			Label string `json:"label"`
			Width int    `json:"width"`
		} `json:"canvases"`
	} `json:"sequences"`
	Structures []struct {
		Id          string   `json:"@id"`
		Type        string   `json:"@type"`
		Label       string   `json:"label"`
		ViewingHint string   `json:"viewingHint"`
		Ranges      []string `json:"ranges,omitempty"`
		Canvases    []string `json:"canvases,omitempty"`
	} `json:"structures"`
}

func (f IIIF) Init(iTask int, sUrl string) (msg string, err error) {
	f.dt = new(DownloadTask)
	f.dt.UrlParsed, err = url.Parse(sUrl)
	f.dt.Url = sUrl
	f.dt.Index = iTask
	f.dt.Jar, _ = cookiejar.New(nil)
	f.dt.BookId = f.getBookId(f.dt.Url)
	if f.dt.BookId == "" {
		return "requested URL was not found.", err
	}
	return f.download()
}

func (f IIIF) InitWithId(iTask int, sUrl string, id string) (msg string, err error) {
	f.dt = new(DownloadTask)
	f.dt.UrlParsed, err = url.Parse(sUrl)
	f.dt.Url = sUrl
	f.dt.Index = iTask
	f.dt.Jar, _ = cookiejar.New(nil)
	f.dt.BookId = id
	return f.download()
}

func (f IIIF) getBookId(sUrl string) (bookId string) {
	m := regexp.MustCompile(`/([^/]+)/manifest.json`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	m = regexp.MustCompile(`/([^/]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func (f IIIF) download() (msg string, err error) {
	f.xmlContent, err = f.getBody(f.dt.Url, f.dt.Jar)
	if err != nil || f.xmlContent == nil {
		return "requested URL was not found.", err
	}
	canvases, err := f.getCanvases(f.dt.Url, f.dt.Jar)
	if err != nil || canvases == nil {
		return
	}
	volumes, err := f.getVolumes(f.dt.Url, f.dt.Jar)
	if err != nil {
		return "getVolumes", err
	}
	if volumes == nil {
		f.dt.SavePath = config.CreateDirectory(f.dt.Url, f.dt.BookId)
		f.do(canvases)
		return "", nil
	}

	log.Printf(" %d volumes, %d pages.\n", len(volumes), len(canvases))
	k := 0
	for i, v := range volumes {
		if config.Conf.Volume > 0 && config.Conf.Volume != i+1 {
			continue
		}
		m := strings.Split(v, "^")
		size, _ := strconv.Atoi(m[1])
		log.Printf(" %d/%d volume, %d pages \n", i+1, len(volumes), size)
		vol := util.GenNumberSorted(i + 1)
		f.dt.SavePath = config.CreateDirectory(f.dt.Url, f.dt.BookId+"_vol."+vol)
		k2 := k + size
		imgUrls := canvases[k:k2]
		if config.Conf.UseDziRs {
			f.doDezoomifyRs(imgUrls)
		} else {
			f.doNormal(imgUrls)
		}
		k += size
	}

	return "", nil
}

func (f IIIF) do(imgUrls []string) (msg string, err error) {
	if config.Conf.UseDziRs {
		f.doDezoomifyRs(imgUrls)
	} else {
		f.doNormal(imgUrls)
	}
	return "", nil
}

func (f IIIF) getVolumes(sUrl string, jar *cookiejar.Jar) (volumes []string, err error) {
	var manifest = new(ResponseManifest)
	if err = json.Unmarshal(f.xmlContent, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Structures) == 0 {
		return
	}
	for _, item := range manifest.Structures {
		if item.ViewingHint == "top" {
			continue
		}
		if item.Ranges != nil {
			volId := fmt.Sprintf("%s^%d", item.Label, len(item.Ranges))
			volumes = append(volumes, volId)
		}
	}
	return volumes, nil
}

func (f IIIF) getCanvases(sUrl string, jar *cookiejar.Jar) (canvases []string, err error) {
	var manifest = new(ResponseManifest)
	if err = json.Unmarshal(f.xmlContent, manifest); err != nil {
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
	canvases = make([]string, 0, size)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			if config.Conf.UseDziRs {
				//iifUrl, _ := url.QueryUnescape(image.Resource.Service.Id)
				//dezoomify-rs URL
				iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
				canvases = append(canvases, iiiInfo)
			} else {
				//JPEG URL
				imgUrl := fmt.Sprintf("%s/%s/0/default.jpg", image.Resource.Service.Id, newWidth)
				canvases = append(canvases, imgUrl)
			}
		}
	}
	return canvases, nil
}

func (f IIIF) getBody(sUrl string, jar *cookiejar.Jar) ([]byte, error) {
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(sUrl)
	if err != nil {
		return nil, err
	}
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return nil, err
	}
	//fix bug https://www.dh-jac.net/db1/books/results-iiif.php?f1==nar-h13-01-01&f12=1&enter=portal
	//delete '?'
	if bs[0] != 123 {
		for i := 0; i < len(bs); i++ {
			if bs[i] == 123 {
				bs = bs[i:]
				break
			}
		}
	}
	return bs, nil
}

func (f IIIF) doDezoomifyRs(iiifUrls []string) bool {
	if iiifUrls == nil {
		return false
	}
	referer := url.QueryEscape(f.dt.Url)
	args := []string{
		"-H", "Origin:" + referer,
		"-H", "Referer:" + referer,
		"-H", "User-Agent:" + config.Conf.UserAgent,
	}
	for i, uri := range iiifUrls {
		if config.SeqContinue(i) {
			continue
		}
		if uri == "" {
			continue
		}
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		filename := sortId + config.Conf.FileExt
		dest := f.dt.SavePath + string(os.PathSeparator) + filename
		util.StartProcess(uri, dest, args)
	}
	return true
}

func (f IIIF) doNormal(imgUrls []string) bool {
	if imgUrls == nil {
		return false
	}
	for i, uri := range imgUrls {
		if config.SeqContinue(i) {
			continue
		}
		if uri == "" {
			continue
		}
		ext := util.FileExt(uri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		filename := sortId + ext
		dest := f.dt.SavePath + string(os.PathSeparator) + filename
		opts := gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: 1,
			CookieFile:  config.Conf.CookieFile,
			CookieJar:   f.dt.Jar,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
			},
		}
		_, err := gohttp.FastGet(uri, opts)
		if err != nil {
			fmt.Println(err)
		}
	}
	return true
}
