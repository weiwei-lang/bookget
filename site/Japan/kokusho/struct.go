package kokusho

import "net/url"

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}
type ResponseDetail struct {
	Bid              string        `json:"bid"`
	Shubetsu         string        `json:"shubetsu"`
	Hshomei          string        `json:"hshomei"`
	Kshomei          []interface{} `json:"kshomei"`
	Hshomeipdf       string        `json:"hshomeipdf"`
	Kshomeipdf       []interface{} `json:"kshomeipdf"`
	Author           []interface{} `json:"author"`
	Kansu            string        `json:"kansu"`
	Kansha           string        `json:"kansha"`
	Writing          []string      `json:"writing"`
	Shuppan          []interface{} `json:"shuppan"`
	Keitai           string        `json:"keitai"`
	Satsu            string        `json:"satsu"`
	Zan              string        `json:"zan"`
	Structure        string        `json:"structure"`
	Keyword          string        `json:"keyword"`
	Chuki            string        `json:"chuki"`
	Collection       string        `json:"collection"`
	WSeikyu          string        `json:"w_seikyu"`
	MSeikyu          string        `json:"m_seikyu"`
	CSeikyu          string        `json:"c_seikyu"`
	DSeikyu          string        `json:"d_seikyu"`
	WCallno          string        `json:"w_callno"`
	MCallno          string        `json:"m_callno"`
	CCallno          string        `json:"c_callno"`
	DCallno          string        `json:"d_callno"`
	MKoma            string        `json:"m_koma"`
	CKoma            string        `json:"c_koma"`
	DKoma            string        `json:"d_koma"`
	MSc              string        `json:"m_sc"`
	CSc              string        `json:"c_sc"`
	DSc              string        `json:"d_sc"`
	Parentbiblio     string        `json:"parentbiblio"`
	Parentbiblioname string        `json:"parentbiblioname"`
	CardSno          string        `json:"card_sno"`
	Seikyuoutput     string        `json:"seikyuoutput"`
	Image            string        `json:"image"`
	Manifestlabel    string        `json:"manifestlabel"`
	Manifest         string        `json:"manifest"`
	Licenseimg       string        `json:"licenseimg"`
	Licensetext1     string        `json:"licensetext1"`
	Licensetext2     string        `json:"licensetext2"`
	Licenselink      string        `json:"licenselink"`
	Imagedownload    string        `json:"imagedownload"`
	Imageindex       []interface{} `json:"imageindex"`
	Taglist          []interface{} `json:"taglist"`
	Reprinttext      []interface{} `json:"reprinttext"`
	Work             []struct {
		Wid           string        `json:"wid"`
		Name          string        `json:"name"`
		Satsu         string        `json:"satsu"`
		Tsuno         string        `json:"tsuno"`
		Year          string        `json:"year"`
		Keyword       string        `json:"keyword"`
		Besho         []interface{} `json:"besho"`
		Note          string        `json:"note"`
		Author        []interface{} `json:"author"`
		Shubetsu      string        `json:"shubetsu"`
		Kokusho       string        `json:"kokusho"`
		Kokushono     string        `json:"kokushono"`
		Kokushoshozai string        `json:"kokushoshozai"`
	} `json:"work"`
	Doi         string `json:"doi"`
	Outsidelink []struct {
		Label   string `json:"label"`
		LabelEn string `json:"label_en"`
		Url     string `json:"url"`
	} `json:"outsidelink"`
	Childbiblio []interface{} `json:"childbiblio"`
}
