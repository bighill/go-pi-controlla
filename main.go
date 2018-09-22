package main

import (
	"net/http"
    "html/template"
    "log"
)

const (
    HOST = ":8080"
    WEB_ROOT = "./web/"
)

type Action struct {
    Do string
}

type WebData struct {
    PageTitle   string
    Actions     []Action
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    data := WebData {
        PageTitle: "go pi controlla",
        Actions: []Action {
            {Do: "on"},
            {Do: "off"},
            {Do: "blink"},
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
