package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	proxy *httputil.ReverseProxy
}

func NewProxy(targetUrl string) Proxy {
	u, _ := url.Parse(targetUrl)
	return Proxy{
		proxy: httputil.NewSingleHostReverseProxy(u),
	}
}

func (p Proxy) handle(w http.ResponseWriter, r *http.Request) {
	// We may want to attach metadata to request header
	r.Header.Add("Header1", "bar")
	r.Header.Set("Header2", "foo")
	p.proxy.ServeHTTP(w, r)
}

func main() {
	proxy := NewProxy("http://127.0.0.1:8888/")

	http.HandleFunc("/", proxy.handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
