package main

import (
	"net/http"
)

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("That's is awesome"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})

	http.ListenAndServe(":35464", mux)
}
