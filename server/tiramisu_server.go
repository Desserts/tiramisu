package main

import (
	"flag"
	"net/http"
	"tiramisu/config"

	log "github.com/golang/glog"
)

var (
	mux *http.ServeMux
)

func init() {
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")
}

func main() {
	startServer()
}

func startServer() {
	mux = http.NewServeMux()
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/", index)
	log.Info("start server: ", config.SERVER_PORT)
	http.ListenAndServe(":"+config.SERVER_PORT, mux)
}
