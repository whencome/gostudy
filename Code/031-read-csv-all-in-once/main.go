package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// open file
	path := "../../Resources/iris_data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("open file fail: %v", err)
	}
	defer file.Close()

	// create csv reader
	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		log.Panicf("read file fail: %v", err)
	}

	// print data
	if len(csvData) == 0 {
		log.Println("empty data")
		return
	}
	for _, line := range csvData {
		fmt.Println(line)
	}
}
