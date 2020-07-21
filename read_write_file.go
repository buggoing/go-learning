package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	err = json.Unmarshal(file, pers)
	return err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "alice.json"
	pers := Person{Name: "Alice", Age: 23}
	err := writeJson(&pers, filename)
	checkError(err)
	pers2 := Person{}
	err = readJson(&pers2, filename)

	checkError(err)
	fmt.Println(pers2)

	tmpFilename := "./read_write_file.go"
	file, err := os.Open(tmpFilename)
	checkError(err)
	defer file.Close()

	content := make([]byte, 1024)
	for {

		n1, err := file.Read(content)
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}
		fmt.Println("n1: ", n1)
		fmt.Println(content[:n1])
	}

}
