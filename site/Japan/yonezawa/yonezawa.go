package yonezawa

import (
	"bookget/config"
	curl "bookget/lib/curl"
	util "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	taskName := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", taskName, taskUrl)

	bookId := getBookId(taskUrl)
	if bookId == "" {
		return
	}
	config.CreateDirectory(taskUrl, bookId)
	StartDownload(taskUrl, bookId)
	return "", err
}

func getBookId(taskUrl string) string {
	m := regexp.MustCompile(`/([A-z\d_-]+)_view.html`).FindStringSubmatch(taskUrl)
	if m != nil {
		return m[1]
	}
	return ""
}

func StartDownload(pageUrl, bookId string) {
	canvases := getCanvases(pageUrl)
	if canvases.Size == 0 {
		return
	}
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
		dest := config.GetDestPath(pageUrl, bookId, fileName)
		curl.FastGet(uri, dest, nil, true)
	}
}

func getCanvases(uri string) (canvases Canvases) {
	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	text := string(bs)
	matches := regexp.MustCompile(`<option\s+value=["']?([A-z\d,_-]+)["']?`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	//var dir='data/AA003/';
	imageDir := regexp.MustCompile(`var\s+dir\s?=\s?["'](\S+)["']`).FindStringSubmatch(text)
	if imageDir == nil {
		return
	}
	pos := strings.LastIndex(uri, "/")
	if pos == -1 {
		return
	}
	host := uri[:pos+1]

	for _, val := range matches {
		imgUrls := getImageUrls(host, imageDir[1], val[1])
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrls...)
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}

func getImageUrls(host, imageDir, val string) (imgUrls []string) {
	m := strings.Split(val, ",")
	if m == nil {
		return
	}
	id := m[0]
	max, _ := strconv.Atoi(m[1])
	imgUrls = make([]string, 0, max)
	for i := 1; i <= max; i++ {
		imgUrl := host + makeUri(imageDir, id, i)
		imgUrls = append(imgUrls, imgUrl)
	}
	return
}

func makeUri(imageDir, val string, i int) string {
	dir2 := val[5:8]
	book := val[0:8]
	page := val[len(val)-3:]
	page = regexp.MustCompile(`^0+0?`).ReplaceAllString(page, "")
	sortId := util.GenNumberLimitLen(i, 3)
	s := fmt.Sprintf("%s%s/%s_%s.jpg", imageDir, dir2, book, sortId)
	return s
}
