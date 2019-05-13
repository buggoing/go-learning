package main

import (
	"fmt"
	"sync"
)

func main() {
	score := map[string]int{
		"foo": 23,
		"bar": 2222,
	}
	newScore := make(map[string]*int)
	for key, sc := range score {
		sco := sc
		fmt.Println(key, &sc)
		newScore[key] = &sco
	}

	for key, sc := range newScore {
		fmt.Println(key, *sc)
	}

	var wait sync.WaitGroup
	// runtime.GOMAXPROCS(1)
	for i := 0; i < 100000; i++ {
		wait.Add(1)
		go func() {
			fmt.Println(i)
			wait.Done()
		}()

	}
	wait.Wait()
}
