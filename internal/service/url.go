package service

import (
	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/lib/collector"
)

type (
	URLService struct {
		base.Worker
		CollectorAPI collector.API
	}
)

func NewURLService(name string, collectorAPI collector.API, log base.Logger) *URLService {
	return &URLService{
		Worker:       base.NewWorker(name, log),
		CollectorAPI: collectorAPI,
	}
}
