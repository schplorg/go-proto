package server

import (
	"fmt"
	"net/http"
)

func Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		fmt.Println("Serving at 58000")

		var handler = func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		}

		http.HandleFunc("/", handler)
		http.ListenAndServe(":58000", nil)
	}()
}
