package kokusho

import (
	"bookget/app"
	"bookget/config"
	"bookget/lib/gohttp"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"time"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "requested URL was not found.", err
	}
	manifestUrl := getManifestUrl(dt.BookId)
	if manifestUrl == "" {
		return "requested URL was not found.", err
	}
	//iiif.StartDownload(manifestUrl, dt.BookId)
	var iiif app.IIIF
	iiif.InitWithId(dt.Index, manifestUrl, dt.BookId)
	return "", nil
}

func getBookId(text string) string {
	var bookId string
	m := regexp.MustCompile(`(?i)biblio/([A-Za-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		return m[1]
	}
	return bookId
}

func getManifestUrl(bookId string) string {
	apiUrl := fmt.Sprintf("https://kokusho.nijl.ac.jp/api/biblioDetail/%s?t=%d", bookId, time.Now().UnixMilli())
	var resp ResponseDetail
	bs, err := getBody(apiUrl)
	if err != nil {
		return ""
	}
	if err = json.Unmarshal(bs, &resp); err != nil {
		return ""
	}
	return resp.Manifest
}

func getBody(apiUrl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
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
