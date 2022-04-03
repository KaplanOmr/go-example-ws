package main

import (
	"flag"
	"net/http"
)

var addr = flag.String("addr", "localhost:8181", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(*addr, nil)
}
