package main

import "fmt"

type Node struct {
	Val []string
}

func updateSlice(s []string) {
	s[0] = "changed"
}

func main() {
	n := Node{
		Val: []string{"hello", "world"},
	}
	fmt.Println("before: ", n)
	n2 := n
	n2.Val[0] = "world"
	fmt.Println("after: ", n)

	l1 := []string{"hello", "world"}
	fmt.Println("before: ", l1)
	l2 := l1
	l2[0] = "world"
	fmt.Println("after: ", l1)

	m1 := make(map[string]string)
	m1["hello"] = "hello"
	fmt.Println("before: ", m1)
	m2 := m1
	m2["hello"] = "world"
	fmt.Println("after: ", m1)

	s1 := []string{"hello", "world"}
	fmt.Println("before: ", s1)
	updateSlice(s1)
	fmt.Println("after: ", s1)
}
