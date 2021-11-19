package service

import "gitlab.com/nasapic/base"

type (
	Collector struct {
		base.Worker
	}
)

func NewCollector(name string, log base.Logger) *Collector {
	return &Collector{
		Worker: base.NewWorker(name, log),
	}
}
