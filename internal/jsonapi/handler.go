package jsonapi

import (
	"fmt"
	"net/http"

	"gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/urlcollector/internal/transport"
)

// SearchURLs endpoint
func (ep *Endpoint) SearchURLs(w http.ResponseWriter, r *http.Request) {
	searchReq := transport.NewSearchRequest(r)

	ep.Log().Debug("Endpoint SearchURLS", "searchReq", searchReq)

	res, err := ep.URLService.GetBetweenDates(searchReq)
	if err != nil {
		ep.sendErrorResponse(w, r, "Cannot get pictures", err)
		return
	}

	jsonStr, err := res.Marshall()
	if err != nil {
		ep.sendErrorResponse(w, r, "Error marshalling search response", err)
		return
	}

	ep.Log().Debug("Endpoint SearchURLS", "search-result", res)

	fmt.Fprintf(w, jsonStr)
}

// Helpers

func (ep *Endpoint) sendErrorResponse(w http.ResponseWriter, r *http.Request, message string, err error) {
	errRes := transport.ErrorResponse{Error: err.Error()}

	errStr, mErr := errRes.Marshall()
	if mErr != nil {
		ep.Log().Error(mErr, "Error marshalling response", "original-error", err)
		return
	}

	ep.Log().Error(err, message)

	fmt.Fprintf(w, errStr)
}
