package server

import "net/http"

func handleFuncMainPage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(218)
	_, _ = res.Write([]byte("Hello from main page!" + req.Method))
}

func handleFuncPingPage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	_, _ = res.Write([]byte("Pong"))
}
