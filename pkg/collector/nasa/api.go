/*
Package nasa provides a basic implementation of collector.Collector.
For the sake of simplicity at the moment is implemented within the service libs
but eventually they could be moved to a custom module in order to facilitate
their reuse.
*/
package nasa

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/nasapic/urlcollector/internal/transport"
	"gitlab.com/nasapic/urlcollector/pkg/collector"
)

type (
	API struct {
		opts   Options
		from   time.Time
		to     time.Time
		client *httpClient
	}

	Result struct {
		list collector.URLS
	}
)

type (
	Options struct {
		APIKey        string
		MaxRequests   int
		TimeoutInSecs int
	}
)

type (
	// NOTE: A future update could include retries with exponential back off.
	httpClient struct {
		*http.Client
	}
)

const (
	baseURL           = "https://api.nasa.gov/planetary/apod"
	dateFormat string = "2006-01-02"
)

func NewAPI(opts Options) *API {
	return &API{
		opts: opts,
		client: &httpClient{
			&http.Client{
				Timeout: time.Second * time.Duration(opts.TimeoutInSecs),
			},
		},
	}
}

func (api API) GetBetweenDates(from, to time.Time) collector.Result {
	for dateElement := rangeDates(from, to); ; {
		date := dateElement()
		if date.IsZero() {
			break
		}

		fmt.Println(toDateString(date))
	}

	return Result{}
}

func (api API) getByDate(sReq *transport.SearchRequest) (picURL string, err error) {
	url := fmt.Sprintf("%s/api_key=%s&date=", baseURL, api.opts.APIKey)

	_, err = api.client.get(url)
	if err != nil {
		return picURL, err
	}

	return "", nil
}

func (r Result) GetList() collector.URLS {
	return r.list
}

// HTTP Client implementation
func (hc *httpClient) get(url string) (res *http.Response, err error) {
	res, err = hc.Client.Get(url)
	if err != nil {
		return res, err
	}

	return res, err
}

// Helpers
func rangeDates(from, to time.Time) (rdFunc func() time.Time) {
	year, month, day := from.Date()
	from = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	year, month, day = to.Date()
	to = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return dateRangeFunc(from, to)
}

func dateRangeFunc(from, to time.Time) (drFunc func() time.Time) {
	return func() time.Time {
		if from.After(to) {
			return time.Time{}
		}

		date := from
		from = from.AddDate(0, 0, 1)
		return date
	}
}

func toDateString(date time.Time) (dateString string) {
	return date.Format(dateFormat)
}
