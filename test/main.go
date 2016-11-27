package main

import (
	"net/http"

	"github.com/fortytw2/httpkit"
)

func main() {
	http.Handle("/", httpkit.ErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(200)
		return nil
	}))
}
