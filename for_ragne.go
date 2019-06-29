package main

import (
	"fmt"
	"sync"
)

type Stu struct {
	Name  string
	Score int
}

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

	stus := map[string]Stu{
		"foo": Stu{"foo", 22},
		"bar": Stu{"bar", 2344},
	}

	var stusList []Stu
	for k, val := range stus {
		fmt.Println(k, val)
		stusList = append(stusList, val)
	}
	fmt.Println(stusList)

	nameS := make(map[string]Stu)
	for i, val := range stusList {
		fmt.Println(i, val)
		nameS[val.Name] = val

	}
	fmt.Println("nameS: ", nameS)

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
