package jsonapi

import (
	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/service"
)

type (
	Endpoint struct {
		*base.Endpoint

		URLService *service.URLService
		Opts       *Options
	}

	Options struct {
		MaxConcurrent int
	}
)

func NewEndpoint(name string, urlSvc *service.URLService, opts *Options, log base.Logger) *Endpoint {
	return &Endpoint{
		Endpoint:   base.NewEndpoint(name, log),
		URLService: urlSvc,
		Opts:       opts,
	}
}
