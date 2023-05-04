package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	msg := "hello world\n"
	fmt.Println("Write" + " " + msg)
	file.Write([]byte(msg))
	file.Close()
}
