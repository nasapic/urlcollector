package jsonapi

import (
	"net/http"

	"gitlab.com/nasapic/urlcollector/internal/transport"
)

// SearchURLs endpoint
func (ep *Endpoint) SearchURLs(w http.ResponseWriter, r *http.Request) {
	// Transport
	searchReq := transport.NewSearchRequest(r)

	ep.Log().Debug("SearchURLs", "searchReq", searchReq)

	ep.URLService.GetBetweenDates(searchReq)

	panic("not fully implemented")
}
