package waseda

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"fmt"
	"log"
	"regexp"
	"sort"
)

// 自定义一个排序类型
type strs []string

func (s strs) Len() int           { return len(s) }
func (s strs) Less(i, j int) bool { return s[i] < s[j] }
func (s strs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Init(iTask int, taskUrl string) (msg string, err error) {
	bookId := ""
	m := regexp.MustCompile(`kosho/[A-Za-z0-9_-]+/([A-Za-z0-9_-]+)/`).FindStringSubmatch(taskUrl)
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

	pdfUrls, size := getMultiplebooks(taskUrl)
	if pdfUrls == nil || size == 0 {
		return
	}
	log.Printf(" %d pages.\n", size)

	for i, uri := range pdfUrls {
		if config.SeqContinue(i) {
			continue
		}
		if uri == "" {
			continue
		}
		ext := util.FileExt(uri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(taskUrl, bookId, fileName)
		gohttp.FastGet(uri, gohttp.Options{
			DestFile:    dest,
			Overwrite:   false,
			Concurrency: config.Conf.Threads,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
}

func getMultiplebooks(bookUrl string) (pdfUrls []string, size int) {
	bs, err := curl.Get(bookUrl, nil)
	if err != nil {
		return
	}
	text := string(bs)
	//取册数
	matches := regexp.MustCompile(`href=["'](.+?)\.pdf["']`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return
	}
	ids := make([]string, 0, len(matches))
	for _, match := range matches {
		ids = append(ids, match[1])
	}
	sort.Sort(strs(ids))
	pdfUrls = make([]string, 0, len(ids))
	for _, v := range ids {
		s := fmt.Sprintf("%s%s.pdf", bookUrl, v)
		pdfUrls = append(pdfUrls, s)
	}
	size = len(pdfUrls)
	return
}
