/*
Package nasa provides a basic implementation of collector.Collector.
For the sake of simplicity at the moment is implemented within the service libs
but eventually they could be moved to a custom module in order to facilitate
their reuse.
*/
package nasa

import (
	"time"

	"gitlab.com/nasapic/urlcollector/internal/lib/collector"
)

type (
	API struct {
		opts Options
		from time.Time
		to   time.Time
	}

	Result struct {
		list collector.URLS
	}
)

type (
	Options struct {
		APIKey string
	}
)

func NewAPI(opts Options) *API {
	return &API{
		opts: opts,
	}
}

func (api API) GetBetweenDates(from, to time.Time) collector.Result {
	return Result{}
}

func (r Result) GetList() collector.URLS {
	return r.list
}
