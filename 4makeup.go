package main

import (
	"github.com/shaoshing/train"
	"github.com/t0pep0/4makeup/controllers"
	"log"
	"net/http"
	"runtime"
	"time"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Proto, r.Method, r.RemoteAddr, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", controllers.RenderController)
	train.ConfigureHttpHandler(http.DefaultServeMux)
	server := &http.Server{
		Addr:           "127.0.0.1:5000",
		Handler:        httpLog(http.DefaultServeMux),
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Start server on 127.0.0.1:5000")
	server.ListenAndServe()
}
