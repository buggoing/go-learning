package main

import (
	"flag"
	"fmt"
)

func fib(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
		fmt.Println(i, b)
	}
	return b
}

func main() {
	num := flag.Int("num", 0, "max num")
	flag.Parse()
	res := fib(*num)
	fmt.Println(res)
}
