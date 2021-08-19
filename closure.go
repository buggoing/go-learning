package main

import (
	"fmt"
	"time"
)

type Point struct {
	X map[string]int
}

func main() {
	m := map[string]int{
		"hello": 3,
	}
	a := &Point{
		X: m,
	}

	go func() {
		for {
			fmt.Printf("in routine a: %p %v\n", a, a)
			time.Sleep(time.Second)
		}
	}()
	for {
		fmt.Printf("int main a: %p %v\n", a, a)
		time.Sleep(time.Second)
		m := map[string]int{
			"world": 3,
		}
		a.X = m
	}
	ch := make(chan int)
	<-ch
}
