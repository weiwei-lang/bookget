package stanford

import (
	"bookget/app"
	"bookget/lib/curl"
	"bookget/lib/util"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`/view/([A-z\d]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		bookId = m[1]
		//config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	bookUrls := getMultiplebooks(taskUrl)
	if bookUrls == nil {
		return
	}
	size := len(bookUrls)
	for i := 0; i < size; i++ {
		log.Printf(" %d/%d volume.\n", i+1, size)
		//iiif.StartDownload(bookUrls[i], fmt.Sprintf("%s_Volume%d", bookId, i+1))
		id := util.GenNumberSorted(i + 1)
		bookId = fmt.Sprintf("%s_vol.%s", bookId, id)
		var iiif app.IIIF
		iiif.InitWithId(iTask, bookUrls[i], bookId)
	}
	return
}

func getMultiplebooks(taskUrl string) (bookUrls []string) {
	bs, err := curl.Get(taskUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	matches := regexp.MustCompile(`data-embed-target\s?=\s?['"]https://purl.stanford.edu/([A-z\d]+)["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	for _, m := range matches {
		manifestUrl := fmt.Sprintf("https://purl.stanford.edu/%s/iiif/manifest", m[1])
		bookUrls = append(bookUrls, manifestUrl)
	}
	return
}
