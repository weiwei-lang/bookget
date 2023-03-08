package jslib

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
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
	apiServer := getApiServer(dt.BookId, dt.UrlParsed)
	tiles, err := getCanvases(dt.BookId, config.Conf.CookieFile, apiServer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := `{
    "Image": {
    "xmlns":    "http://schemas.microsoft.com/deepzoom/2009",
    "Url":      "%s",
    "Format":   "%s",
    "Overlap":  "1", 
	"MaxLevel": "0",
	"Separator": "/",
        "TileSize": "%d",
        "Size": {
            "Height": "%d",
            "Width":  "%d"
        }
    }
}
`
	dziUrls := make([]string, 0, len(tiles))
	ext := ""
	for key, item := range tiles {
		k := regexp.MustCompile(`(\d+)`).FindString(key)
		i, _ := strconv.Atoi(k)
		sortId := fmt.Sprintf("%s.json", util.GenNumberSorted(i))
		dest := config.GetDestPath(dt.Url, dt.BookId, sortId)
		serverUrl := fmt.Sprintf("%s/tiles/%s/", apiServer, key)
		if ext == "" {
			ext = "." + strings.ToLower(item.Extension)
		}
		txt := fmt.Sprintf(text, serverUrl, item.Extension, item.TileSize.W, item.Height, item.Width)
		util.FileWrite([]byte(txt), dest)
		dziUrls = append(dziUrls, sortId)
	}
	sort.Sort(strs(dziUrls))
	storePath := dt.SavePath + string(os.PathSeparator)
	args := []string{"--dezoomer=deepzoom",
		"-H", "Origin:" + url.QueryEscape(dt.Url),
		"-H", "Referer:" + url.QueryEscape(dt.Url),
		"-H", "User-Agent:" + config.Conf.UserAgent,
	}
	for i, val := range dziUrls {
		if config.SeqContinue(i) {
			continue
		}
		inputUri := storePath + val
		outfile := storePath + util.GenNumberSorted(i+1) + ext
		if ret := util.StartProcess(inputUri, outfile, args); ret == true {
			os.Remove(inputUri)
		}
		util.PrintSleepTime(config.Conf.Speed)
	}
	return "", nil
}

func getApiServer(bookId string, u *url.URL) string {
	apiServer := fmt.Sprintf("%s://%s/medias/%s", u.Scheme, u.Host, bookId)
	return apiServer
}

func getBookId(text string) string {
	bookId := ""
	m := regexp.MustCompile(`bookid=([A-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		return m[1]
	}
	m = regexp.MustCompile(`id=([A-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func getCanvases(bookId string, cookieFile string, apiServer string) (tiles map[string]Item, err error) {
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	apiUrl := fmt.Sprintf("%s/tiles/infos.json", apiServer)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
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
		return
	}
	var result ResponseBody
	if err = json.Unmarshal(bs, &result); err != nil {
		return
	}
	if result.Tiles == nil {
		return
	}
	return result.Tiles, nil
}
