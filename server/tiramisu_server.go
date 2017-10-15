package main

import (
	"net/http"
	"tiramisu/config"
)

var (
	mux *http.ServeMux
)

func main() {
	startServer()
}

func startServer() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	http.ListenAndServe(":"+config.SERVER_PORT, mux)
}

func login(w http.ResponseWriter, r *http.Request) {

}
