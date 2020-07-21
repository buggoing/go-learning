package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	arrLen = 1000
)

func useAppend(arr []string) []string {
	var res []string
	for _, value := range arr {
		res = append(res, value)
	}
	return res
}

func useMake(arr []string) []string {
	res := make([]string, len(arr))
	for i, value := range arr {
		res[i] = value
	}
	return res
}

func main() {
	arr := make([]string, arrLen)
	for i, _ := range arr {
		arr[i] = "helloalksdaflskdfjalsdk" + strconv.FormatInt(int64(i), 10)
	}
	start := time.Now()
	for i := 0; i < 100; i++ {
		useAppend(arr)
	}
	fmt.Println("elpase: ", time.Since(start))

	start = time.Now()
	for i := 0; i < 100; i++ {
		useMake(arr)
	}
	fmt.Println("elpase: ", time.Since(start))
}
