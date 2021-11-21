package transport

import (
	"net/http"
	"time"
)

type (
	SearchRequest struct {
		From     string `json:"from"`
		To       string `json:"to"`
		fromDate time.Time
		toDate   time.Time
		Errors   map[string]error
	}

	SearchResponse struct {
		URLS []string `json:"urls"`
	}
)

const (
	dateFormat string = "2006-01-02"
)

func NewSearchRequest(r *http.Request) *SearchRequest {
	sr := SearchRequest{}

	sr.From = r.URL.Query().Get("from")
	sr.To = r.URL.Query().Get("to")

	sr.ParseInput()

	return &sr
}

func (sr *SearchRequest) FromDate() time.Time {
	return sr.fromDate
}

func (sr *SearchRequest) ToDate() time.Time {
	return sr.toDate
}

func (sr *SearchRequest) ParseInput() (hasErrors bool) {
	sr.clearErrors()

	date, err := toTime(sr.From)
	if err != nil {
		sr.Errors["from"] = err
		hasErrors = hasErrors && true
	}

	sr.fromDate = date

	date, err = toTime(sr.To)
	if err != nil {
		sr.Errors["to"] = err
		hasErrors = hasErrors && true
	}

	sr.toDate = date

	return hasErrors
}

func (sr *SearchRequest) Validate() (hasErrors bool) {
	// At the moment it only validates parsing errors.
	// Implement more complex date validations.
	return len(sr.Errors) > 0
}

func (sr *SearchRequest) clearErrors() {
	sr.Errors = make(map[string]error)
}

func toTime(dateString string) (date time.Time, err error) {
	date, err = time.Parse(dateFormat, dateString)

	if err != nil {
		return date, err
	}

	return date, nil
}

func toTimeString(date time.Time) (dateString string) {
	return date.Format(dateFormat)
}
