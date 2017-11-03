package main

import "fmt"

func randXorshit32(seed uint32) func() uint32 {
	var y uint32 = seed
	inner := func() uint32 {
		y = y ^ (y << 13)
		y = y ^ (y >> 17)
		y = y ^ (y << 5)
		return y
	}
	return inner
}

func uniform(generator func() uint32) func() float32 {
	inner := func() float32 {
		return float32(generator()) / float32(^uint32(0))
	}
	return inner
}

func calcPi() {
	xor32 := randXorshit32(2463534242)
	u01 := uniform(xor32)
	N := 100000000
	counter := 0
	for i := 0; i < N; i++ {
		x := u01()
		y := u01()
		if x*x+y*y < 1.0 {
			counter += 1
		}
	}
	fmt.Println(4 * float64(counter) / float64(N))
}

func test1() {
	xor32 := randXorshit32(2463534242)
	for i := 0; i < 10; i++ {
		fmt.Println(xor32())
	}

}

func main() {
	calcPi()
}
