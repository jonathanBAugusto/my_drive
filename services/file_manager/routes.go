package main

import "net/http"

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	mux.HandleFunc("GET /space", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}
