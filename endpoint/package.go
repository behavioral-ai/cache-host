package endpoint

import (
	"github.com/behavioral-ai/core/httpx"
	http2 "net/http"
)

var (
	cache = httpx.NewResponseCache()
)

// Exchange - HTTP exchange function
func Exchange(w http2.ResponseWriter, r *http2.Request) {
	var (
		resp *http2.Response
	)

	switch r.Method {
	case http2.MethodGet:
		resp = cache.Get(r.URL.String())
		if resp.StatusCode == http2.StatusNotFound {
			w.WriteHeader(resp.StatusCode)
			return
		}
		httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, nil)
	case http2.MethodPut:
		cache.Put(r.URL.String(), httpx.CreateResponse(r))
		w.WriteHeader(http2.StatusOK)
	default:
		w.WriteHeader(http2.StatusMethodNotAllowed)
		return
	}
}
