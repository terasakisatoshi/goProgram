package main

import "fmt"

func addFloat64(x float64, y float64) float64 {
	return x + y
}

func addFloat32(x float32, y float32) float32 {
	return x + y
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(addFloat64(num1, num2))

	var num3, num4 float64 = 2.1, 3.3

	fmt.Println(addFloat64(num3, num4))

	x, y := float32(4.5), float32(6.3)

	fmt.Println(addFloat32(x, y))

}
