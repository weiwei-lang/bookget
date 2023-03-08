package ndl

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
	m := regexp.MustCompile(`/pid/([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(num int, uri, bookId string) {
	name := util.GenNumberSorted(num)
	log.Printf("Get %s  %s\n", name, uri)

	pages := getPages(bookId)
	log.Printf(" %d pages.\n", len(pages))

	for i, imgUri := range pages {
		if config.SeqContinue(i) {
			continue
		}
		if imgUri == "" {
			continue
		}
		ext := util.FileExt(imgUri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, imgUri)
		fileName := sortId + ext
		dest := config.GetDestPath(uri, bookId, fileName)
		curl.FastGet(imgUri, dest, nil, true)
	}
}

func getPages(bookId string) (pages []string) {
	var manifest = new(Manifest)
	bs, err := curl.Get(fmt.Sprintf("https://www.dl.ndl.go.jp/api/iiif/%s/manifest.json", bookId), nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	for _, sequence := range manifest.Sequences {
		for _, canvase := range sequence.Canvases {
			for _, image := range canvase.Images {
				imgUrl := image.Resource.Id
				pages = append(pages, imgUrl)
				break
			}
		}
	}
	return
}
