package usthk

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
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

func getBookId(text string) string {
	text = strings.ToLower(text)
	bookId := ""
	m := regexp.MustCompile(`bib/([A-z0-9_-]+)`).FindStringSubmatch(text)
	if m != nil {
		bookId = m[1]
	}
	return bookId
}

func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	canvases := getCanvases(dt)
	log.Printf(" %d pages.\n", canvases.Size)
	for i, uri := range canvases.ImgUrls {
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
		curl.FastGet(uri, dest, nil, true)
	}

	return
}

func getCanvases(dt *DownloadTask) (canvases Canvases) {
	bs, err := curl.Get(dt.Url, nil)
	if err != nil {
		return
	}
	text := string(bs)

	//view_book('6/o/b1129168/ebook'
	matches := regexp.MustCompile(`view_book\(["'](\S+)["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	imgUrls := make([]string, 0, 1000)
	for _, m := range matches {
		sPath := m[1]
		uri := fmt.Sprintf("https://%s/bookreader/getfilelist.php?path=%s", dt.UrlParsed.Host, sPath)
		bs, err = curl.Get(uri, nil)
		if err != nil {
			return
		}
		respFiles := new(ResponseFiles)
		if err = json.Unmarshal(bs, respFiles); err != nil {
			log.Printf("json.Unmarshal failed: %s\n", err)
			return
		}
		//imgUrls := make([]string, 0, len(result.FileList))
		for _, v := range respFiles.FileList {
			imgUrl := fmt.Sprintf("https://%s/obj/%s/%s", dt.UrlParsed.Host, sPath, v)
			imgUrls = append(imgUrls, imgUrl)
		}
	}
	canvases.ImgUrls = imgUrls
	canvases.Size = len(imgUrls)
	return
}
