package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hi"))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
