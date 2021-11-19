package jsonapi

import (
	"net/http"
)

// SearchURLs endpoint
func (ep *Endpoint) SearchURLs(w http.ResponseWriter, r *http.Request) {

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	ep.Log().Debug("SearchURLs", "from", from, "to", to)

	panic("not implemented")
}
