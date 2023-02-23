package bjdpm

type DziFormat struct {
	Xmlns    string `json:"xmlns"`
	Url      string `json:"Url"`
	Overlap  int    `json:"Overlap"`
	TileSize int    `json:"TileSize"`
	Format   string `json:"Format"`
	Size     struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
	} `json:"Size"`
}
