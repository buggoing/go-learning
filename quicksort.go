package main

import "fmt"

type item int

func quickSort(arr []item) []item {
	if len(arr) <= 1 {
		return arr
	}

	pvit := arr[0]
	var left, right []item
	for i := 1; i < len(arr); i++ {
		if arr[i] < pvit {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	left = append(left, pvit)
	left = append(left, right...)
	return left
}

func main() {
	arr := []item{1, 4, 342456, 234, 22, 222, 4534, 455}
	// arr := []item{23, 23, 1}
	result := quickSort(arr)
	fmt.Println(arr)
	fmt.Println(result)
}
