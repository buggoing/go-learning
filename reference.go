package main

import "fmt"

type Node struct {
	Val []string
}

func main() {
	n := Node{
		Val: []string{"hello", "world"},
	}
	n2 := n
	n2.Val[0] = "world"
	fmt.Println("n: ", n)
}
