package proto

import (
	"fmt"
	"net/http"
)

func StartServer(wor *World) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		fmt.Println("Serving at 58000")

		var handler = func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			outp := wor.GetJSON()
			fmt.Fprintf(w, outp)
			//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		}

		http.HandleFunc("/", handler)
		http.ListenAndServe(":58000", nil)
	}()
}
