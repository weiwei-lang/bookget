package router

import (
	"bookget/app"
	"bookget/site/China/idp"
	"bookget/site/Universal"
)

type NormalIIIF struct{}

func (p NormalIIIF) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		Universal.AutoDetectManifest(i+1, s)
	}
	return nil, nil
}

type NormalHttp struct{}

func (p NormalHttp) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		Universal.StartDownload(i+1, s)
	}
	return nil, nil
}

type IDP struct{}

func (p IDP) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		idp.Init(i+1, s)
	}
	return nil, nil
}

type KyudbSnu struct{}

func (p KyudbSnu) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		n := i + 1
		var kyudbsnu app.KyudbSnu
		kyudbsnu.Init(n, s)
	}
	return nil, nil
}
