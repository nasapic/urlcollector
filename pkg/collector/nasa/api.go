/*
Package nasa provides a basic implementation of collector.Collector.
For the sake of simplicity at the moment is implemented within the service libs
but eventually they could be moved to a custom module in order to facilitate
their reuse.

Sample JSON response
{
	"date":"2020-01-01",
	"explanation":"Why is Betelgeuse fading?  No one knows.  Betelgeuse, one of the brightest and most recognized stars in the night sky, is only half as bright as it used to be only five months ago.  Such variability is likely just  normal behavior for this famously variable supergiant, but the recent dimming has rekindled discussion on how long it may be before Betelgeuse does go supernova.  Known for its red color, Betelgeuse is one of the few stars to be resolved by modern telescopes, although only barely.  The featured artist's illustration imagines how Betelgeuse might look up close. Betelgeuse is thought to have a complex and tumultuous surface that frequently throws impressive flares.  Were it to replace the Sun (not recommended), its surface would extend out near the orbit of Jupiter, while gas plumes would bubble out past Neptune.  Since Betelgeuse is about 700 light years away, its eventual supernova will not endanger life on Earth even though its brightness may rival that of a full Moon.  Astronomers -- both amateur and professional -- will surely continue to monitor Betelgeuse as this new decade unfolds.    Free Presentation: APOD Editor to show best astronomy images of 2019 -- and the decade -- in NYC on January 3",
	"hdurl":"https://apod.nasa.gov/apod/image/2001/BetelgeuseImagined_EsoCalcada_2662.jpg",
	"media_type":"image",
	"service_version":"v1",
	"title":"Betelgeuse Imagined",
	"url":"https://apod.nasa.gov/apod/image/2001/BetelgeuseImagined_EsoCalcada_960.jpg"
}
*/
package nasa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/base"
	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/urlcollector/pkg/collector"
)

type (
	API struct {
		opts   Options
		log    base.Logger
		from   time.Time
		to     time.Time
		client *httpClient
	}

	Result struct {
		list []string
	}
)

type (
	Options struct {
		APIKey        string
		TimeoutInSecs int
		MaxConcurrent int
	}
)

type (
	APIResponse struct {
		Copyright      string `json:"copyright"`
		Date           string `json:"date"`
		Explanation    string `json:"explanation"`
		HDURL          string `json:"hdurl"`
		MediaType      string `json:"media_type"`
		ServiceVersion string `json:"service_version"`
		Title          string `json:"title"`
		URL            string `json:"url"`
	}
)

type (
	Job struct {
		ID   uuid.UUID
		Date time.Time
	}

	JobResult struct {
		ID     uuid.UUID
		Date   time.Time
		Result *APIResponse
		Error  error
	}
)

type (
	// NOTE: A future update could include retries with exponential back off.
	httpClient struct {
		*http.Client
	}
)

const (
	baseURL                   = "https://api.nasa.gov/planetary/apod"
	dateFormat         string = "2006-01-02"
	chanClearanceIndex        = 15
)

func NewAPI(opts Options, log base.Logger) *API {
	return &API{
		opts: opts,
		log:  log,
		client: &httpClient{
			&http.Client{
				Timeout: time.Second * time.Duration(opts.TimeoutInSecs),
			},
		},
	}
}

func (api *API) GetBetweenDates(from, to time.Time) (cr collector.Result, err error) {
	dates := []time.Time{}

	if to.Before(from) {
		return Result{}, errors.New("invalid date range")
	}

	func() error {
		for dateElement := rangeDates(from, to); ; {
			date := dateElement()
			if date.IsZero() {
				return nil
			}

			dates = append(dates, date)
		}
	}()

	chanSize := chanSafeSize(len(dates))
	max := api.opts.MaxConcurrent

	jobsChan := make(chan Job, chanSize)
	resultsChan := make(chan JobResult, chanSize)

	// Create API callers
	for c := 1; c <= max; c++ {
		id := uuid.New()
		go api.caller(id, jobsChan, resultsChan)
	}

	// Send Jobs to API callers
	jobIDs := []uuid.UUID{}

	for _, date := range dates {
		id := uuid.New()

		job := Job{
			ID:   id,
			Date: date,
		}

		select {
		case jobsChan <- job:
			jobIDs = append(jobIDs, id)
		default:
			err := errors.New("max allowed concurrent requests reached")
			api.Log().Error(err, "Max allowed concurrent requests reached", "max-value", max)
		}
	}
	close(jobsChan)

	// Collect results
	list := []string{}
	for range jobIDs {

		res := <-resultsChan
		if res.Error != nil {
			api.Log().Error(err, "API call error", "error", res.Error, "job-id", res.ID, "picture-date", res.Date)
			continue
		}

		list = append(list, res.Result.URL)
	}

	cr = &Result{
		list: list,
	}

	return cr, nil
}

func (api *API) getByDate(date time.Time) (ar *APIResponse, err error) {
	if date.IsZero() {
		return ar, errors.New("Invalid date")
	}

	dateStr := toDateString(date)

	url := fmt.Sprintf("%s/?api_key=%s&date=%s", baseURL, api.opts.APIKey, dateStr)

	api.Log().Debug("Getting URL", "url", url)

	res, err := api.client.get(url)
	if err != nil {
		return ar, fmt.Errorf("Error retrievieng picture URL: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return ar, fmt.Errorf("Non OK response: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ar, err
	}

	if res.StatusCode == http.StatusOK {
		json.Unmarshal(body, &ar)
	}

	return ar, nil
}

func (r Result) GetList() []string {
	return r.list
}

func (api API) caller(id uuid.UUID, jobs <-chan Job, results chan<- JobResult) {
	for j := range jobs {
		ar, err := api.getByDate(j.Date)
		results <- JobResult{
			ID:     id,
			Result: ar,
			Error:  err,
		}
	}
}

func (api *API) Log() base.Logger {
	return api.log
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

func chanSafeSize(toProcQty int) int {
	return toProcQty * (1 + chanClearanceIndex/100)
}

func toDateString(date time.Time) (dateString string) {
	return date.Format(dateFormat)
}
