package jsonapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/nasapic/urlcollector/internal/transport"
)

// SearchURLs endpoint
func (ep *Endpoint) SearchURLs(w http.ResponseWriter, r *http.Request) {
	// Transport
	searchReq := transport.NewSearchRequest(r)

	ep.Log().Debug("Endpoint SearchURLS", "searchReq", searchReq)

	// Service call
	searchRes, err := ep.URLService.GetBetweenDates(searchReq)
	if err != nil {
		// Error response
		return
	}

	// OK response marshalling
	jsonRes, err := json.MarshalIndent(searchRes, "", "  ")
	if err != nil {
		// Error response
		return
	}

	ep.Log().Debug("Endpoint SearchURLS", "searchRes", searchRes)

	fmt.Fprintf(w, string(jsonRes))
}
