package niiac

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
	m := regexp.MustCompile(`toyobunko/([^/]+)/([^/]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = fmt.Sprintf("%s.%s", m[1], m[2])
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	imageUrls, iiifUrls := getImageUrls(bookId, taskUrl)
	if imageUrls == nil || iiifUrls == nil {
		return
	}
	size := len(imageUrls)
	log.Printf(" %d pages.\n", size)

	destPath := config.CreateDirectory(taskUrl, bookId)
	util.CreateShell(destPath, iiifUrls, nil)
	for i, uri := range imageUrls {
		if uri == "" {
			continue
		}
		ext := util.FileExt(uri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		curl.FastGet(uri, dest, nil, true)
	}
}

func getImageUrls(bookId, bookUrl string) (imgUrls []string, iiifUrls []string) {
	uri := fmt.Sprintf("%s/manifest.json", bookUrl)
	var manifest = new(Manifest)
	bs, err := curl.Get(uri, nil)
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
	imgUri := make([]string, 0, i)
	iiifUri := make([]string, 0, i)
	newWidth := ""
	//>6400使用原图
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "/full/full/0/default.jpg"
	} else if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("/full/%d,/0/default.jpg", config.Conf.FullImageWidth)
	}
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//dezoomify-rs URL
			iiiInfo := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
			iiifUri = append(iiifUri, iiiInfo)

			//JPEG URL
			imgUrl := ""
			if newWidth == "" {
				imgUrl = image.Resource.Id
			} else {
				imgUrl = fmt.Sprintf("%s%s", image.Resource.Service.Id, newWidth)
			}
			imgUri = append(imgUri, imgUrl)
		}
	}
	return imgUri, iiifUri
}
