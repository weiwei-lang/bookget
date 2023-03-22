package router

import (
	"bookget/app"
	"bookget/site/USA/harvard"
	"bookget/site/USA/loc"
	"bookget/site/USA/princeton"
)

type Harvard struct {
}

func (p Harvard) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		harvard.Init(i+1, s)
	}
	return nil, nil
}

type Hathitrust struct {
}

func (p Hathitrust) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		var hathitrust app.Hathitrust
		hathitrust.Init(i+1, s)
	}
	return nil, nil
}

type Princeton struct {
}

func (p Princeton) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		princeton.Init(i+1, s)
	}
	return nil, nil
}

type UsLoc struct {
}

func (p UsLoc) getRouterInit(sUrl []string) (map[string]interface{}, error) {
	for i, s := range sUrl {
		loc.Init(i+1, s)
	}
	return nil, nil
}
