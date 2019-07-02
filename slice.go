package main

import "fmt"

func main() {
	sli1 := []int{23, 34, 56}

	sli2 := sli1

	sli3 := sli1
	sli4 := sli1[:]
	sli5 := make([]int, len(sli1))
	copy(sli5, sli1)

	sli2[0] = -100

	fmt.Println(sli1)
	fmt.Println(sli2)
	fmt.Println(sli3)
	fmt.Println(sli4)
	fmt.Println(sli5)
	fmt.Println("cap: ", cap(sli1), cap(sli5))

	li := sli1[1:]
	// li = append(li, 33, 22, 11, 3, 444444, 5)
	fmt.Println(li)
	li[1] = 9999
	fmt.Println(sli1)
	fmt.Println(li)

}
