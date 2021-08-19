package main

import "fmt"

type St struct {
	Number       float64
	Depreciation float64
}

func update(arr []St) {
	for i, _ := range arr {
		arr[i].Number = 0
	}
}

func main() {
	rows := []St{
		{
			Number:       3,
			Depreciation: 0.3,
		},
		{
			Number:       4,
			Depreciation: 0.1,
		},
	}
	for _, row := range rows {
		row.Number *= row.Depreciation
	}
	fmt.Printf("%+v\n", rows)

	update(rows)
	fmt.Printf("%+v\n", rows)

}
