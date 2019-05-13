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
	newScore1 := make(map[string]*int)
	for key, sc := range score {
		sco := sc
		fmt.Println(key, &sc)
		newScore[key] = &sco
		newScore1[key] = &sc
	}

	for key, sc := range newScore {
		fmt.Println(key, *sc)
	}
	for key, sc := range newScore1 {
		fmt.Println(key, *sc)
	}

	var wait sync.WaitGroup

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			fmt.Println(i)
			wait.Done()
		}()

	}

	for i := 200; i < 300; i++ {
		wait.Add(1)
		go func(index int) {
			fmt.Println(index)
			wait.Done()
		}(i)

	}
	wait.Wait()
}
