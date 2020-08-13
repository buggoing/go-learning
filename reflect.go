package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string `json:"name"`
	Number int64  `json:"number"`
}

func main() {
	s := &Student{
		Name:   "name",
		Number: 123,
	}
	values := reflect.ValueOf(s).Elem()
	types := reflect.TypeOf(s).Elem()
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i)
		if v.Type().Kind() == reflect.String {
			t := types.Field(i)
			jsonTag := t.Tag.Get("json")
			fmt.Printf("json tag: %v is type string\n", jsonTag)
		}
	}
}
