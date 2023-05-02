package main

import (
	"fmt"
	"math"
)

const msg string = "constant"
const az = "Azarashi"

func main() {
	fmt.Println(msg)
	fmt.Println(az)

	const n = 100
	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))
	fmt.Println(math.Pi)
	s := math.Sin(n)
	c := math.Cos(n)
	fmt.Println(s * s + c * c)
}