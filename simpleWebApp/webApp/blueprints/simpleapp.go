package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
        <html>
            <head>
                <title> Chat </title>
            </head>
            <body>
                Let's have fun
            </body>
        </html>
        `))
}
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
