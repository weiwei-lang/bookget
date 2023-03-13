package emuseum

import (
	"bookget/app"
	"bookget/config"
	"bookget/lib/curl"
	util "bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`content_base_id=([A-Za-z0-9]+)&content_part_id=([A-Za-z0-9]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		if len(m[2]) < 3 {
			m[2] = "00" + m[2]
		}
		bookId = fmt.Sprintf("%s%s", m[1], m[2])
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(num int, pageUrl, bookId string) {
	name := util.GenNumberSorted(num)
	log.Printf("Get %s  %s\n", name, pageUrl)

	pages, iiifInfo := getPages(pageUrl)
	log.Printf(" %d pages.\n", len(pages))

	config.CreateDirectory(pageUrl, bookId)
	if config.Conf.UseDziRs {
		app.DziDownload(pageUrl, bookId, iiifInfo)
	} else {
		app.NormalDownload(pageUrl, bookId, pages, nil)
	}
}

func getPages(uri string) (pages []string, iiifInfo []string) {

	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	matches := regexp.MustCompile(`https://emuseum.nich.go.jp/iiifapi/([A-Za-z0-9]+)/manifest.json`).FindStringSubmatch(string(bs))
	if matches == nil {
		return
	}
	bookId := matches[1]
	var manifest = new(Manifest)
	bs, err = curl.Get(fmt.Sprintf("https://emuseum.nich.go.jp/iiifapi/%s/manifest.json", bookId), nil)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}

	i := len(manifest.Sequences[0].Canvases)
	pages = make([]string, 0, i)
	newWidth := ""
	//此站最大只支持5000
	if config.Conf.FullImageWidth > 6400 {
		newWidth = "full/full/"
	}
	if config.Conf.FullImageWidth >= 1000 {
		newWidth = fmt.Sprintf("full/%d,/", config.Conf.FullImageWidth)
	}
	for _, sequence := range manifest.Sequences {
		for _, canvase := range sequence.Canvases {
			for _, image := range canvase.Images {
				if strings.Contains(image.Resource.Service.Id, "/100001001002.tif") {
					image.Resource.Service.Id = strings.Replace(image.Resource.Service.Id, "/100001001002.tif", "/100001001001.tif", 1)
					image.Resource.Id = strings.Replace(image.Resource.Id, "/100001001002.tif", "/100001001001.tif", 1)
				}
				//dezoomify-rs URL
				iiifUrl := fmt.Sprintf("%s/info.json", image.Resource.Service.Id)
				iiifInfo = append(iiifInfo, iiifUrl)

				//JPEG URL
				imgUrl := image.Resource.Id
				if newWidth != "" {
					imgUrl = strings.Replace(image.Resource.Id, "full/full/", newWidth, 1)
				}
				pages = append(pages, imgUrl)
				break
			}
		}
	}
	return
}
