package familysearch

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"bookget/lib/zhash"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

var API_URL = "https://www.familysearch.org/search/filmdata/filmdatainfo"

var Cookies []*http.Cookie
var CookieJar *cookiejar.Jar

func Init(iTask int, sUrl string) (msg string, err error) {
	CookieJar, _ = cookiejar.New(nil)
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	dt.BookId = getBookId(dt)
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)
	if dt.UrlType == 1 {
		return ImagesDownload(dt)
	}
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {
	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	imageData, err := getImageData(dt.Url, config.Conf.CookieFile)
	if err != nil {
		return "", err
	}
	filmData, err := getFilmData(dt.Url, imageData.DgsNum, config.Conf.CookieFile)
	if err != nil {
		return "", err
	}
	size := len(filmData.Images)
	log.Printf(" %d Pages.\n", size)

	//{id} = 3:1:3QSQ-G9SM-C8SC
	//{image} = image.xml 或 dist.jpg?proxy=true
	iifDown(dt, filmData, config.Conf.CookieFile)
	return "", nil
}

func dasDown(dt *DownloadTask, filmData ResultFilmData) {
	jar, _ := cookiejar.New(nil)
	//filmData.Templates.DasTemplate
	//https://sg30p0.familysearch.org/service/records/storage/dascloud/das/v2/{id}/{image}
	dasTemplate := regexp.MustCompile(`\{[A-z]+\}`).ReplaceAllString(filmData.Templates.DasTemplate, "%s")
	for index, image := range filmData.Images {
		// https://familysearch.org/ark:/61903/3:1:3QSQ-G9SM-C8SC/image.xml
		m := regexp.MustCompile(`/([^/]+)/image.xml`).FindStringSubmatch(image)
		if m == nil {
			continue
		}
		id := m[1]
		dUrl := fmt.Sprintf(dasTemplate, id, "dist.jpg?proxy=true")
		sortId := util.GenNumberSorted(index + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		fileName := sortId + ".jpg"
		dest := config.GetDestPath(dt.Url, dt.BookId, fileName)
		for {
			_, err := gohttp.FastGet(dUrl, gohttp.Options{
				DestFile:    dest,
				Overwrite:   false,
				Concurrency: config.Conf.Threads,
				CookieJar:   jar,
				CookieFile:  config.Conf.CookieFile,
				Headers: map[string]interface{}{
					"user-agent": config.UserAgent,
				},
			})
			if err != nil {
				fmt.Println(err)
				util.PrintSleepTime(60)
				continue
			}
			break
		}
		util.PrintSleepTime(config.Conf.Speed)
	}
}

func iifDown(dt *DownloadTask, filmData ResultFilmData, cookieFile string) {
	//filmData.Templates.DzTemplate
	//https://sg30p0.familysearch.org/service/records/storage/deepzoomcloud/dz/v1/{id}/{image}
	dzTpl := regexp.MustCompile(`\{[A-z]+\}`).ReplaceAllString(filmData.Templates.DzTemplate, "%s")
	header, _ := curl.GetHeaderFile(cookieFile)
	args := []string{"--dezoomer=deepzoom",
		"-H", "authority:www.familysearch.org",
		"-H", "referer:" + url.QueryEscape(dt.Url),
		"-H", "User-Agent:" + header["User-Agent"],
		"-H", "cookie:" + header["Cookie"],
	}
	storePath := dt.SavePath + string(os.PathSeparator)
	for i, image := range filmData.Images {
		if config.SeqContinue(i) {
			continue
		}
		m := regexp.MustCompile(`/([^/]+)/image.xml`).FindStringSubmatch(image)
		if m == nil {
			continue
		}
		id := m[1]
		sortId := util.GenNumberSorted(i + 1)
		inputUri := fmt.Sprintf(dzTpl, id, "image.xml")
		log.Printf("Get %s  %s\n", sortId, inputUri)
		outfile := storePath + sortId + config.Conf.FileExt
		util.StartProcess(inputUri, outfile, args)
		util.PrintSleepTime(config.Conf.Speed)
	}
	return
}

func getBookId(dt *DownloadTask) string {
	bookId := ""
	m := regexp.MustCompile(`(?i)wc=([^&]+)`).FindStringSubmatch(dt.Url)
	if m != nil {
		bookId = strconv.FormatUint(uint64(zhash.CRC32(m[1])), 10)
		dt.UrlType = 0 //中國族譜收藏 1239-2014年 https://www.familysearch.org/search/collection/1787988
	}
	m = regexp.MustCompile(`(?i)rmsId=([A-z\d-_]+)`).FindStringSubmatch(dt.Url)
	if m != nil {
		bookId = m[1]
		dt.UrlType = 1 //家谱图像 https://www.familysearch.org/records/images/
	}
	m = regexp.MustCompile(`(?i)groupId=([A-z\d-_]+)`).FindStringSubmatch(dt.Url)
	if m != nil {
		bookId = m[1]
		dt.UrlType = 1 //家谱图像 https://www.familysearch.org/ark:/61903/3:1:3QS7-L9S9-WS92?view=explore&groupId=M94X-6HR
	}
	return bookId
}

func getImageData(sUrl, cookieFile string) (result ResultImageData, err error) {
	u, err := url.Parse(sUrl)
	if err != nil {
		return
	}
	q := u.Query()
	var d = ImageData{}
	d.Type = "image-data"
	d.Args.ImageURL = sUrl
	d.Args.Locale = "zh"
	d.Args.State.Wc = q.Get("wc")
	d.Args.State.Cc = q.Get("cc")
	d.Args.State.ImageOrFilmUrl = u.Path
	d.Args.State.CollectionContext = q.Get("cc")
	d.Args.State.ViewMode = "i"
	d.Args.State.SelectedImageIndex = -1
	d.Args.State.WaypointContext = "/service/cds/recapi/waypoints/" + q.Get("wc")

	//post json
	cli := gohttp.NewClient()
	resp, err := cli.Post(API_URL, gohttp.Options{
		CookieFile: cookieFile,
		CookieJar:  CookieJar,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
			"accept":       "application/json",
		},
		JSON: d,
	})
	if err != nil {
		return
	}
	body, err := resp.GetBody()
	if err != nil {
		return
	}
	var resultError ResultError
	if err = json.Unmarshal(body, &resultError); resultError.Error.StatusCode != 0 {
		msg := fmt.Sprintf("StatusCode: %d, Message: %s", resultError.Error.StatusCode, resultError.Error.Message)
		err = errors.New(msg)
		return
	}

	cookieURL, _ := url.Parse(API_URL)
	Cookies = CookieJar.Cookies(cookieURL)
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}

func getFilmData(sUrl, dgsNum, cookieFile string) (result ResultFilmData, err error) {
	u, err := url.Parse(sUrl)
	if err != nil {
		return
	}
	var d = FilmData{}
	d.Type = "film-data"
	d.Args.WaypointURL = ""
	d.Args.DgsNum = dgsNum
	d.Args.State.ImageOrFilmUrl = u.Path
	d.Args.State.ViewMode = "i"
	d.Args.State.SelectedImageIndex = -1
	d.Args.Locale = "zh"

	//post json
	cli := gohttp.NewClient()
	resp, err := cli.Post(API_URL, gohttp.Options{
		CookieFile: cookieFile,
		CookieJar:  CookieJar,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
			"accept":       "application/json",
		},
		JSON: d,
	})
	if err != nil {
		return
	}
	body, err := resp.GetBody()
	if err != nil {
		return
	}
	cookieURL, _ := url.Parse(API_URL)
	Cookies = CookieJar.Cookies(cookieURL)
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
