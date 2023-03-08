package familysearch

import (
	"bookget/config"
	"bookget/lib/curl"
	"bookget/lib/gohttp"
	util "bookget/lib/util"
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

// 家谱图像 https://www.familysearch.org/records/images/
func ImagesDownload(dt *DownloadTask) (msg string, err error) {

	name := util.GenNumberSorted(dt.Index)
	log.Printf("Get %s  %s\n", name, dt.Url)

	canvases, err := getCanvases(dt.BookId, config.Conf.CookieFile)
	if err != nil {
		return "", err
	}
	//用户自定义起始页
	log.Printf(" %d Pages.\n", canvases.Size)

	header, _ := curl.GetHeaderFile(config.Conf.CookieFile)
	args := []string{"--dezoomer=deepzoom",
		"-H", "authority:www.familysearch.org",
		"-H", "referer:" + url.QueryEscape(dt.Url),
		"-H", "User-Agent:" + header["User-Agent"],
		"-H", "cookie:" + header["Cookie"],
	}
	storePath := dt.SavePath + string(os.PathSeparator)
	for i, inputUri := range canvases.IiifUrls {
		if config.SeqContinue(i) {
			continue
		}
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, inputUri)
		outfile := storePath + sortId + config.Conf.FileExt
		util.StartProcess(inputUri, outfile, args)
		util.PrintSleepTime(config.Conf.Speed)
	}

	return "", nil
}

func dasImageDown(dt *DownloadTask, imageUrls []string) {
	jar, _ := cookiejar.New(nil)
	for i, dUrl := range imageUrls {
		if config.SeqContinue(i) {
			continue
		}
		if dUrl == "" {
			continue
		}
		sortId := util.GenNumberSorted(i + 1)
		log.Printf("Get %s  %s\n", sortId, dUrl)
		fileName := sortId + config.Conf.FileExt
		dest := config.GetDestPath(dt.Url, dt.BookId, fileName)
		for {
			_, err := gohttp.FastGet(dUrl, gohttp.Options{
				DestFile:    dest,
				Overwrite:   false,
				Concurrency: config.Conf.Threads,
				CookieJar:   jar,
				CookieFile:  config.Conf.CookieFile,
				Headers: map[string]interface{}{
					"user-agent": config.UserAgent,
				},
			})
			if err != nil {
				fmt.Println(err)
				util.PrintSleepTime(60)
				continue
			}
			break
		}
		util.PrintSleepTime(config.Conf.Speed)
	}

}

func getCanvases(bookId, cookieFile string) (Canvases, error) {
	canvases := Canvases{}
	apiUrl := fmt.Sprintf("https://www.familysearch.org/records/images/api/imageDetails/groups/%s?properties&changeLog&coverageIndex=null", bookId)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: cookieFile,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
	})
	resp, err := cli.Get(apiUrl)
	if err != nil {
		return canvases, err
	}
	imageGroups := ImageGroups{}
	if err = resp.GetJsonDecodeBody(&imageGroups); err != nil {
		return canvases, err
	}
	canvases.IiifUrls = make([]string, 0, imageGroups.VolumeSet.ChildCount)
	canvases.ImageUrls = make([]string, 0, imageGroups.VolumeSet.ChildCount)
	for _, group := range imageGroups.Groups {
		for _, v := range group.ImageUrls {
			dzUrl := fmt.Sprintf("%s/image.xml", v)
			u, _ := url.Parse(v)
			i := strings.LastIndex(u.Path, "/")
			id := u.Path[i+1:]
			imgUrl := fmt.Sprintf("%s://%s/service/records/storage/dascloud/das/v2/%s/dist.jpg?proxy=true", u.Scheme, u.Host, id)
			canvases.IiifUrls = append(canvases.IiifUrls, dzUrl)
			canvases.ImageUrls = append(canvases.ImageUrls, imgUrl)
		}
	}
	canvases.Size = len(canvases.IiifUrls)
	return canvases, nil
}
