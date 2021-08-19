package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now().Format("2006-01-02 15:04:05 MST")
	fmt.Println("time: ", timeNow)
}
