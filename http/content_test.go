package http

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
	"time"
)

/*
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


*/

func putCache(url string, timeout time.Duration) (*http.Response, error) {
	// create request and process exchange
	ctx, cancel := httpx.NewContext(nil, timeout)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(iox.AcceptEncoding, iox.GzipEncoding)
	resp, err1 := httpx.ExchangeWithTimeout(timeout, nil)(req)
	if err1 != nil {
		return resp, err1
	}

	status := cache.put(url, resp)
	return resp, errors.New(fmt.Sprintf("code:%v", status))
}

func ExampleCache_No_Timeout() {
	url := "https://www.google.com/search?q=golang"
	timeout := time.Millisecond * 0
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(iox.AcceptEncoding, iox.GzipEncoding)
	resp1, status := cache.get(req.URL.String())
	fmt.Printf("test: cache.get() [status:%v] [header:%v] [code:%v]\n", resp1.StatusCode, resp.Header != nil, status)

	// verify that the response body can be read
	if status == http.StatusOK {
		buf, err1 := io.ReadAll(resp1.Body)
		fmt.Printf("test: io.ReadAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=golang] [timeout:0s]
	//test: cache.Put() [status:200] [code:200]
	//test: cache.Get() [status:200] [header:true] [code:200]
	//test: io.ReadAll() [err:<nil>] [buf:40984]

}

func ExampleCache_Timeout_504() {
	url := "https://www.google.com/search?q=erlang"
	timeout := time.Millisecond * 10
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(iox.AcceptEncoding, iox.GzipEncoding)
	_, status := cache.get(req.URL.String())
	fmt.Printf("test: cache.get() [code:%v]\n", status)

	// verify that the response body can be read
	if status == http.StatusOK {
		buf, err1 := io.ReadAll(resp.Body)
		fmt.Printf("test: io.ReadAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=erlang] [timeout:10ms]
	//test: cache.put() [status:504] [Get "https://www.google.com/search?q=erlang": context deadline exceeded]
	//test: cache.get() [code:404]

}

func ExampleCache_Timeout_200() {
	url := "https://www.google.com/search?q=pascal"
	timeout := time.Second * 5
	fmt.Printf("test: ExampleCache() [url:%v] [timeout:%v]\n", url, timeout)

	resp, err := putCache(url, timeout)
	fmt.Printf("test: cache.put() [status:%v] [%v]\n", resp.StatusCode, err)

	// Get cached response
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = make(http.Header)
	req.Header.Add(iox.AcceptEncoding, iox.GzipEncoding)
	resp1, status := cache.get(req.URL.String())
	fmt.Printf("test: cache.get() [status:%v] [header:%v] [code:%v]\n", resp.StatusCode, resp.Header != nil, status)

	// verify that the response body can be read
	if status == http.StatusOK {
		buf, err1 := io.ReadAll(resp1.Body)
		fmt.Printf("test: io.ReadAll() [err:%v] [buf:%v]\n", err1, len(buf))
	}

	//Output:
	//test: ExampleCache() [url:https://www.google.com/search?q=pascal] [timeout:5s]
	//test: cache.put() [status:200] [code:200]
	//test: cache.get() [status:200] [header:true] [code:200]
	//test: io.ReadAll() [err:<nil>] [buf:40912]

}
