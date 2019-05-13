package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	intCh := make(chan int, 1)
	stringCh := make(chan string, 1)
	intCh <- 1
	stringCh <- "hello"
	select {
	case value := <-intCh:
		fmt.Println(value)
	case value := <-stringCh:
		panic(value)
	}

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
