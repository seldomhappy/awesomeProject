package server

import (
	"log"
	"net/http"
)

func Start() {
	//mux := http.ServeMux{}
	http.HandleFunc("/", handleFuncMainPage)
	http.HandleFunc("/ping", handleFuncPingPage)

	log.Println("Listen HTTP server on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
