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
	URL struct {
		from time.Time
		to   time.Time
	}

	Result struct {
		list collector.URLS
	}
)

func (url *URL) GetBetweenDates(from, to time.Time) Result {
	return Result{}
}

func (r *Result) GetList() collector.URLS {
	return r.list
}
