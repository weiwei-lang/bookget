package tianyige

import "net/url"

type DownloadTask struct {
	Index     int
	Url       string
	UrlParsed *url.URL
	SavePath  string
	BookId    string
}

type BookDetail struct {
	CatalogId            string      `json:"catalogId"`
	TitleRolls           string      `json:"titleRolls"`
	TradeTitleRolls      interface{} `json:"tradeTitleRolls"`
	Identifier           string      `json:"identifier"`
	TradeIdentifier      interface{} `json:"tradeIdentifier"`
	CallNumber           string      `json:"callNumber"`
	TradeCallNumber      interface{} `json:"tradeCallNumber"`
	ParentIdentifier     string      `json:"parentIdentifier"`
	SeriesOrder          interface{} `json:"seriesOrder"`
	Type                 int         `json:"type"`
	Title                string      `json:"title"`
	TradeTitle           interface{} `json:"tradeTitle"`
	Author               string      `json:"author"`
	TradeAuthor          interface{} `json:"tradeAuthor"`
	RollEndTitle         string      `json:"rollEndTitle"`
	CoverTitle           string      `json:"coverTitle"`
	CenterTitle          string      `json:"centerTitle"`
	Principal            interface{} `json:"principal"`
	TradePrincipal       interface{} `json:"tradePrincipal"`
	CreatorYear          string      `json:"creatorYear"`
	CreatorDutyType      string      `json:"creatorDutyType"`
	Contributor          string      `json:"contributor"`
	CensusNumber         string      `json:"censusNumber"`
	TradeCensusNumber    interface{} `json:"tradeCensusNumber"`
	ContributorYear      string      `json:"contributorYear"`
	ContributorDutyType  string      `json:"contributorDutyType"`
	Language             string      `json:"language"`
	Subject              string      `json:"subject"`
	CoverageAddress      string      `json:"coverageAddress"`
	CoverageDynasty      string      `json:"coverageDynasty"`
	TradeCoverageDynasty interface{} `json:"tradeCoverageDynasty"`
	MadeCategory         interface{} `json:"madeCategory"`
	PublishDynasty       string      `json:"publishDynasty"`
	TradePublishDynasty  interface{} `json:"tradePublishDynasty"`
	PublishYear          string      `json:"publishYear"`
	GetType              string      `json:"getType"`
	Marking              string      `json:"marking"`
	StoreAddress         string      `json:"storeAddress"`
	StoreNumber          string      `json:"storeNumber"`
	CollectionNumber     interface{} `json:"collectionNumber"`
	Criticism            string      `json:"criticism"`
	TradeCriticism       interface{} `json:"tradeCriticism"`
	CategoryItem         string      `json:"categoryItem"`
	Publisher            string      `json:"publisher"`
	PublishPlace         string      `json:"publishPlace"`
	VersionType          string      `json:"versionType"`
	NationalMapShooting  interface{} `json:"nationalMapShooting"`
	OriginalPublication  interface{} `json:"originalPublication"`
	PackagesNum          string      `json:"packagesNum"`
	Volumes              string      `json:"volumes"`
	HighSize             interface{} `json:"highSize"`
	WideSize             interface{} `json:"wideSize"`
	DamageCondition      string      `json:"damageCondition"`
	Vector               string      `json:"vector"`
	RightDeclare         string      `json:"rightDeclare"`
	ClassCode            string      `json:"classCode"`
	IsRelated            int         `json:"isRelated"`
	IsHide               int         `json:"isHide"`
	IsPublish            int         `json:"isPublish"`
	IsRecommend          int         `json:"isRecommend"`
	IsCollect            int         `json:"isCollect"`
	ScanNumber           string      `json:"scanNumber"`
	DiskNumber           string      `json:"diskNumber"`
	Introduction         interface{} `json:"introduction"`
	IsScaned             int         `json:"isScaned"`
	Clicks               int         `json:"clicks"`
	XmlName              string      `json:"xmlName"`
	Annotate             interface{} `json:"annotate"`
	Auditor              interface{} `json:"auditor"`
	IsCover              int         `json:"isCover"`
	Checknumber          string      `json:"checknumber"`
	Creator              interface{} `json:"creator"`
	CreateTime           string      `json:"createTime"`
	Updator              string      `json:"updator"`
	UpdateTime           string      `json:"updateTime"`
	IsDeleted            int         `json:"isDeleted"`
	FascicleCount        interface{} `json:"fascicleCount"`
	ImageCount           interface{} `json:"imageCount"`
	FilePath             string      `json:"filePath"`
	Fascicles            interface{} `json:"fascicles"`
	AnbCatalogDetail     struct {
		CatalogDetailId              string      `json:"catalogDetailId"`
		CatalogId                    string      `json:"catalogId"`
		Checknumber                  string      `json:"checknumber"`
		Oldchecknumber               string      `json:"oldchecknumber"`
		IndexNo                      string      `json:"indexNo"`
		Classification               string      `json:"classification"`
		TitleOfVolumes               string      `json:"titleOfVolumes"`
		Writer                       string      `json:"writer"`
		Edition                      string      `json:"edition"`
		EditionDate                  string      `json:"editionDate"`
		EditionType                  string      `json:"editionType"`
		Format                       string      `json:"format"`
		BindingForm                  string      `json:"bindingForm"`
		BookNumber                   string      `json:"bookNumber"`
		StorageVolume                string      `json:"storageVolume"`
		SchoolOfInscription          string      `json:"schoolOfInscription"`
		OwnerBookName                string      `json:"ownerBookName"`
		Subtitle                     string      `json:"subtitle"`
		SubtitleAuthor               interface{} `json:"subtitleAuthor"`
		SubtitleDate                 interface{} `json:"subtitleDate"`
		Unit                         string      `json:"unit"`
		Annotation                   string      `json:"annotation"`
		DataProducer                 string      `json:"dataProducer"`
		BasicAnnotation              string      `json:"basicAnnotation"`
		CollectionUnitClassification string      `json:"collectionUnitClassification"`
		ClassifyAnnotation           string      `json:"classifyAnnotation"`
		TitleVolumeNote              string      `json:"titleVolumeNote"`
		AuthorAnnotation             string      `json:"authorAnnotation"`
		AuthorBirth                  string      `json:"authorBirth"`
		OtherNominate                string      `json:"otherNominate"`
		TotalVolume                  string      `json:"totalVolume"`
		ActualVolume                 string      `json:"actualVolume"`
		LackVolume                   string      `json:"lackVolume"`
		VolumeCountAnnotation        string      `json:"volumeCountAnnotation"`
		EditionAnnotation            string      `json:"editionAnnotation"`
		CardConent                   string      `json:"cardConent"`
		BoardWidth                   string      `json:"boardWidth"`
		FormatAnnotation             string      `json:"formatAnnotation"`
		BookSize                     string      `json:"bookSize"`
		Function                     string      `json:"function"`
		BindingAnnotation            string      `json:"bindingAnnotation"`
		Fittings                     string      `json:"fittings"`
		XuBa                         string      `json:"xuBa"`
		Inscription                  string      `json:"inscription"`
		Details                      string      `json:"details"`
		QianYin                      string      `json:"qianYin"`
		Accessory                    string      `json:"accessory"`
		RegistrationNo               string      `json:"registrationNo"`
		MoneyTransfer                string      `json:"moneyTransfer"`
		BookSourceAnnotation         string      `json:"bookSourceAnnotation"`
		RepaireHistory               string      `json:"repaireHistory"`
		GradingPeople                string      `json:"gradingPeople"`
		GradingLevel                 string      `json:"gradingLevel"`
		GradingDate                  string      `json:"gradingDate"`
		AncientGrading               string      `json:"ancientGrading"`
		GradingReason                string      `json:"gradingReason"`
		GradingAnnotation            string      `json:"gradingAnnotation"`
		BookDamageLevel              string      `json:"bookDamageLevel"`
		RepairSuggest                string      `json:"repairSuggest"`
		Damage                       string      `json:"damage"`
		PeopleFee                    string      `json:"peopleFee"`
		TotalLossNumber              string      `json:"totalLossNumber"`
		DamageAnnotation             string      `json:"damageAnnotation"`
		RatingBooks                  string      `json:"ratingBooks"`
		DamageBooks                  string      `json:"damageBooks"`
		Creator                      interface{} `json:"creator"`
		CreateTime                   interface{} `json:"createTime"`
		Updator                      interface{} `json:"updator"`
		UpdateTime                   interface{} `json:"updateTime"`
		IsDeleted                    int         `json:"isDeleted"`
	} `json:"anbCatalogDetail"`
	IsHaveFasc  int         `json:"isHaveFasc"`
	IsHaveDire  interface{} `json:"isHaveDire"`
	IsHaveOcr   int         `json:"isHaveOcr"`
	Subtitles   interface{} `json:"subtitles"`
	IsPrivilege interface{} `json:"isPrivilege"`
	ShowName    interface{} `json:"showName"`
	Num         interface{} `json:"num"`
	IsPrint     interface{} `json:"isPrint"`
	IsFamily    interface{} `json:"isFamily"`
	OrgId       string      `json:"orgId"`
}

type BookCatalog struct {
	Records []struct {
		DirectoryId string      `json:"directoryId"`
		FascicleId  string      `json:"fascicleId"`
		CatalogId   string      `json:"catalogId"`
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
		PageId      string      `json:"pageId"`
		GradeId     string      `json:"gradeId"`
		Region      string      `json:"region"`
		Sort        int         `json:"sort"`
		Creator     interface{} `json:"creator"`
		CreateTime  string      `json:"createTime"`
		Updator     interface{} `json:"updator"`
		UpdateTime  interface{} `json:"updateTime"`
		IsDeleted   int         `json:"isDeleted"`
	} `json:"records"`
	Total       int  `json:"total"`
	Size        int  `json:"size"`
	Current     int  `json:"current"`
	SearchCount bool `json:"searchCount"`
	Pages       int  `json:"pages"`
}

type BookFascicle struct {
	FascicleId   string      `json:"fascicleId"`
	CatalogId    string      `json:"catalogId"`
	Name         string      `json:"name"`
	Introduction interface{} `json:"introduction"`
	GradeId      string      `json:"gradeId"`
	Sort         int         `json:"sort"`
	Creator      interface{} `json:"creator"`
	CreateTime   string      `json:"createTime"`
	Updator      interface{} `json:"updator"`
	UpdateTime   string      `json:"updateTime"`
	IsDeleted    int         `json:"isDeleted"`
	FilePath     interface{} `json:"filePath"`
	ImageCount   interface{} `json:"imageCount"`
}

type PageImage struct {
	Records     []ImageRecord `json:"records"`
	Total       int           `json:"total"`
	Size        int           `json:"size"`
	Current     int           `json:"current"`
	SearchCount bool          `json:"searchCount"`
	Pages       int           `json:"pages"`
}
type ImageRecord struct {
	ImageId     string      `json:"imageId"`
	ImageName   string      `json:"imageName"`
	DirectoryId string      `json:"directoryId"`
	FascicleId  string      `json:"fascicleId"`
	CatalogId   string      `json:"catalogId"`
	Sort        int         `json:"sort"`
	Type        int         `json:"type"`
	IsParse     interface{} `json:"isParse"`
	Description interface{} `json:"description"`
	Creator     string      `json:"creator"`
	CreateTime  string      `json:"createTime"`
	Updator     string      `json:"updator"`
	UpdateTime  string      `json:"updateTime"`
	IsDeleted   int         `json:"isDeleted"`
	OcrInfo     interface{} `json:"ocrInfo"`
	File        interface{} `json:"file"`
}

// 基本信息
type ResponseDetail struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data BookDetail `json:"data"`
}

// 几册
type ResponseCatalog struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data BookCatalog `json:"data"`
}

// 分卷
type ResponseFascicle struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data []BookFascicle `json:"data"`
}

// 页面
type ResponsePage struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data PageImage `json:"data"`
}

// 图像
type ResponseObject struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	File []struct {
		FileName    string `json:"fileName"`
		FileSuffix  string `json:"fileSuffix"`
		FilePath    string `json:"filePath"`
		UpdateTime  string `json:"updateTime"`
		CreateTime  string `json:"createTime"`
		FileSize    int    `json:"fileSize"`
		FileOldname string `json:"fileOldname"`
	} `json:"file"`
}

type ResponseFile struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OcrInfo struct {
			OcrId        string      `json:"ocrId"`
			ImageId      string      `json:"imageId"`
			CatalogId    string      `json:"catalogId"`
			OcrText      string      `json:"ocrText"`
			TradeOcrText interface{} `json:"tradeOcrText"`
			OcrJson      string      `json:"ocrJson"`
			Width        int         `json:"width"`
			Height       int         `json:"height"`
			FontWidth    interface{} `json:"fontWidth"`
			CutLevel     interface{} `json:"cutLevel"`
			BitwiseNot   int         `json:"bitwiseNot"`
			Creator      string      `json:"creator"`
			CreateTime   string      `json:"createTime"`
			Updator      interface{} `json:"updator"`
			UpdateTime   interface{} `json:"updateTime"`
			IsDeleted    int         `json:"isDeleted"`
			FilePath     []struct {
				FileName    string `json:"fileName"`
				FileSuffix  string `json:"fileSuffix"`
				FilePath    string `json:"filePath"`
				UpdateTime  string `json:"updateTime"`
				Sort        string `json:"sort"`
				CreateTime  string `json:"createTime"`
				FileSize    int    `json:"fileSize"`
				FileOldname string `json:"fileOldname"`
				FileInfoId  string `json:"fileInfoId"`
			} `json:"filePath"`
			FascicleId    interface{} `json:"fascicleId"`
			FascicleName  interface{} `json:"fascicleName"`
			DirectoryId   interface{} `json:"directoryId"`
			DirectoryName interface{} `json:"directoryName"`
			CatalogName   interface{} `json:"catalogName"`
		} `json:"ocrInfo"`
		File []struct {
			FileName    string `json:"fileName"`
			FileSuffix  string `json:"fileSuffix"`
			FilePath    string `json:"filePath"`
			UpdateTime  string `json:"updateTime"`
			CreateTime  string `json:"createTime"`
			FileSize    int    `json:"fileSize"`
			FileOldname string `json:"fileOldname"`
		} `json:"file"`
	} `json:"data"`
}

type Canvases struct {
	ImgUrls    []string
	ImgOcrUrls []string
	Size       int
}

type ResponseVolume struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []Volume `json:"data"`
}

type Volume struct {
	FascicleId   string      `json:"fascicleId"`
	CatalogId    string      `json:"catalogId"`
	Name         string      `json:"name"`
	Introduction interface{} `json:"introduction"`
	GradeId      string      `json:"gradeId"`
	Sort         int         `json:"sort"`
	Creator      interface{} `json:"creator"`
	CreateTime   string      `json:"createTime"`
	Updator      interface{} `json:"updator"`
	UpdateTime   string      `json:"updateTime"`
	IsDeleted    int         `json:"isDeleted"`
	FilePath     interface{} `json:"filePath"`
	ImageCount   interface{} `json:"imageCount"`
}

type Parts map[string][]ImageRecord
