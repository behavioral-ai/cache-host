package http

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

func (c *contentT) get(key string) (*http.Response, int) {
	value, ok := c.m.Load(key)
	if !ok {
		return nil, http.StatusNotFound
	}
	if r, ok1 := value.(*http.Response); ok1 {
		return r, http.StatusOK
	}
	return nil, http.StatusNotFound
}

func (c *contentT) put(key string, resp *http.Response) int {
	c.m.Store(key, resp)
	return http.StatusOK
}
