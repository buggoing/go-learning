package main

import "fmt"

func main() {
	m1 := map[string]string{
		"hello": "world",
	}
	m2 := m1

	m2["well"] = "done"

	fmt.Printf("m1: %+v\n", m1)
	fmt.Printf("m2: %+v\n", m2)
}
