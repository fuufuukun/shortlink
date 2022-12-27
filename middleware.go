package main

import (
	"log"
	"net/http"
	"time"
)

type Middlerware struct {
}

// LoggingHandler log the time-cosuming of http request
func (m Middlerware) LoggingHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

// RecoverHandler recover panic
func (m Middlerware) RecoverHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
