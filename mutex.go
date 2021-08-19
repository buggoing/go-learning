package main

import (
	"fmt"
	"sync"
	"time"
)

type reportList struct {
	deviceID string
	ipList   []string
}

func main() {
	// name := make(map[int]int)
	// for i := 0; i < 1000; i++ {
	// 	name[i] = i
	// }
	// for i := 0; i < 100; i++ {
	// 	go func(i int) {
	// 		for {
	// 			n := rand.Intn(1000)
	// 			fmt.Printf("routine: %v, k: %v v: %v\n", i, n, name[n])
	// 			name[n] = n
	// 			// time.Sleep(time.Second)
	// 		}
	// 	}(i)
	// }
	// ch := make(chan int)
	// <-ch

	var lo sync.Mutex
	var wg sync.WaitGroup
	v := make([]reportList, 3000)
	for i, item := range v {
		item.deviceID = fmt.Sprintf("%d", i)
		item.ipList = make([]string, 10)
		for j := range item.ipList {
			item.ipList[j] = fmt.Sprint("%d", j)
		}
		v[i] = item
	}
	// start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(t int, wg *sync.WaitGroup) {
			defer wg.Done()
			start := time.Now()

			lo.Lock()
			// b := v
			lo.Unlock()

			// ipSet := make(map[string]string)
			// for _, item := range b {
			// 	if item.deviceID == "test" {
			// 		continue
			// 	}
			// 	for _, ip := range item.ipList {
			// 		addr := fmt.Sprintf("%d-%d", item.deviceID, ip)
			// 		ipSet[addr] = item.deviceID
			// 	}
			// }
			var res []string
			// for k := range ipSet {
			// 	res = append(res, k)
			// 	if len(res) == 1000 {
			// 		break
			// 	}
			// }
			fmt.Printf("elapse: %v, %d\n", time.Since(start), len(res))
		}(i, &wg)
	}
	wg.Wait()
	// fmt.Println("elapse: ", time.Since(start))
}
