package app

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"errors"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
)

type Hathitrust struct {
	dt *DownloadTask
}

func (h Hathitrust) Init(iTask int, sUrl string) (msg string, err error) {
	h.dt = new(DownloadTask)
	h.dt.UrlParsed, err = url.Parse(sUrl)
	h.dt.Url = sUrl
	h.dt.Index = iTask
	h.dt.BookId = h.getBookId(h.dt.Url)
	if h.dt.BookId == "" {
		return "requested URL was not found.", err
	}
	h.dt.Jar, _ = cookiejar.New(nil)
	return h.download()
}

func (h Hathitrust) getBookId(sUrl string) (bookId string) {
	m := regexp.MustCompile(`id=([^&]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func (h Hathitrust) download() (msg string, err error) {
	name := util.GenNumberSorted(h.dt.Index)
	log.Printf("Get %s  %s\n", name, h.dt.Url)
	canvases, err := h.getCanvases(h.dt.Url, h.dt.Jar)
	if err != nil {
		return "requested URL was not found.", err
	}
	h.dt.SavePath = config.CreateDirectory(h.dt.Url, h.dt.BookId)
	msg, err = h.do(canvases)
	return msg, err
}

func (h Hathitrust) do(imgUrls []string) (msg string, err error) {
	if imgUrls == nil {
		return
	}
	fmt.Println()
	referer := url.QueryEscape(h.dt.Url)
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
		dest := config.GetDestPath(uri, h.dt.BookId, filename)
		if FileExist(dest) {
			continue
		}
		log.Printf("Get %d/%d page, URL: %s\n", i+1, size, uri)
		opts := gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: 1,
			CookieFile:  config.Conf.CookieFile,
			CookieJar:   h.dt.Jar,
			Headers: map[string]interface{}{
				"User-Agent": config.Conf.UserAgent,
				"Referer":    referer,
			},
		}
		for true {
			_, err := gohttp.FastGet(uri, opts)
			if err != nil {
				fmt.Println(err)
				//log.Println("images (1 file per page, watermarked,  max. 20 MB / 1 min), image quality:Full")
				util.PrintSleepTime(60)
				continue
			}
			break
		}
	}
	fmt.Println()
	return "", err
}

func (h Hathitrust) getVolumes(sUrl string, jar *cookiejar.Jar) (volumes []string, err error) {
	//TODO implement me
	panic("implement me")
}

func (h Hathitrust) getCanvases(sUrl string, jar *cookiejar.Jar) (canvases []string, err error) {
	bs, err := h.getBody(h.dt.Url, h.dt.Jar)
	if err != nil || bs == nil {
		return nil, err
	}
	//取页数
	// <input id="range-seq" class="navigator-range" type="range" min="1" max="1036" value="2" aria-label="Progress" dir="rtl" />
	matches := regexp.MustCompile(`<input(?:[^>]+)id="range-seq"(?:[^>]+)max="([0-9]+)"`).FindStringSubmatch(string(bs))
	if matches == nil {
		return
	}
	size, _ := strconv.Atoi(matches[1])
	canvases = make([]string, 0, size)
	ext := config.Conf.FileExt
	format := "jpeg"
	if ext == ".png" {
		format = "png"
	} else if ext == ".tif" {
		format = "tiff"
	}
	for i := 0; i < size; i++ {
		imgurl := fmt.Sprintf("https://babel.hathitrust.org/cgi/imgsrv/image?id=%s&attachment=1&size=full&format=image/%s&seq=%d", h.dt.BookId, format, i+1)
		canvases = append(canvases, imgurl)
	}
	return canvases, err
}

func (h Hathitrust) getBody(apiUrl string, jar *cookiejar.Jar) ([]byte, error) {
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return nil, err
	}
	return bs, nil
}
