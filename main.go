package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultPort     = "8080"
	defaultBody     = "default backend - 404\n"
	healthzEndpoint = "/healthz"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, defaultBody)
	})

	http.HandleFunc(healthzEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
