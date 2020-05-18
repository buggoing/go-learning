package main

import "fmt"

func main() {
	f := hello
	defer f()
	f = hello2
	defer f()
}

func hello() {
	fmt.Println("hello")
}

func hello2() {
	fmt.Println("hello2")
}
