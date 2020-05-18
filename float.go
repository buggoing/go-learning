package main

import (
	"fmt"
	"math"
)

func main() {

	maxFloat32 := float32(1 << 24)
	var a float32 = 2
	var b float32 = 5
	fmt.Printf("a: %b\n", math.Float32bits(a))
	fmt.Printf("b: %b\n", math.Float32bits(b))
	fmt.Println("b-a: ", b-a)

	a = maxFloat32
	b = maxFloat32 + 1
	fmt.Printf("b: %b, a: %f, b-a: %d", math.Float32bits(b), a, b-a)

}
