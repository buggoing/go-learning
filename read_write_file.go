package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func writeJson(pers *Person, filename string) (err error) {
	// jsonData, err := json.Marshal(pers)
	jsonData, err := json.MarshalIndent(pers, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	jsonFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	return nil
}

func readJson(pers *Person, filename string) (err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_ = json.Unmarshal([]byte(file), pers)
	return nil
}

func main() {
	filename := "alice.json"
	pers := Person{Name: "Alice", Age: 23}
	err := writeJson(&pers, filename)
	if err != nil {
		panic(err)
	}
	pers2 := Person{}
	err = readJson(&pers2, filename)

	if err != nil {
		panic(err)
	}
	fmt.Println(pers2)
}
