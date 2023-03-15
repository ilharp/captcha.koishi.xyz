package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	var listen string
	if len(os.Args) > 1 {
		listen = os.Args[1]
	} else {
		listen = "127.0.0.1:8081"
	}

	proxy := &httputil.ReverseProxy{Director: func(r *http.Request) {
		r.Host = "captcha.go-cqhttp.org"
		r.URL.Host = "captcha.go-cqhttp.org"
		r.URL.Scheme = "https"
	}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")

		if strings.HasPrefix(r.URL.Path, "/captcha") || strings.HasPrefix(r.URL.Path, "/sdk") {
			log.Printf("SERVE %s", ip)
			proxy.ServeHTTP(w, r)

			return
		}

		log.Printf("REDIR %s", ip)
		w.Header().Add("Location", "https://koishi.chat")
		w.WriteHeader(http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(listen, nil))
}
