package main

import "fmt"

func cntGen() func() uint64 {
	var counter uint64
	inner := func() uint64 {
		counter += 1
		return counter
	}
	return inner
}

func main() {
	count := cntGen()
	fmt.Println(count()) //1
	fmt.Println(count()) //2
	fmt.Println(count()) //3
	fmt.Println(count()) //4
}
