package npmtw

import (
	"bookget/config"
	curl "bookget/lib/curl"
	util "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`\?pid=(\d+)`).FindStringSubmatch(taskUrl)
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
	canvases := getImageUrls(bookId, taskUrl)
	if canvases.ImgUrls == nil {
		return
	}
	log.Printf(" %d pages.\n", canvases.Size)
	destPath := config.CreateDirectory(taskUrl, bookId)
	util.CreateShell(destPath, canvases.IiifUrls, nil)
	for i, uri := range canvases.ImgUrls {
		if uri == "" {
			continue
		}
		ext := util.FileExt(uri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		dest := config.GetDestPath(taskUrl, bookId, sortId+ext)
		curl.FastGet(uri, dest, nil, true)
	}
}

func getImageUrls(bookId, taskUrl string) (canvases Canvases) {
	var manifest = new(Manifest)
	u := fmt.Sprintf("https://digitalarchive.npm.gov.tw/Painting/setJson?pid=%s&Dept=P", bookId)
	bs, err := curl.Get(u, nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}

	i := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, i)
	canvases.IiifUrls = make([]string, 0, i)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			u := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			canvases.IiifUrls = append(canvases.IiifUrls, u)
			canvases.ImgUrls = append(canvases.ImgUrls, image.Resource.Id)
		}
	}
	canvases.Size = len(canvases.ImgUrls)
	return
}
