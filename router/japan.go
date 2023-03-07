package router

import (
	"bookget/site/Japan/emuseum"
	"bookget/site/Japan/gprime"
	"bookget/site/Japan/kanjikyoto"
	"bookget/site/Japan/keio"
	"bookget/site/Japan/khirin"
	"bookget/site/Japan/kokusho"
	"bookget/site/Japan/kyoto"
	"bookget/site/Japan/national"
	"bookget/site/Japan/ndl"
	"bookget/site/Japan/niiac"
	"bookget/site/Japan/utokyo"
	"bookget/site/Japan/waseda"
	"bookget/site/Japan/yonezawa"
)

type RmdaKyoto struct{}

func (p RmdaKyoto) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		kyoto.Init(i+1, s)
	}
	return nil, nil
}

type NdlGo struct{}

func (p NdlGo) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		ndl.Init(i+1, s)
	}
	return nil, nil
}

type EmuseumNich struct{}

func (p EmuseumNich) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		emuseum.Init(i+1, s)
	}
	return nil, nil
}

type SidoKeio struct{}

func (p SidoKeio) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		keio.Init(i+1, s)
	}
	return nil, nil
}

type ShanbenuTokyo struct{}

func (p ShanbenuTokyo) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		utokyo.Init(i+1, s)
	}
	return nil, nil
}

type ArchivesGo struct{}

func (p ArchivesGo) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		national.Init(i+1, s)
	}
	return nil, nil
}

type DsrNiiAc struct{}

func (p DsrNiiAc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		niiac.Init(i+1, s)
	}
	return nil, nil
}

type WulWasedaAc struct{}

func (p WulWasedaAc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		waseda.Init(i+1, s)
	}
	return nil, nil
}

type KokushoNijlAc struct{}

func (p KokushoNijlAc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		kokusho.Init(i+1, s)
	}
	return nil, nil
}

type KanjiZinbunKyotouAc struct{}

func (p KanjiZinbunKyotouAc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		kanjikyoto.Init(i+1, s)
	}
	return nil, nil
}

type ElibGprime struct{}

func (p ElibGprime) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		gprime.Init(i+1, s)
	}
	return nil, nil
}

type KhirinRekihaku struct{}

func (p KhirinRekihaku) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		khirin.Init(i+1, s)
	}
	return nil, nil
}

type LibYonezawa struct{}

func (p LibYonezawa) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		yonezawa.Init(i+1, s)
	}
	return nil, nil
}
