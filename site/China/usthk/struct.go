package usthk

import "net/url"

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}

type Canvases struct {
	ImgUrls []string
	Size    int
}

type ResponseFiles struct {
	FileList []string `json:"file_list"`
}
