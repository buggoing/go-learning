package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type City struct {
	Name       string `json:"name"`
	Admaster   string `json:"admaster"`
	NameEn     string `json:"name_en"`
	Province   string `json:"province"`
	ProvinceEn string `json:"province_en"`
	Tier       string `json:"tier"`
}

func main() {
	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = '	'
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	citis := make(map[string]City)
	for _, line := range lines {
		city := City{
			Name:       line[0],
			Admaster:   line[1],
			NameEn:     line[2],
			Province:   line[3],
			ProvinceEn: line[4],
			Tier:       line[6],
		}
		citis[city.NameEn] = city
	}
	jsonData, err := json.MarshalIndent(citis, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Create("city_en.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)

	fmt.Printf("city: %v", citis)
}
