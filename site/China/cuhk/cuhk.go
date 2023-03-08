package cuhk

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := getBookId(taskUrl)
	if bookId == "" {
		return "", errors.New("requested URL was not found")
	}
	StartDownload(iTask, taskUrl, bookId)
	return "", nil
}

func getBookId(bookUrl string) (bookId string) {
	m := regexp.MustCompile(`item/cuhk-([A-Za-z0-9]+)`).FindStringSubmatch(bookUrl)
	if m != nil {
		bookId = m[1]
	}
	return
}

func StartDownload(iTask int, taskUrl, bookId string) {
	name := util.GenNumberSorted(iTask)
	log.Printf("Get %s  %s\n", name, taskUrl)
	bookUrls := getMultiplebooks(taskUrl)
	if bookUrls == nil || len(bookUrls) == 0 {
		log.Printf("requested URL was not found.\n")
		return
	}
	size := len(bookUrls)
	log.Printf(" %d volumes \n", size)
	for i := 0; i < size; i++ {
		uri := bookUrls[i]
		log.Printf("Test volume %d ... \n", i+1)
		id := fmt.Sprintf("%s_volume%s", bookId, util.GenNumberSorted(i+1))
		config.CreateDirectory(uri, id)
		do(id, uri)
	}
}

func do(bookId, bookUrl string) {
	imagePages, cookies := getJpg2000Urls(bookUrl)
	if imagePages == nil {
		log.Printf("requested URL was not found.")
		return
	}
	size := len(imagePages)
	log.Printf(" %d pages.\n", size)
	sCookie := curl.HttpCookie2String(cookies)
	for i, v := range imagePages {
		if config.SeqContinue(i) {
			continue
		}
		imgUrl := formateUrl(v.Pid)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, imgUrl)

		filename := sortId + ".jp2"
		dest := config.GetDestPath(bookUrl, bookId, filename)

		header := make(map[string]string, 8)
		header["Cookie"] = sCookie
		header["Referer"] = bookUrl
		if v.Token != "" {
			header["X-ISLANDORA-TOKEN"] = v.Token
		}
		curl.FastGet(imgUrl, dest, header, true)
		util.PrintSleepTime(config.Conf.Speed)
	}
}

func getMultiplebooks(bookUrl string) (uri []string) {
	bs, err := curl.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	subText := util.SubText(text, "id=\"block-islandora-compound-object-compound-navigation-select-list\"", "id=\"book-viewer\">")
	matches := regexp.MustCompile(`value=['"]([A-z\d:_-]+)['"]`).FindAllStringSubmatch(subText, -1)
	if matches == nil {
		uri = append(uri, bookUrl)
		return
	}
	for _, m := range matches {
		//value='ignore'
		if m[1] == "ignore" {
			continue
		}
		id := strings.Replace(m[1], ":", "-", 1)
		uri = append(uri, fmt.Sprintf("https://repository.lib.cuhk.edu.hk/sc/item/%s#page/1/mode/2up", id))
	}
	return
}

// header.Set("X-ISLANDORA-TOKEN")
func getTokenHeader(text string) bool {
	return true //TODO
	useToken := false
	matchToken := regexp.MustCompile(`"tokenHeader":([a-zA-z]+)`).FindStringSubmatch(text)
	if matchToken != nil {
		//请求头信息是否包含token
		if strings.ToLower(matchToken[1]) == "true" {
			useToken = true
		}
	}
	return useToken
}

func getJpg2000Urls(bookUrl string) (imagePage []ImagePage, c []*http.Cookie) {
	bs, c, err := curl.GetWithCookie(bookUrl, nil)
	if err != nil {
		return
	}
	var resp ResponsePage
	matches := regexp.MustCompile(`"pages":([^]]+)]`).FindSubmatch(bs)
	if matches != nil {
		data := []byte("{\"pages\":" + string(matches[1]) + "]}")
		if err = json.Unmarshal(data, &resp); err != nil {
			log.Printf("json.Unmarshal failed: %s\n", err)
		}
		imagePage = make([]ImagePage, len(resp.ImagePage))
		copy(imagePage, resp.ImagePage)
	}
	return imagePage, c
}

func formateUrl(id string) string {
	template := "https://repository.lib.cuhk.edu.hk/islandora/object/%s/datastream/JP2"
	//imgUrl := fmt.Sprintf("https://repository.lib.cuhk.edu.hk/%s/full/full/0/default.jpg", page.Identifier)
	return fmt.Sprintf(template, id)
}
