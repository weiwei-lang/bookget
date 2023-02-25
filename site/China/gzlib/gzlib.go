package gzlib

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}

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
		return "", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	header, _ := curl.GetHeaderFile(config.Conf.CookieFile)
	pdfUrl, err := getPdfUrl(dt, config.Conf.CookieFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Printf("Get %d  %s\n", dt.Index, pdfUrl)
	ext := util.FileExt(pdfUrl)
	dest := config.GetDestPath(dt.Url, dt.BookId, dt.BookId+ext)

	fi, err := os.Stat(dest)
	if err == nil && fi.Size() > 0 {
		return
	}
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		DestFile:   dest,
		Headers: map[string]interface{}{
			"User-Agent":     "ReaderEx 2.3",
			"Accept-Range":   "bytes=0-",
			"Range":          "bytes=0-",
			"Request-Cookie": header["Cookie"],
		},
	})
	_, err = cli.Get(pdfUrl)
	if err != nil {
		return "", err
	}
	return "ok", nil
}

func getBookId(text string) string {
	sUrl := strings.ToLower(text)
	bookId := ""
	m := regexp.MustCompile(`bookid=([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	m = regexp.MustCompile(`filename=([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func getPdfUrl(dt *DownloadTask, cookieFile string) (string, error) {
	jar, _ := cookiejar.New(nil)
	apiUrl := fmt.Sprintf("http://%s/Hrcanton/Search/ResultDetail?BookId=%s", dt.UrlParsed.Host, dt.BookId)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return "", err
	}
	bs, _ := resp.GetBody()
	text := string(bs)
	pdfUrl := ""
	//var fileUrl = "http://113.108.173.156" + subStr;
	m := regexp.MustCompile(`fileUrl[\s]+=[\s]+["'](\S+)["']`).FindStringSubmatch(text)
	if m != nil {
		pdfUrl = m[1]
	}
	//var subStr = "/OnlineViewServer/onlineview.aspx?filename=GZDD034601001.pdf"
	m = regexp.MustCompile(`subStr[\s]+=[\s]+["'](\S+)["']`).FindStringSubmatch(text)
	if m != nil {
		pdfUrl += m[1]
	}
	if pdfUrl == "" {
		pdfUrl = fmt.Sprintf("http://113.108.173.156/OnlineViewServer/onlineview.aspx?filename=%s.pdf", dt.BookId)
	}
	return pdfUrl, nil
}
