package loc

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`item/([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	imageUrls := getPages(bookId)
	size := len(imageUrls)
	log.Printf(" %d pages.\n", size)

	//cookie 处理
	jar, _ := cookiejar.New(nil)
	for i, dUrl := range imageUrls {
		if dUrl == "" {
			continue
		}
		ext := util.FileExt(dUrl)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		gohttp.FastGet(dUrl, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			CookieJar:   jar,
			CookieFile:  config.Conf.CookieFile,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
		util.PrintSleepTime(config.Conf.Speed)
	}
}

func getBody(apiUrl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
			"authority":  "www.loc.gov",
			"origin":     "https://www.loc.gov",
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

func getPages(bookId string) (pages []string) {
	apiUrl := fmt.Sprintf("https://www.loc.gov/item/%s/?fo=json", bookId)
	bs, err := getBody(apiUrl)
	if err != nil {
		return
	}
	var manifests = new(ManifestsJson)
	if err = json.Unmarshal(bs, manifests); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	//fmt.Println(manifests)
	newWidth := ""
	//限制图片最大宽度
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/pct:100/"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}
	//一本书有N卷
	for _, resource := range manifests.Resources {
		//每卷有P页
		for _, file := range resource.Files {
			//每页有6种下载方式
			imgUrl, ok := getImagePage(file, newWidth)
			if ok {
				pages = append(pages, imgUrl)
			}
		}
	}
	return
}

func getImagePage(fileUrls []ImageFile, newWidth string) (downloadUrl string, ok bool) {
	for _, f := range fileUrls {
		if config.Conf.FileExt == ".jpg" && f.Mimetype == "image/jpeg" {
			if strings.Contains(f.Url, "full/pct:100/") {
				if newWidth != "" && newWidth != "full/pct:100/" {
					downloadUrl = strings.Replace(f.Url, "full/pct:100/", newWidth, 1)
				} else {
					downloadUrl = f.Url
				}
				ok = true
				break
			}
		} else if f.Mimetype != "image/jpeg" {
			if !config.Conf.UseCDN {
				downloadUrl = strings.Replace(f.Url, "https://tile.loc.gov/storage-services/", "http://140.147.239.202/", 1)
			} else {
				downloadUrl = f.Url
			}
			ok = true
			break
		}
	}
	return
}
