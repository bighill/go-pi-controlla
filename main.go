package main

import (
	"net/http"
    "log"
)

const (
    HOST = ":8080"
    WEB_ROOT = "index.html"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, WEB_ROOT)
}

func main() {
	hub := newHub()
	go hub.run()

    http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

    log.Fatal(http.ListenAndServe(HOST, nil))
}
