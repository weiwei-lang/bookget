package gprime

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
	"strconv"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	dt.Jar, _ = cookiejar.New(nil)
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)
	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s %s %s\n", name, dt.Title, dt.Url)
	canvases := getCanvases(dt.BookId, dt.Jar)
	log.Printf(" %d pages \n", len(canvases))
	if canvases == nil || len(canvases) == 0 {
		log.Printf("requested URL was not found.\n")
	}
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
		dest := config.GetDestPath(dt.Url, dt.BookId, fileName)
		gohttp.FastGet(uri, gohttp.Options{
			Concurrency: 1,
			DestFile:    dest,
			Overwrite:   false,
			CookieJar:   dt.Jar,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
	return "", nil
}

func getBookId(bookUrl string) (bookId string) {
	m := regexp.MustCompile(`(?i)tilcod=([A-z0-9_-]+)`).FindStringSubmatch(bookUrl)
	if m != nil {
		bookId = m[1]
	}
	return
}

func getCanvases(bookId string, jar *cookiejar.Jar) []string {
	urlTemplate := "http://e-library2.gprime.jp/lib_pref_osaka/da/download?id=%s&size=full&type=image&file=%s"
	var canvases = make([]string, 0, 1000)
	for i := 1; i < 10000; i++ {
		bs, err := getBody(bookId, i, jar)
		if err != nil || bs == nil {
			continue
		}
		var resImage ResponseImage
		if err := json.Unmarshal(bs, &resImage); err != nil {
			continue
		}
		for _, v := range resImage.ImagePath {
			sUrl := fmt.Sprintf(urlTemplate, bookId, v)
			canvases = append(canvases, sUrl)
		}
		if !resImage.IsNext || resImage.Size == 0 || resImage.Total < i {
			break
		}
	}
	return canvases
}

func getBody(bookId string, page int, jar *cookiejar.Jar) ([]byte, error) {
	apiUrl := "http://e-library2.gprime.jp/lib_pref_osaka/da/ajax/image"
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent":       config.Conf.UserAgent,
			"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With": "XMLHttpRequest",
			"Origin":           "http://e-library2.gprime.jp",
		},
		FormParams: map[string]interface{}{
			"tilcod": bookId,
			"start":  "0",
			"page":   strconv.Itoa(page),
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
	return bs, nil
}
