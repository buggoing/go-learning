package main

import (
	"fmt"
	"time"
)

func run(i int) {
	// defer func() {
	// 	fmt.Println("i am", i)
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered in f", r)
	// 	}
	// }()

	// go func() {
	// 	panic("inner go")
	// }()
	panic("haha")
	i += 1
}

func main() {

	go func() {
		defer func() {
			fmt.Println("i am")
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		panic("meme")
		run(2)
	}()
	for {
		time.Sleep(10 * time.Hour)
	}
}
