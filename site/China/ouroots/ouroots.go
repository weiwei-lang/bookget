package ouroots

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"bookget/lib/util"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
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
		return "requested URL was not found.", err
	}
	dt.SavePath = config.CreateDirectory(dt.Url, dt.BookId)

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	userKey := ""
	token, err := getToken()
	if err != nil {
		return "token not found.", err
	}
	respVolume, err := getVolumes(dt.BookId)
	if err != nil || respVolume.StatusCode != "200" {
		return "requested URL was not found.", err
	}
	index := 0
	for k, vol := range respVolume.Volume {
		log.Printf(" %d/%d volume, %d pages \n", k+1, len(respVolume.Volume), vol.Pages)
		for i := 1; i <= vol.Pages; i++ {
			respImage, err := getBase64Image(dt.BookId, vol.VolumeId, i, userKey, token)
			if err != nil || respImage.StatusCode != "200" {
				log.Println(err)
				continue
			}
			if pos := strings.Index(respImage.ImagePath, "data:image/jpeg;base64,"); pos != -1 {
				data := respImage.ImagePath[pos+len("data:image/jpeg;base64,"):]
				bs, err := base64.StdEncoding.DecodeString(data)
				if err != nil || bs == nil {
					log.Println(err)
					continue
				}
				//
				sortId := fmt.Sprintf("%s.jpg", util.GenNumberSorted(index+1))
				dest := config.GetDestPath(dt.Url, dt.BookId, sortId)

				log.Printf("Get %d/%d volume, %d/%d pages.\n", k+1, len(respVolume.Volume), i, vol.Pages)
				util.FileWrite(bs, dest)
				index++
			}
		}
	}
	fmt.Println()
	return "", nil
}

func getBookId(text string) string {
	text = strings.ToLower(text)
	var bookId string
	m := regexp.MustCompile(`\.html\?([A-z0-9]+)`).FindStringSubmatch(text)
	if m != nil {
		return m[1]
	}
	return bookId
}

func getToken() (string, error) {
	apiUrl := "http://dsNode.ouroots.nlc.cn/loginAnonymousUser"
	bs, err := getBody(apiUrl)
	if err != nil {
		return "", err
	}
	var respLoginAnonymousUser ResponseLoginAnonymousUser
	if err = json.Unmarshal(bs, &respLoginAnonymousUser); err != nil {
		return "", err
	}
	return respLoginAnonymousUser.Token, nil
}

func getVolumes(catalogKey string) (respVolume ResponseVolume, err error) {
	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
		Query: map[string]interface{}{
			"catalogKey": catalogKey,
			"bookid":     "", //目录索引，不重要
		},
	})
	resp, err := cli.Get("http://dsnode.ouroots.nlc.cn/gtService/data/catalogVolume")
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return
	}
	if err = json.Unmarshal(bs, &respVolume); err != nil {
		return
	}
	return respVolume, nil
}

func getBase64Image(catalogKey string, volumeId, page int, userKey, token string) (respImage ResponseCatalogImage, err error) {
	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
		Query: map[string]interface{}{
			"catalogKey": catalogKey,
			"volumeId":   strconv.FormatInt(int64(volumeId), 10),
			"page":       strconv.FormatInt(int64(page), 10),
			"userKey":    userKey,
			"token":      token,
		},
	})
	resp, err := cli.Get("http://dsnode.ouroots.nlc.cn/data/catalogImage")
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return
	}
	if err = json.Unmarshal(bs, &respImage); err != nil {
		return
	}
	return respImage, nil
}

func getBody(apiUrl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent": config.Conf.UserAgent,
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	bs, _ := resp.GetBody()
	if bs == nil {
		err = errors.New(resp.GetReasonPhrase())
		return nil, err
	}
	return bs, nil
}
