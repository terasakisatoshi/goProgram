/*
webserver.go shows a simple web server.
usage:
open terminal:
$ go run webserver.go

Then,acccess http://localhost:8080/
we must get some HTML to prompt starting funny chat.
*/

package main

import (
    "log"
    "net/http"
    "path/filepath"
    "sync"
    "text/template"
)

// templ represents a single template
type templateHandler struct {
    once     sync.Once
    filename string
    templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })

    t.templ.Execute(w, nil)
}

func main() {
    http.Handle("/", &templateHandler{filename: "chat.html"})
    //start web server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
