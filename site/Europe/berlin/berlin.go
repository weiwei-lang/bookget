package berlin

import (
	"bookget/app"
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	taskName := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", taskName, taskUrl)
	bookId := getBookId(taskUrl)
	config.CreateDirectory(taskUrl, bookId)

	StartDownload(taskUrl, bookId)
	return "", err
}

func getBookId(taskUrl string) string {
	m := regexp.MustCompile(`PPN=([A-Za-z0-9_-]+)`).FindStringSubmatch(taskUrl)
	if m != nil {
		return m[1]
	}
	return ""
}

func StartDownload(pageUrl, bookId string) {
	canvases := getCanvases(bookId)
	if canvases.Size == 0 {
		return
	}
	log.Printf(" %d pages.\n", canvases.Size)

	config.CreateDirectory(pageUrl, bookId)
	if config.Conf.UseDziRs {
		app.DziDownload(pageUrl, bookId, canvases.IiifUrls)
	} else {
		app.NormalDownload(pageUrl, bookId, canvases.ImgUrls, nil)
	}
	return
}

func getCanvases(bookId string) (canvases Canvases) {
	apiUrl := fmt.Sprintf("https://content.staatsbibliothek-berlin.de/dc/%s/manifest", bookId)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	var manifest = new(Manifest)
	if err = json.Unmarshal(bs, manifest); err != nil {
		log.Printf("json.Unmarshal failed: %s\n", err)
		return
	}
	if len(manifest.Sequences) == 0 {
		return
	}

	size := len(manifest.Sequences[0].Canvases)
	canvases.ImgUrls = make([]string, 0, size)
	canvases.IiifUrls = make([]string, 0, size)
	for _, canvase := range manifest.Sequences[0].Canvases {
		for _, image := range canvase.Images {
			//dezoomify-rs URL
			//https://ngcs-core.staatsbibliothek-berlin.de/dzi/PPN3303598630/PHYS_0001.dzi
			m := regexp.MustCompile("/dc/([A-z0-9]+)-([A-z0-9]+)/full").FindStringSubmatch(image.Resource.Id)
			iiiInfo := fmt.Sprintf("https://ngcs-core.staatsbibliothek-berlin.de/dzi/%s/PHYS_%s.dzi", bookId, m[2])
			canvases.IiifUrls = append(canvases.IiifUrls, iiiInfo)

			//JPEG URL
			//https://content.staatsbibliothek-berlin.de/dc/3303598630-0001/full/full/0/default.jpg
			canvases.ImgUrls = append(canvases.ImgUrls, image.Resource.Id)
		}
	}
	canvases.Size = size
	return
}

// tif
func singleImageUrl(bookId, id string) string {
	uri := fmt.Sprintf("https://content.staatsbibliothek-berlin.de/dms/%s/full/0/0000%s.jpg?original=true",
		bookId, id)
	//https://content.staatsbibliothek-berlin.de/dms/PPN3303598630/full/0/00000001.jpg?original=true
	return uri
}

func singleDziUrl(bookId, id string) string {
	uri := fmt.Sprintf("https://content.staatsbibliothek-berlin.de/?action=metsImage&metsFile=%s&divID=%s&dzi=true", bookId, id)
	//uri := fmt.Sprintf("https://ngcs-core.staatsbibliothek-berlin.de/dzi/%s/%s.dzi", bookId, id)
	return uri
}
