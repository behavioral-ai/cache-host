package http

import (
	"github.com/behavioral-ai/core/httpx"
	http2 "net/http"
)

var (
	cache = newCache()
)

// Exchange - HTTP exchange function
func Exchange(w http2.ResponseWriter, r *http2.Request) {
	var (
		resp   *http2.Response
		status int
	)

	switch r.Method {
	case http2.MethodGet:
		resp, status = cache.get(r.URL.String())
		if status == http2.StatusNotFound {
			w.WriteHeader(status)
			return
		}
		httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, nil)
	case http2.MethodPut:
		cache.put(r.URL.String(), &http2.Response{StatusCode: http2.StatusOK, Header: httpx.CloneHeader(r.Header), Body: r.Body})
		w.WriteHeader(http2.StatusOK)
	default:
		w.WriteHeader(http2.StatusMethodNotAllowed)
		return
	}
}
