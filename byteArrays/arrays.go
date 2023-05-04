package main

import (
	"fmt"
	"os"
	"bytes"
)

func main(){
	byteArray := []byte{0x41, 0x53,0x43,0x49,0x49}

	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write(byteArray)
	var buf bytes.Buffer
	buf.Write(byteArray)
	buf.Write([]byte(" ")) // add space
	buf.Write(byteArray)
	// Display message to stdout
	fmt.Println("buf = " + buf.String())

	s := string(byteArray)
	fmt.Println("s = " + s)
}

