package httpx

import (
	"net/http"
	"sync"
)

var (
	notFoundResponse = &http.Response{StatusCode: http.StatusNotFound}
	okResponse       = &http.Response{StatusCode: http.StatusOK}
)

type contentT struct {
	m *sync.Map
}

func newCache() *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	return c
}

func (c *contentT) get(req *http.Request) *http.Response {
	value, ok := c.m.Load(req.URL.String())
	if !ok {
		return notFoundResponse
	}
	if r, ok1 := value.(*http.Response); ok1 {
		return r
	}
	return notFoundResponse
}

func (c *contentT) put(req *http.Request) *http.Response {
	resp := &http.Response{StatusCode: http.StatusOK, Header: req.Header, Body: req.Body}
	c.m.Store(req.URL.String(), resp)
	return okResponse
}
