package gprime

import (
	"net/http/cookiejar"
	"net/url"
)

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
	Title     string
	Jar       *cookiejar.Jar
}

type ResponseImage struct {
	MetaDataId string   `json:"metaDataId"`
	ImagePath  []string `json:"imagePath"`
	Size       int      `json:"size"`
	IsNext     bool     `json:"isNext"`
	Page       int      `json:"page"`
	Total      int      `json:"total"`
}
