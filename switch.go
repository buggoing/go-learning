package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	i := 0
	if len(os.Args) > 1 {
		i, _ = strconv.Atoi(os.Args[1])
	}

	switch i {
	case -1:
		fmt.Println("less than 0")
		fmt.Println("less than 0")
	case 0:
		fmt.Println("less than 1")
		fmt.Println("less than 1")
	default:
		fmt.Println("default")
		fmt.Println("default")
	}

	var j interface{} = 2.3
	switch t := j.(type) {
	case float32:
		fmt.Println("float32: ", t)
	case float64:
		fmt.Println("float64: ", t)
	default:
		fmt.Println("not float: %v", t)
	}

}
