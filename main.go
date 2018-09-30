package main

import (
	"net/http"
    "html/template"
    "time"
    "log"
)

const (
    HOST = ":8080"
    WEB_ROOT = "./web/"
)

var timestamp = time.Now().Unix()

type Action struct {
    Pin string
}
type WebData struct {
    Timestamp   int64
    PageTitle   string
    Actions     []Action
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    data := WebData {
        Timestamp: timestamp,
        PageTitle: "go pi controlla",
        Actions: []Action {
            {Pin: "eighteen"},
            {Pin: "twentythree"},
        },
    }
    index := template.Must(template.ParseFiles(WEB_ROOT + "index.html"))
    index.Execute(w, data)
}

func main() {
	hub := newHub()
	go hub.run()

    assets := http.FileServer(http.Dir(WEB_ROOT + "assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", assets))

    http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

    log.Print("starting server on localhost" + HOST)
    log.Fatal(http.ListenAndServe(HOST, nil))
}
