package sdutcm

import "net/url"

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}

type ResponseBody struct {
	Url       string `json:"url"`
	Text      string `json:"text"`
	Charmax   int    `json:"charmax"`
	ColNum    int    `json:"colNum"`
	PageNum   string `json:"pageNum"`
	ImageList struct {
	} `json:"imageList"`
}

type Canvases struct {
	ImgUrls []string
	Size    int
}
