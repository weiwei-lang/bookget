package ncpssd

import "net/url"

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}

type ResponseServerDate struct {
	Result bool   `json:"result"`
	Code   int    `json:"code"`
	Data   string `json:"data"`
	Succee bool   `json:"succee"`
}

type ResponseReadUrl struct {
	Url string `json:"url"`
}
