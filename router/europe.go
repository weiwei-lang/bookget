package router

import (
	"bookget/site/Europe/bavaria"
	"bookget/site/Europe/berlin"
	"bookget/site/Europe/bluk"
	"bookget/site/Europe/oxacuk"
	"bookget/site/USA/familysearch"
	"bookget/site/USA/stanford"
)

type OxacUk struct{}

func (p OxacUk) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		oxacuk.Init(i+1, s)
	}
	return nil, nil
}

type DigitalBerlin struct{}

func (p DigitalBerlin) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		berlin.Init(i+1, s)
	}
	return nil, nil
}

type BlUk struct{}

func (p BlUk) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		bluk.Init(i+1, s)
	}
	return nil, nil
}

type OstasienBavaria struct{}

func (p OstasienBavaria) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		bavaria.Init(i+1, s)
	}
	return nil, nil
}

type SearchworksStanford struct{}

func (p SearchworksStanford) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		stanford.Init(i+1, s)
	}
	return nil, nil
}

type FamilySearch struct{}

func (p FamilySearch) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		familysearch.Init(i+1, s)
	}
	return nil, nil
}
