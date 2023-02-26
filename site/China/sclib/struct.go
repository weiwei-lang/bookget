package sclib

// 自定义一个排序类型
type strs []string

func (s strs) Len() int           { return len(s) }
func (s strs) Less(i, j int) bool { return s[i] < s[j] }
func (s strs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type Item struct {
	Extension   string `json:"extension"`
	Height      int    `json:"height"`
	Resolutions int    `json:"resolutions"`
	TileSize    struct {
		H int `json:"h"`
		W int `json:"w"`
	} `json:"tile_size"`
	TileSize2 struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"tileSize"`
	Width int `json:"width"`
}

type ResponseBody struct {
	Tiles map[string]Item `json:"tiles"`
}

type ResponseServerBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Title      string   `json:"title"`
		ServerBase string   `json:"serverBase"`
		Images     []string `json:"images"`
	} `json:"data"`
}
