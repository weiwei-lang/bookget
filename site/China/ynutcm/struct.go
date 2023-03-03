package ynutcm

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

type ResponseDetail struct {
	DataFields []struct {
		Id               int         `json:"Id"`
		Name             string      `json:"Name"`
		ShowName         string      `json:"ShowName"`
		Type             string      `json:"Type"`
		ContentType      string      `json:"ContentType"`
		DictKeyId        interface{} `json:"DictKeyId"`
		Options          interface{} `json:"Options"`
		IsMultiple       string      `json:"IsMultiple"`
		IsRepeat         string      `json:"IsRepeat"`
		DataType         string      `json:"DataType"`
		BindDataField    string      `json:"BindDataField"`
		Abbreviation     interface{} `json:"Abbreviation"`
		Placeholder      interface{} `json:"Placeholder"`
		Unit             interface{} `json:"Unit"`
		Description      *string     `json:"Description"`
		SortOrder        int         `json:"SortOrder"`
		OperateAuthority string      `json:"OperateAuthority"`
		IsRequired       string      `json:"IsRequired"`
		DataSource       interface{} `json:"DataSource"`
		DefaultValue     interface{} `json:"DefaultValue"`
		IsSearch         string      `json:"IsSearch"`
		IsUnifySearch    string      `json:"IsUnifySearch"`
		GroupType        string      `json:"GroupType"`
		GroupId          int         `json:"GroupId"`
		CreateIp         string      `json:"CreateIp"`
		CreateBy         int         `json:"CreateBy"`
		CreateTime       string      `json:"CreateTime"`
		IsDeleted        string      `json:"IsDeleted"`
		SiteId           int         `json:"SiteId"`
	} `json:"dataFields"`
	Detail struct {
		Creator      string      `json:"creator"`
		Title        string      `json:"title"`
		Author       string      `json:"author"`
		Number       string      `json:"number"`
		Totalnum     string      `json:"totalnum"`
		Callnum      string      `json:"callnum"`
		DocNo        string      `json:"docNo"`
		Class        string      `json:"class"`
		Age          string      `json:"age"`
		Version      string      `json:"version"`
		Vol          interface{} `json:"vol"`
		CopiesCount  string      `json:"copiesCount"`
		Volume       string      `json:"volume"`
		Readirection string      `json:"readirection"`
		BooSize      string      `json:"booSize"`
		Binding      string      `json:"binding"`
		Content      string      `json:"content"`
		Organs       interface{} `json:"organs"`
		Location     string      `json:"location"`
		IsRecommend  bool        `json:"isRecommend"`
		FullTextPath string      `json:"fullTextPath"`
		Cover        string      `json:"cover"`
		Id           int         `json:"Id"`
		GroupType    string      `json:"GroupType"`
		GroupId      int         `json:"GroupId"`
		ViewCount    int         `json:"ViewCount"`
		PraiseCount  int         `json:"PraiseCount"`
		CollectCount int         `json:"CollectCount"`
		DownCount    int         `json:"DownCount"`
		CreateIp     string      `json:"CreateIp"`
		CreateBy     interface{} `json:"CreateBy"`
		CreateTime   interface{} `json:"CreateTime"`
		UpdateTime   string      `json:"UpdateTime"`
		IsDeleted    string      `json:"IsDeleted"`
		SiteId       int         `json:"SiteId"`
	} `json:"detail"`
	Fulltextpath []struct {
		Name  string `json:"name"`
		Tpath string `json:"tpath"`
	} `json:"fulltextpath"`
	Channel struct {
		Id             int         `json:"Id"`
		ParentId       interface{} `json:"ParentId"`
		Name           string      `json:"Name"`
		Unikey         string      `json:"Unikey"`
		Module         string      `json:"Module"`
		CreateTempalte interface{} `json:"CreateTempalte"`
		ListTemplate   string      `json:"ListTemplate"`
		DetailTemplate string      `json:"DetailTemplate"`
		CrawlWebIds    interface{} `json:"CrawlWebIds"`
		ZtRule         interface{} `json:"ZtRule"`
		SortIndex      int         `json:"SortIndex"`
		SiteId         int         `json:"SiteId"`
		IsWork         string      `json:"IsWork"`
		IsDeleted      string      `json:"IsDeleted"`
	} `json:"channel"`
	IsCollect bool `json:"isCollect"`
	Aboutlist []struct {
		Creator      string      `json:"creator"`
		Title        string      `json:"title"`
		Author       []string    `json:"author"`
		Number       string      `json:"number"`
		Totalnum     string      `json:"totalnum"`
		Callnum      string      `json:"callnum"`
		DocNo        string      `json:"docNo"`
		Class        []string    `json:"class"`
		Age          string      `json:"age"`
		Version      string      `json:"version"`
		Vol          interface{} `json:"vol"`
		CopiesCount  string      `json:"copiesCount"`
		Volume       string      `json:"volume"`
		Readirection string      `json:"readirection"`
		BooSize      string      `json:"booSize"`
		Binding      string      `json:"binding"`
		Content      string      `json:"content"`
		Organs       interface{} `json:"organs"`
		Location     string      `json:"location"`
		IsRecommend  string      `json:"isRecommend"`
		FullTextPath string      `json:"fullTextPath"`
		Cover        string      `json:"cover"`
		Id           string      `json:"Id"`
		GroupType    string      `json:"GroupType"`
		GroupId      string      `json:"GroupId"`
		ViewCount    int         `json:"ViewCount"`
		PraiseCount  string      `json:"PraiseCount"`
		CollectCount string      `json:"CollectCount"`
		DownCount    string      `json:"DownCount"`
		CreateIp     string      `json:"CreateIp"`
		CreateBy     interface{} `json:"CreateBy"`
		CreateTime   interface{} `json:"CreateTime"`
		UpdateTime   string      `json:"UpdateTime"`
		IsDeleted    string      `json:"IsDeleted"`
		SiteId       string      `json:"SiteId"`
	} `json:"aboutlist"`
}
