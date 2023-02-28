package bjdpm

import (
	xcrypt "bookget/lib/crypt"
	"fmt"
	"strconv"
	"strings"
)

const (
	//故宫名画 minghuaji.dpm.org.cn
	AES_KEY = "hQKWqRCPUFjUXv0q"
	AES_IV  = "SH8csHyhBEnAPtwb"

	//数字文物 digicol.dpm.org.cn
	AES2_KEY = "tNzf3IrAXDCepOVQ"
	AES2_IV  = "nE0d1QQdSy45uBX3"
)

func getDziJson(host string, text []byte) (dziJson string, dzi DziFormat) {
	template := `{
  "xmlns": "http://schemas.microsoft.com/deepzoom/2009",
  "Url": "%s",
  "Overlap": "%d",
  "TileSize": "%d",
  "Format": "%s",
  "Size": {
    "Width": "%d",
    "Height": "%d"
  }
}
`
	var recoveredPt []byte
	var err error
	if host == "digicol.dpm.org.cn" {
		recoveredPt, err = xcrypt.DecryptByAes(string(text), []byte(AES2_KEY), []byte(AES2_IV))
	} else {
		recoveredPt, err = xcrypt.DecryptByAes(string(text), []byte(AES_KEY), []byte(AES_IV))
	}
	if err != nil {
		return
	}
	m := strings.Split(string(recoveredPt), "^")
	if m == nil || len(m) != 6 {
		return
	}
	//fmt.Printf("Split plaintext: %+v\n", m)
	dzi.Url = m[0]
	dzi.Format = m[1]
	dzi.TileSize, _ = strconv.Atoi(m[4])
	dzi.Overlap, _ = strconv.Atoi(m[5])
	if strings.Contains(m[2], ".") {
		w, _ := strconv.ParseFloat(m[2], 32)
		dzi.Size.Width = int(w)
	} else {
		dzi.Size.Width, _ = strconv.Atoi(m[2])
	}
	if strings.Contains(m[3], ".") {
		h, _ := strconv.ParseFloat(m[3], 32)
		dzi.Size.Height = int(h)
	} else {
		dzi.Size.Height, _ = strconv.Atoi(m[3])
	}
	dziJson = fmt.Sprintf(template, dzi.Url, dzi.Overlap, dzi.TileSize, dzi.Format, dzi.Size.Width, dzi.Size.Height)
	return
}
