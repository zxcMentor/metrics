package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		targetUrl, err := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))

		if !strings.HasPrefix(r.URL.Path, "/api") && !strings.HasPrefix(r.URL.Path, "/swagger") {

			if err != nil {
				fmt.Println("url parsing error:", err)
			}

			proxy := httputil.NewSingleHostReverseProxy(targetUrl)

			proxy.ServeHTTP(w, r)

			return

		}

		handler.ServeHTTP(w, r)

	})
}
