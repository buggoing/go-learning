package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type st struct {
		a bool
		b int32
		c float64
		d []uint
		e *st
	}
	s := st{}
	fmt.Println(unsafe.Sizeof(s.a))
	fmt.Println(unsafe.Sizeof(s.b))
	fmt.Println(unsafe.Sizeof(s.c))
	fmt.Println(unsafe.Sizeof(s.d))
	fmt.Println(unsafe.Sizeof(s.e))
}
