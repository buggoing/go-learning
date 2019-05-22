package main

import (
	"fmt"
	"time"
)

func producer(container chan<- interface{}, interval time.Duration, maxNum int, done chan int) {
	for num := 0; num < maxNum; num++ {
		container <- num
		fmt.Printf("producer num: %d\n", num)
		time.Sleep(interval)
	}
	close(container)
	done <- 0
}

func customer(container <-chan interface{}, customerID int, interval time.Duration) {
	for {
		num, ok := <-container
		if !ok {
			return
		}
		fmt.Printf("customer: %d num: %v\n", customerID, num)
		time.Sleep(interval)
	}
}

func main() {
	container := make(chan interface{}, 100)
	done := make(chan int)
	go producer(container, 10*time.Millisecond, 1000, done)

	for i := 0; i < 5; i++ {
		// go customer(container, i, 1*time.Second)
		go customer(container, i, 20*time.Millisecond)
	}
	<-done
}
