package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hey there </h1>")
	fmt.Fprintf(w, "<p> Go is simple </p>")
	fmt.Fprintf(w, "<p> Go is fast to write </p>")
	fmt.Fprintf(w, "<p> you %s even add %s</p> ", "can", "<strong>variables</strong>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1> Hey there</h1>
        <p> Go is simple rather than C/C++</p>`)
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}
