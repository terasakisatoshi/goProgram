package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	targetURL := "https://pythonprogramming.net/go/parsing-go-language-programming-tutorial/"
	response, _ := http.Get(targetURL)
	bytes, _ := ioutil.ReadAll(response.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	response.Body.Close()
}
