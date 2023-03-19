package hathitrust

import (
	"bookget/config"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"fmt"
	"log"
	"net/http/cookiejar"
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

	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(uri)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
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
	ext := config.Conf.FileExt
	format := "jpeg"
	if ext == ".png" {
		format = "png"
	} else if ext == ".tif" {
		format = "tiff"
	}
	for i := 0; i < size; i++ {
		if !config.PageRange(i, size) {
			continue
		}
		for true {
			sortId := util.GenNumberSorted(i + 1)
			imgurl := fmt.Sprintf("https://babel.hathitrust.org/cgi/imgsrv/image?id=%s&attachment=1&size=full&format=image/%s&seq=%d", bookId, format, i+1)
			log.Printf("Get %d/%d  %s\n", i+1, size, imgurl)

			fileName := sortId + ext
			dest := config.GetDestPath(uri, bookId, fileName)

			opts := gohttp.Options{
				DestFile:    dest,
				Overwrite:   false,
				Concurrency: 1,
				CookieFile:  config.Conf.CookieFile,
				CookieJar:   jar,
				Headers: map[string]interface{}{
					"User-Agent": config.Conf.UserAgent,
				},
			}
			_, err := gohttp.FastGet(imgurl, opts)
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
