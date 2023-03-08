package ynutcm

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

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	dt.Jar, _ = cookiejar.New(nil)
	return Download(dt)
}

func getBookId(bookUrl string) (bookId string) {
	m := regexp.MustCompile(`(?i)id=([A-Za-z0-9]+)`).FindStringSubmatch(bookUrl)
	if m != nil {
		bookId = m[1]
	}
	return
}
func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "", err
	}
	//dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)
	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s %s %s\n", name, dt.Title, dt.Url)

	bookUrls, err := getMultiplebooks(dt)
	if bookUrls == nil || len(bookUrls) == 0 {
		log.Printf("requested URL was not found.\n")
		return
	}
	log.Printf(" %d volumes(parts).\n", len(bookUrls))
	serverUrl := fmt.Sprintf("%s://%s", dt.UrlParsed.Scheme, dt.UrlParsed.Host)
	for k, page := range bookUrls {
		if config.Conf.Volume > 0 && config.Conf.Volume != k+1 {
			continue
		}
		id := fmt.Sprintf("%s_volume%s", dt.BookId, util.GenNumberSorted(k+1))
		config.CreateDirectory(page, id)
		canvases := getCanvases(serverUrl, page, dt.Jar)
		log.Printf(" %d/%d volume, %d pages \n", k+1, len(bookUrls), len(canvases))
		if canvases == nil || len(canvases) == 0 {
			continue
		}
		dt.SavePath = config.CreateDirectory(dt.Url, id)
		for i, uri := range canvases {
			if config.SeqContinue(i) {
				continue
			}
			if uri == "" {
				continue
			}
			ext := util.FileExt(uri)
			sortId := util.GenNumberSorted(i + 1)
			log.Printf("Get %s  %s\n", sortId, uri)
			fileName := sortId + ext
			dest := config.GetDestPath(dt.Url, id, fileName)
			gohttp.FastGet(uri, gohttp.Options{
				Concurrency: config.Conf.Threads,
				DestFile:    dest,
				Overwrite:   false,
				Headers: map[string]interface{}{
					"user-agent": config.UserAgent,
				},
			})
			util.PrintSleepTime(config.Conf.Speed)
		}
	}
	return "", nil
}

func getCanvases(serverUrl, bookUrl string, jar *cookiejar.Jar) []string {
	bs, err := getBody(bookUrl, jar)
	if bs == nil || err != nil {
		log.Printf("requested URL was not found.")
		return nil
	}
	matches := regexp.MustCompile(`data-original=["']?([^"']+?)["']`).FindAllStringSubmatch(string(bs), -1)
	if matches == nil {
		return nil
	}
	var canvases = make([]string, 0, len(matches[1]))
	for _, v := range matches {
		imgUrl := serverUrl + strings.Replace(v[1], "\\", "/", -1)
		canvases = append(canvases, imgUrl)
	}
	return canvases
}

func getMultiplebooks(dt *DownloadTask) ([]string, error) {
	apiUrl := fmt.Sprintf("%s://%s/Yngj/Adm/Data/Detail_gj", dt.UrlParsed.Scheme, dt.UrlParsed.Host)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  dt.Jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
		FormParams: map[string]interface{}{
			"id":     dt.BookId,
			"isView": "true",
		},
	})
	resp, err := cli.Post(apiUrl)
	if err != nil {
		return nil, err
	}
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return nil, err
	}
	var respDetail ResponseDetail
	if err := json.Unmarshal(bs, &respDetail); err != nil {
		return nil, err
	}
	pageUrls := make([]string, 0, len(respDetail.Fulltextpath))
	urlTemplate := "%s://%s/Yngj/Data/PicView?number=%s&title=%s&totalnum=%s&path=%s"
	for _, v := range respDetail.Fulltextpath {
		sUrl := fmt.Sprintf(urlTemplate, dt.UrlParsed.Scheme, dt.UrlParsed.Host,
			respDetail.Detail.Number, url.QueryEscape(respDetail.Detail.Title),
			respDetail.Detail.Totalnum, url.QueryEscape(v.Tpath))
		pageUrls = append(pageUrls, sUrl)
	}
	return pageUrls, nil
}

func getBody(apiUrl string, jar *cookiejar.Jar) ([]byte, error) {
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
