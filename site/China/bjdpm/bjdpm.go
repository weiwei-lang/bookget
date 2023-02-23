package bjdpm

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
	"strings"
)

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
	Title     string
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

	bs, err := getBody(dt.Url)
	if err != nil {
		return "Error:", err
	}
	cipherText := getCipherText(bs)
	dt.Title = getTitle(bs)

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s %s %s\n", name, dt.Title, dt.Url)

	if cipherText == nil || len(cipherText) == 0 {
		return "not found cipherText", err
	}
	dziJson, dziFormat := getDziJson(cipherText)
	sortId := fmt.Sprintf("%s.json", dt.BookId)
	dest := config.GetDestPath(dt.Url, dt.BookId, sortId)

	log.Printf("Create a new file %s \n", sortId)
	util.FileWrite([]byte(dziJson), dest)

	config.Conf.FileExt = "." + dziFormat.Format
	dziUrls := make([]string, 0)
	dziUrls = append(dziUrls, dest)

	header := make(map[string]string, 4)
	header["Origin"] = "https://minghuaji.dpm.org.cn"
	header["Referer"] = "https://minghuaji.dpm.org.cn"
	header["User-Agent"] = config.Conf.UserAgent

	util.CreateShell(dt.SavePath, dziUrls, header)
	return "请手动运行 dezoomify-rs.urls 文件", nil
}

func getBookId(text string) string {
	bookId := ""
	m := regexp.MustCompile(`id=([A-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func getTitle(bs []byte) string {
	//<title>赵孟頫水村图卷-故宫名画记</title>
	m := regexp.MustCompile(`<title>([^<]+)</title>`).FindSubmatch(bs)
	if m == nil {
		return ""
	}
	title := regexp.MustCompile("([|/\\:+\\?]+)").ReplaceAll(m[1], nil)
	return strings.Replace(string(title), "-故宫名画记", "", -1)
}

func getCipherText(bs []byte) []byte {
	//gv.init("",...)
	m := regexp.MustCompile(`gv.init\("([^"]+)",`).FindSubmatch(bs)
	if m == nil {
		return nil
	}
	return m[1]
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
