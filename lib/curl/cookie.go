package curl

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func GetHeaderFile(srcPath string) (header map[string]string, err error) {
	fp, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer fp.Close()

	bsHeader, err := io.ReadAll(fp)
	if err != nil {
		return
	}
	sHeader := string(bsHeader)
	mHeader := strings.Split(sHeader, "\n")

	header = make(map[string]string, 10)
	for _, line := range mHeader {
		s := strings.Trim(line, "\r")
		i := strings.Index(s, ":")
		if i == -1 {
			continue
		}
		k := s[:i]
		v := strings.Trim(s[i+1:], " ")
		if v == "" {
			continue
		}
		if "cookie" == strings.ToLower(k) {
			header["Cookie"] = CookieEscape(v)
		} else if "user-agent" == strings.ToLower(k) {
			header["User-Agent"] = v
		} else {
			header[k] = v
		}
	}
	return header, nil
}

func CookieEscape(text string) string {
	matches := regexp.MustCompile(`([^=]+)\=([^;]+);`).FindAllStringSubmatch(text, -1)
	if matches == nil {
		return text
	}
	cookieText := ""
	for _, m := range matches {
		k := strings.TrimSpace(m[1])
		v := strings.TrimSpace(m[2])
		cookieText += k + "=" + url.QueryEscape(v) + ";"
	}
	return cookieText
}

func HttpCookie2String(cookie []*http.Cookie) string {
	sCookie := ""
	if cookie != nil {
		for _, c := range cookie {
			sCookie += fmt.Sprintf("%s=%s;", c.Name, c.Value)
		}
	}
	return sCookie
}

func GetHeaderFmtValues(srcPath string) (header url.Values, err error) {
	fp, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer fp.Close()
	bsHeader, err := io.ReadAll(fp)
	if err != nil {
		return
	}
	sHeader := string(bsHeader)
	mHeader := strings.Split(sHeader, "\n")

	header = url.Values{}
	for _, line := range mHeader {
		m := strings.Split(strings.Trim(line, "\r"), ":")
		if len(m) != 2 {
			continue
		}
		if "cookie" == strings.ToLower(m[0]) {
			header.Set("Cookie", m[1])
		} else if "user-agent" == strings.ToLower(m[0]) {
			header.Set("User-Agent", m[1])
		} else {
			header.Set(m[0], m[1])
		}
	}
	return header, nil
}
