package main

import (
	"fmt"
	"time"
)

func run(i int) {
	defer func() {
		fmt.Println("i am", i)
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// go func() {
	// 	panic("inner go")
	// }()
	panic("haha")
	i += 1
}

func main() {
	for {
		go run(2)

		time.Sleep(time.Second)
	}
}
