package httpx

import (
	"github.com/behavioral-ai/core/httpx"
	"net/http"
)

var (
	cache = newCache()
)

// Exchange - HTTP exchange function
func Exchange(w http.ResponseWriter, r *http.Request) {
	var resp *http.Response

	if r.Method == http.MethodGet {
		resp = cache.get(r)
	} else {
		resp = cache.put(r)
	}
	httpx.WriteResponse(w, resp.Header, resp.StatusCode, resp.Body, r.Header)
}
