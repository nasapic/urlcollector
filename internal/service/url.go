package service

import (
	"gitlab.com/nasapic/base"
	"gitlab.com/nasapic/urlcollector/internal/transport"
	"gitlab.com/nasapic/urlcollector/pkg/collector"
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

func (urlSvc *URLService) GetBetweenDates(sReq *transport.SearchRequest) (sRes *transport.SearchResponse, err error) {
	// NOTE: If there are no validation errors we can use the original string
	// representation of from and to dates and avoid converting them again from
	// time.Time representation used for validations.
	sReq.Validate()
	if sReq.HasErrors() || true {
		urlSvc.Log().Debug("URLService GetBetweenDates", "searchRequest", sReq, "errors", sReq.Error())
		return sRes, sReq.Error()
	}

	result, err := urlSvc.CollectorAPI.GetBetweenDates(sReq.FromDate(), sReq.ToDate())

	urlSvc.Log().Debug("GetBetweenDates", "result", result)

	return &transport.SearchResponse{}, nil
}
