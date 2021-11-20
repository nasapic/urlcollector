package jsonapi

import (
	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/service"
)

type (
	Endpoint struct {
		*base.Endpoint
		URLService *service.URLService
	}
)

func NewEndpoint(name string, urlSvc *service.URLService, log base.Logger) *Endpoint {
	return &Endpoint{
		Endpoint:   base.NewEndpoint(name, log),
		URLService: urlSvc,
	}
}
