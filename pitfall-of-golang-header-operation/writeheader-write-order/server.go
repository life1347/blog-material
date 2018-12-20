package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foobar"))
		w.WriteHeader(http.StatusBadRequest)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
