package endpoint

import (
	"github.com/behavioral-ai/core/httpx"
	"net/http"
)

var (
	cache = httpx.NewResponseCache()
)

// Exchange - HTTP exchange function
func Exchange(w http.ResponseWriter, r *http.Request) {
	var (
		resp *http.Response
	)

	switch r.Method {
	case http.MethodGet:
		resp = cache.Get(r.URL.String())
		if resp.StatusCode == http.StatusNotFound {
			w.WriteHeader(resp.StatusCode)
			return
		}
		httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, nil)
	case http.MethodPut:
		cache.Put(r.URL.String(), httpx.CreateResponse(r))
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
