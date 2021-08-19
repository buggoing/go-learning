package main

import "fmt"

func main() {
	f := hello
	defer f()
	f = hello2
	defer f()

	for i := 0; i < 100000; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func hello() {
	fmt.Println("hello")
}

func hello2() {
	fmt.Println("hello2")
}
