package ncpssd

import (
	"bookget/config"
	"bookget/lib/gohttp"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/cookiejar"
)

func getToken(sUrl string) (string, error) {
	apiUrl := "https://www.ncpssd.org/get/server/date"
	bs, err := getBody(apiUrl, sUrl)
	if err != nil {
		return "", err
	}
	var responseServerDate ResponseServerDate
	if err = json.Unmarshal(bs, &responseServerDate); err != nil {
		return "", err
	}
	h := md5.New()
	h.Write([]byte("L!N45S26y1SGzq9^" + responseServerDate.Data))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token, nil
}

func getBody(apiUrl string, sUrl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
	cli := gohttp.NewClient(gohttp.Options{
		CookieFile: config.Conf.CookieFile,
		CookieJar:  jar,
		Headers: map[string]interface{}{
			"User-Agent":       config.Conf.UserAgent,
			"X-Requested-With": "XMLHttpRequest",
			"Referer":          sUrl,
			"Content-Type":     "application/json; charset=utf-8",
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
