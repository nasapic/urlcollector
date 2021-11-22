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

func (urlSvc *URLService) GetBetweenDates(searchReq *transport.SearchRequest) (searchRes *transport.SearchResponse, err error) {
	// NOTE: If there are no validation errors we can use the original string
	// representation of from and to dates and avoid converting them again from
	// time.Time representation used for transport validations.
	searchReq.Validate()

	urlSvc.Log().Debug("URLService GetBetweenDates", "searchRequest", searchReq)

	if searchReq.HasErrors() {
		urlSvc.Log().Debug("URLService GetBetweenDates", "errors", searchReq.Error())
		return searchRes, searchReq.Error()
	}

	result, err := urlSvc.CollectorAPI.GetBetweenDates(searchReq.FromDate(), searchReq.ToDate())
	if err != nil {
		urlSvc.Log().Debug("URLService GetBetweenDates", "errors", err)
		return searchRes, err
	}

	searchRes = &transport.SearchResponse{
		URLS: result.GetList(),
	}

	urlSvc.Log().Debug("URLService GetBetweenDates", "searchRes", searchRes)

	return searchRes, nil
}
