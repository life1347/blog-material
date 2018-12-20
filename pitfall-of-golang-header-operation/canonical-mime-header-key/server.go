package main

import (
	"fmt"
	"log"
	"net/http"
	"net/textproto"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		v1 := r.Header["foo"]
		v2 := r.Header[textproto.CanonicalMIMEHeaderKey("foo")]

		msg := fmt.Sprintf("V1: %v, V2: %v", v1, v2)

		w.Write([]byte(msg))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
