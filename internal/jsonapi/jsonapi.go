package jsonapi

import (
	"gitlab.com/nasapic/base"
)

type (
	Endpoint struct {
		*base.Endpoint
	}
)

func NewEndpoint(name string, log base.Logger) *Endpoint {
	return &Endpoint{
		Endpoint: base.NewEndpoint(name, log),
	}
}
