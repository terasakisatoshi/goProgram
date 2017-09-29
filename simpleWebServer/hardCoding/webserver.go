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
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
            <html>
                <head>
                    <title> let's fun chat!</title>
                </head>
                <body> let's get started our funny chat!</body>
            </html>
        `))
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
