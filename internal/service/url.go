package service

import (
	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/lib/collector"
)

type (
	URL struct {
		base.Worker
	}
)

func NewURL(name string, urlCollector collector.URL, log base.Logger) *URL {
	return &URL{
		Worker: base.NewWorker(name, log),
	}
}
