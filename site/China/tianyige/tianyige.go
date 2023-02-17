package tianyige

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

func Init(iTask int, sUrl string) (msg string, err error) {
	dt := new(DownloadTask)
	dt.UrlParsed, err = url.Parse(sUrl)
	dt.Url = sUrl
	dt.Index = iTask
	return Download(dt)
}

func Download(dt *DownloadTask) (msg string, err error) {
	dt.BookId = getBookId(dt.Url)
	if dt.BookId == "" {
		return "", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	imageIds, err := getImageIds(dt.BookId, config.Conf.CookieFile)
	if err != nil {
		log.Println("A cookie file is required.")
		return
	}
	maxSize := len(imageIds)
	log.Printf(" %d pages.\n", maxSize)
	//用户自定义起始页
	i := util.LoopIndexStart(maxSize)
	for ; i < maxSize; i++ {
		uri, _, err := getImageById(imageIds[i], config.Conf.CookieFile)
		if uri == "" || err != nil {
			continue
		}
		ext := util.FileExt(uri)
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, uri)
		fileName := sortId + ext
		dest := config.GetDestPath(dt.Url, dt.BookId, fileName)
		gohttp.FastGet(uri, gohttp.Options{
			Concurrency: config.Conf.Threads,
			DestFile:    dest,
			Overwrite:   false,
			Headers: map[string]interface{}{
				"user-agent": config.UserAgent,
			},
		})
	}
	return "", nil
}

func getBookId(text string) string {
	sUrl := strings.ToLower(text)
	bookId := ""
	m := regexp.MustCompile(`searchpage/([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
	if m != nil {
		bookId = m[1]
	} else {
		m = regexp.MustCompile(`catalogid=([A-z0-9_-]+)`).FindStringSubmatch(sUrl)
		if m != nil {
			bookId = m[1]
		}
	}
	return bookId
}

func getImageIds(imageId string, cookieFile string) (imageIds []string, err error) {
	//https://gj.tianyige.com.cn/fileUpload/56956d82679111ec85ee7020840b69ac/ANB/ANB_IMAGE_PHOTO/ANB/ANB_IMAGE_PHOTO/20220324/febd8c1dcd134c33b5c1cad8883dd1cd1648107167499.jpg
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	apiUrl := fmt.Sprintf("https://gj.tianyige.com.cn/g/sw-anb/api/queryImageByCatalog?catalogId=%s", imageId)

	token := getToken()
	type dataParam struct {
		Param struct {
			PageNum  int `json:"pageNum"`
			PageSize int `json:"pageSize"`
		} `json:"param"`
	}

	dataJson := dataParam{}
	dataJson.Param.PageNum = 1
	dataJson.Param.PageSize = 999

	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent":   config.Conf.UserAgent,
			"Content-Type": "application/json;charset=UTF-8",
			"token":        token,
			"appId":        APP_ID,
		},
		JSON: dataJson,
	})
	resp, err := cli.Post(apiUrl)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	if bs == nil || resp.GetStatusCode() == 401 {
		fmt.Printf("Please try again later.[401 %s]\n", resp.GetReasonPhrase())
		return
	}
	var resObj ResponsePage
	if err = json.Unmarshal(bs, &resObj); resObj.Code != 200 {
		return
	}
	imageIds = make([]string, 0, len(resObj.Data.Records))
	for _, d := range resObj.Data.Records {
		imageIds = append(imageIds, d.ImageId)
	}
	return
}

func getCanvases(imageIds []string, cookieFile string) (canvases Canvases) {
	fmt.Println()
	for i, id := range imageIds {
		imgUrl, ocrUrl, err := getImageById(id, cookieFile)
		if err != nil {
			continue
		}
		sortId := util.GenNumberSorted(i + 1)
		fmt.Printf("\r Test page %s ... ", sortId)
		canvases.ImgUrls = append(canvases.ImgUrls, imgUrl)
		canvases.ImgOcrUrls = append(canvases.ImgOcrUrls, ocrUrl)
	}
	fmt.Println()
	canvases.Size = len(canvases.ImgUrls)
	return
}

func getImageById(imageId, cookieFile string) (imgUrl, ocrUrl string, err error) {
	//https://gj.tianyige.com.cn/fileUpload/56956d82679111ec85ee7020840b69ac/ANB/ANB_IMAGE_PHOTO/ANB/ANB_IMAGE_PHOTO/20220324/febd8c1dcd134c33b5c1cad8883dd1cd1648107167499.jpg
	//cookie 处理
	jar, _ := cookiejar.New(nil)
	apiUrl := fmt.Sprintf("https://gj.tianyige.com.cn/g/sw-anb/api/queryOcrFileByimageId?imageId=%s", imageId)

	token := getToken()
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent":   config.Conf.UserAgent,
			"Content-Type": "application/json;charset=UTF-8",
			"token":        token,
			"appId":        APP_ID,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return
	}
	bs, _ := resp.GetBody()
	if bs == nil || resp.GetStatusCode() == 401 {
		fmt.Printf("Please try again later.[401 %s]\n", resp.GetReasonPhrase())
		return
	}
	var resObj ResponseFile
	if err = json.Unmarshal(bs, &resObj); err != nil {
		fmt.Println(err)
		return
	}

	for _, ossFile := range resObj.Data.File {
		if strings.Contains(ossFile.FileOldname, "_c") {
			ocrUrl = fmt.Sprintf("https://gj.tianyige.com.cn/fileUpload/%s/%s", ossFile.FilePath, ossFile.FileName)
		} else {
			imgUrl = fmt.Sprintf("https://gj.tianyige.com.cn/fileUpload/%s/%s", ossFile.FilePath, ossFile.FileName)
		}
	}
	return
}
