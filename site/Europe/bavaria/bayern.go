package bavaria

import (
	"bookget/app"
	"bookget/config"
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
		config.CreateDirectory(taskUrl, bookId)
		StartDownload(iTask, taskUrl, bookId)
	}
	return "", err
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)

	manifestUrl := fmt.Sprintf("https://api.digitale-sammlungen.de/iiif/presentation/v2/%s/manifest", bookId)
	//iiif.StartDownload(manifestUrl, bookId)
	var iiif app.IIIF
	iiif.InitWithId(iTask, manifestUrl, bookId)
	return
}
