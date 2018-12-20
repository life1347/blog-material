package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		h1 := r.Header["Header1"]
		h2 := r.Header["Header2"]
		msg := fmt.Sprintf("header1: %v, header2: %v", h1, h2)
		w.Write([]byte(msg))
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
