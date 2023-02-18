package berlin

type Canvases struct {
	ImgUrls  []string
	IiifUrls []string
	Size     int
}

type Manifest struct {
	Sequences []Sequence `json:"sequences"`
}
type Sequence struct {
	Canvases []struct {
		Images []struct {
			Id         string `json:"@id"`
			Type       string `json:"@type"`
			Motivation string `json:"motivation"`
			Resource   struct {
				Id      string `json:"@id"`
				Type    string `json:"@type"`
				Format  string `json:"format"`
				Height  int    `json:"height"`
				Width   int    `json:"width"`
				Service struct {
					Context string `json:"@context"`
					Id      string `json:"@id"`
					Profile string `json:"profile"`
				} `json:"service"`
			} `json:"resource"`
			On string `json:"on"`
		} `json:"images"`
	} `json:"canvases"`
}
