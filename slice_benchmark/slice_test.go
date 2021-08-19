package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var arr []string

func TestMain(m *testing.M) {
	arr := make([]string, arrLen)
	for i, _ := range arr {
		arr[i] = "helloalksdaflskdfjalsdk" + strconv.FormatInt(int64(i), 10)
	}
	fmt.Println("arro: ", arr[1])
	os.Exit(m.Run())
}

func BenchmarkUseAppend(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		useAppend(arr)
	}
}

func BenchmarkUseMake(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		useMake(arr)
	}
}
