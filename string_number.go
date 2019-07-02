package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "23233"
	num, _ := strconv.Atoi(s)

	newS := strconv.Itoa(222111)
	oct, _ := strconv.ParseInt("23", 8, 0)

	fmt.Println(num, newS)
	fmt.Println(oct)
}
