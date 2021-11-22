package transport

import (
	"errors"
	"net/http"
	"strings"
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
	sr.ClearErrors()

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
	// At the moment it only validates parsing errors
	// and do a simple dates checks.
	// NOTE: It may be a good idea to get rid of previous parsing errors.
	sr.ClearErrors()

	if sr.fromDate.IsZero() {
		sr.Errors["from"] = errors.New("invalid 'from' date")
	}

	if sr.toDate.IsZero() {
		sr.Errors["to"] = errors.New("invalid 'to' date")
	}

	if sr.toDate.Before(sr.fromDate) {
		sr.Errors["from"] = errors.New("'from' date should be prior to 'to' date")
		sr.Errors["to"] = errors.New("'to' date should be after 'from' date")
	}

	return len(sr.Errors) > 0
}

func (sr *SearchRequest) Error() error {
	var sb strings.Builder

	last := len(sr.Errors) - 1
	if last < 0 {
		return nil
	}

	pos := 0
	sep := ""
	for _, e := range sr.Errors {
		sb.WriteString(sep)
		sb.WriteString(e.Error())
		if pos == last {
			sep = ""
		} else {
			sep = ", "
		}

		pos++
	}

	return errors.New(sb.String())
}

func (sr *SearchRequest) HasErrors() (hasErrors bool) {
	return len(sr.Errors) > 0
}

func (sr *SearchRequest) ClearErrors() {
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
