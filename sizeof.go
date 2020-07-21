package main

import (
	"fmt"
	"unsafe"
)

type T struct{}

func main() {
	t := T{}
	fmt.Printf("sizeof(%T) is: %d\n", t, unsafe.Sizeof(t))
	i := int64(0)
	fmt.Printf("sizeof(%T) is: %d\n", i, unsafe.Sizeof(i))
	m := make(map[string]interface{})
	m["1"] = nil
	m["2"] = nil
	fmt.Printf("sizeof(%T) is: %d\n", m, unsafe.Sizeof(m))
	b := true
	fmt.Printf("sizeof(%T) is: %d\n", b, unsafe.Sizeof(b))
	var s interface{}
	fmt.Printf("sizeof(%T) is: %d\n", s, unsafe.Sizeof(s))
}
