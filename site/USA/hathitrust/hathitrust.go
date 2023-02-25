package hathitrust

import (
	"bookget/config"
	curl "bookget/lib/curl"
	util "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`id=([^&]+)`).FindStringSubmatch(taskUrl)
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

	bs, err := curl.Get(uri, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取页数
	// <input id="range-seq" class="navigator-range" type="range" min="1" max="1036" value="2" aria-label="Progress" dir="rtl" />
	matches := regexp.MustCompile(`<input(?:[^>]+)id="range-seq"(?:[^>]+)max="([0-9]+)"`).FindStringSubmatch(text)
	if matches == nil {
		return
	}
	size := 0
	if matches[1] != "" {
		size, _ = strconv.Atoi(matches[1])
	}
	log.Printf(" %d pages.\n", size)
	ext := ".jpeg"
	for i := 0; i < size; i++ {
		for true {
			sortId := util.GenNumberSorted(i + 1)
			imgurl := fmt.Sprintf("https://babel.hathitrust.org/cgi/imgsrv/image?id=%s&attachment=1&size=full&format=image/jpeg&seq=%d", bookId, i+1)
			log.Printf("Get %s  %s\n", sortId, imgurl)

			fileName := sortId + ext
			dest := config.GetDestPath(uri, bookId, fileName)

			header := make(map[string]string)
			_, err = curl.FastGet(imgurl, dest, header, true)
			if err != nil {
				fmt.Println(err)
				//log.Println("images (1 file per page, watermarked,  max. 20 MB / 1 min), image quality:Full")
				util.PrintSleepTime(60)
				continue
			}
			break
		}
	}

}
