package sclib

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util2 "bookget/lib/util"
	"encoding/json"
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

	name := util2.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	tiles, err := getCanvases(dt.BookId, config.Conf.CookieFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dziUrls := make([]string, 0, len(tiles))
	for key, item := range tiles {
		i, _ := strconv.Atoi(key)
		sortId := fmt.Sprintf("%s.json", util2.GenNumberSorted(i))
		dest := config.GetDestPath(dt.Url, dt.BookId, sortId)
		//文件存在，跳过
		fi, err := os.Stat(dest)
		if err == nil && fi.Size() > 0 {
			continue
		}
		serverUrl := fmt.Sprintf("http://guji.sclib.org/medias/%s/tiles/%s/", dt.BookId, key)

		text := `{
    "Image": {
        "xmlns":    "http://schemas.microsoft.com/deepzoom/2008",
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
		txt := fmt.Sprintf(text, serverUrl, item.Extension, item.TileSize.W, item.Height, item.Width)
		log.Printf("Create a new file %s \n", sortId)
		util2.FileWrite([]byte(txt), dest)
		dziUrls = append(dziUrls, sortId)
	}
	sort.Sort(strs(dziUrls))
	util2.CreateShell(dt.SavePath, dziUrls, nil)
	return "请手动运行 dezoomify-rs.urls 文件", nil
}

func getBookId(text string) string {
	sUrl := strings.ToLower(text)
	bookId := ""
	m := regexp.MustCompile(`id=([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func getCanvases(bookId string, cookieFile string) (tiles map[string]Item, err error) {
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	apiUrl := fmt.Sprintf("http://guji.sclib.org/medias/%s/tiles/infos.json", bookId)
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
	var result Result
	if err = json.Unmarshal(bs, &result); err != nil {
		return
	}
	if result.Tiles == nil {
		return
	}
	return result.Tiles, nil
}
