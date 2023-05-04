package main

import "fmt"

func msg2ch(ch chan<- string, msg string) {
	ch <- msg
}

func ch2ch(ch_one <-chan string, ch_another chan<- string) {
	msg := <-ch_one
	ch_another <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	msg2ch(pings, "passed message 1")
	ch2ch(pings, pongs)
	fmt.Println(<-pongs)

	msg2ch(pings, "passed message 2")
	ch2ch(pings, pongs)
	fmt.Println(<-pongs)
}
