package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	name := make(map[int]int)
	for i := 0; i < 1000; i++ {
		name[i] = i
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				n := rand.Intn(1000)
				fmt.Printf("routine: %v, k: %v v: %v\n", i, n, name[n])
				name[n] = n
				time.Sleep(time.Second)
			}
		}(i)
	}
	ch := make(chan int)
	<-ch
}
