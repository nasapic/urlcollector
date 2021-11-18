package jsonapi

import (
	"gitlab.com/nasapic/base"
)

type (
	Endpoint struct {
		*base.Endpoint
	}
)

func NewEndpoint(name string) *Endpoint {
	return &Endpoint{
		Endpoint: base.NewEndpoint(name),
	}
}
