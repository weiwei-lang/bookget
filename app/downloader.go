package app

import (
	"net/http/cookiejar"
	"net/url"
)

type Downloader interface {
	Init(iTask int, sUrl string) (msg string, err error)
	getBookId(sUrl string) (bookId string)
	download() (msg string, err error)
	do(imgUrls []string) (msg string, err error)
	getVolumes(sUrl string, jar *cookiejar.Jar) (volumes []string, err error)
	getCanvases(sUrl string, jar *cookiejar.Jar) (canvases []string, err error)
}

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
	Title     string
	Param     map[string]interface{} //备用参数
	Jar       *cookiejar.Jar
}
