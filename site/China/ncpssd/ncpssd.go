package ncpssd

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"encoding/json"
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
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {

	dUrl := ""
	if strings.Contains(dt.Url, "fullTextRead?filePath=") {
		dUrl = getPdfUrl(dt.Url)
		dt.BookId = getBookId(dUrl)
	} else {
		dt.BookId = getBookId(dt.Url)
		if dt.BookId == "" {
			return "requested URL was not found.", err
		}
		name := util.GenNumberSorted(dt.Index)
		log.Printf("Get %s  %s\n", name, dt.Url)
		dUrl, _ = getReadUrl(dt.BookId, dt.Url)
	}
	if dUrl == "" {
		return "requested URL was not found.", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, "")
	token, _ := getToken(dt.Url)
	ext := util.FileExt(dUrl)
	//sortId := util.GenNumberSorted(1) + "_" + dt.BookId
	log.Printf("Get %s  %s\n", dt.BookId, dUrl)
	fileName := dt.BookId + ext
	dest := config.GetDestPath(dt.Url, "", fileName)
	jar, _ := cookiejar.New(nil)
	gohttp.FastGet(dUrl, gohttp.Options{
		DestFile:    dest,
		Overwrite:   false,
		Concurrency: 1,
		CookieJar:   jar,
		CookieFile:  config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"user-agent": config.UserAgent,
			"Referer":    "https://www.ncpssd.org/",
			"Origin":     "https://www.ncpssd.org/",
			"site":       "npssd",
			"sign":       token,
		},
	})
	return "", err
}

func getBookId(text string) string {
	var bookId string
	m := regexp.MustCompile(`(?i)barcodenum=([A-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		return m[1]
	}
	m = regexp.MustCompile(`(?i)pdf/([A-z0-9_-]+)\.pdf`).FindStringSubmatch(text)
	if m != nil {
		return m[1]
	}
	return bookId
}

func getReadUrl(bookId string, sUrl string) (string, error) {
	apiUrl := fmt.Sprintf("https://www.ncpssd.org/Literature/readurl?id=%s&type=3", bookId)
	bs, err := getBody(apiUrl, sUrl)
	if err != nil {
		return "", err
	}
	var respReadUrl ResponseReadUrl
	if err = json.Unmarshal(bs, &respReadUrl); err != nil {
		return "", err
	}
	return respReadUrl.Url, nil
}

func getPdfUrl(sUrl string) string {
	var pdfUrl string
	m := regexp.MustCompile(`(?i)filePath=(.+)\.pdf`).FindStringSubmatch(sUrl)
	if m != nil {
		s, _ := url.QueryUnescape(m[1])
		pdfUrl = s + ".pdf"
	}
	return pdfUrl
}
