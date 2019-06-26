package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	isNeg := false
	if x < 0 {
		x = -x
		isNeg = true
	}
	xstring := strconv.Itoa(x)
	var newx []byte
	for i := len(xstring) - 1; i >= 0; i-- {
		newx = append(newx, xstring[i])

	}
	y, _ := strconv.Atoi(string(newx))
	if isNeg {
		y = -y
	}
	return y
}

func main() {
	test := []int{233, 4455, 533, 2232, -1235, 2100, -3201, -980, 0, 2, -1, -0}
	for _, value := range test {
		re := reverse(value)
		fmt.Printf("a: %d\tre: %d\n", value, re)
	}
}
