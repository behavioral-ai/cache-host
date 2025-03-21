package http2

import (
	"bytes"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
)

func ExampleNewCache() {
	uri := "https://www.google.com/search?q=golang"
	s := "this is string content"
	c := newCache()

	req, _ := http.NewRequest(http.MethodPut, uri, io.NopCloser(bytes.NewReader([]byte(s))))
	req.Header.Set("key-1", "value-1")
	req.Header.Set("key-2", "value-2")
	req.Header.Set("key-3", "value-3")
	c.put(req)

	req2, _ := http.NewRequest(http.MethodGet, uri, nil)
	resp := c.get(req2)
	buf, err := iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewCache() -> [%v] [%v] [%v] [err:%v]\n", resp.StatusCode, resp.Header, string(buf), err)

	req3, _ := http.NewRequest(http.MethodGet, "https://bing.com", nil)
	resp = c.get(req3)
	buf, err = iox.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewCache() -> [%v] [%v] [%v] [err:%v]\n", resp.StatusCode, resp.Header, buf, err)

	//Output:
	//test: NewCache() -> [200] [map[Key-1:[value-1] Key-2:[value-2] Key-3:[value-3]]] [this is string content] [err:<nil>]
	//test: NewCache() -> [404] [map[]] [[]] [err:<nil>]

}
