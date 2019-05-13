package main

import "fmt"

// unbuffer chan 的send和receive都会阻塞
func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("ch: ", <-ch)

	unbufCh := make(chan int)
	unbufCh <- 1
	fmt.Println("unbufCh: ", <-unbufCh)

}
