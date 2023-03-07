package router

import (
	"bookget/site/China/bjdpm"
	"bookget/site/China/cuhk"
	"bookget/site/China/gzlib"
	"bookget/site/China/jslib"
	"bookget/site/China/luoyang"
	"bookget/site/China/ncpssd"
	"bookget/site/China/npmtw"
	"bookget/site/China/ouroots"
	"bookget/site/China/rbkdocnpmtw"
	"bookget/site/China/sclib"
	"bookget/site/China/sdutcm"
	"bookget/site/China/szlib"
	"bookget/site/China/tianyige"
	"bookget/site/China/twnlc"
	"bookget/site/China/usthk"
	"bookget/site/China/wzlib"
	"bookget/site/China/ynutcm"
)
import "bookget/site/China/nlc"

type ChinaNcl struct{}

func (p ChinaNcl) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		nlc.Init(i+1, s)
	}
	return nil, nil
}

type RbookNcl struct{}

func (p RbookNcl) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		twnlc.Init(i+1, s)
	}
	return nil, nil
}

type RbkdocNpmTw struct{}

func (p RbkdocNpmTw) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		rbkdocnpmtw.Init(i+1, s)
	}
	return nil, nil
}

type DigitalarchiveNpmTw struct{}

func (p DigitalarchiveNpmTw) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		npmtw.Init(i+1, s)
	}
	return nil, nil
}

type CuHk struct{}

func (p CuHk) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		cuhk.Init(i+1, s)
	}
	return nil, nil
}

type UstHk struct{}

func (p UstHk) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		usthk.Init(i+1, s)
	}
	return nil, nil
}

type LuoYang struct{}

func (p LuoYang) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		luoyang.Init(i+1, s)
	}
	return nil, nil
}

type OyjyWzlib struct{}

func (p OyjyWzlib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		wzlib.Init(i+1, s)
	}
	return nil, nil
}

type YunSzlib struct{}

func (p YunSzlib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		szlib.Init(i+1, s)
	}
	return nil, nil
}

type GzddGzlib struct{}

func (p GzddGzlib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		gzlib.Init(i+1, s)
	}
	return nil, nil
}

type TianYiGeLib struct{}

func (p TianYiGeLib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		tianyige.Init(i+1, s)
	}
	return nil, nil
}

type GujiSclib struct{}

func (p GujiSclib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		sclib.Init(i+1, s)
	}
	return nil, nil
}

type GuijiJslib struct{}

func (p GuijiJslib) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		jslib.Init(i+1, s)
	}
	return nil, nil
}

type MinghuajiBjDpm struct{}

func (p MinghuajiBjDpm) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		bjdpm.Init(i+1, s)
	}
	return nil, nil
}

type OurootsNlc struct{}

func (p OurootsNlc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		ouroots.Init(i+1, s)
	}
	return nil, nil
}

type Ncpssd struct{}

func (p Ncpssd) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		ncpssd.Init(i+1, s)
	}
	return nil, nil
}

type GujiYnutcm struct{}

func (p GujiYnutcm) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		ynutcm.Init(i+1, s)
	}
	return nil, nil
}

type Sdutcm struct{}

func (p Sdutcm) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		sdutcm.Init(i+1, s)
	}
	return nil, nil
}
