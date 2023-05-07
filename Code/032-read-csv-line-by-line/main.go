package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	csvData := make([][]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panicf("read err: %v", err)
		}
		csvData = append(csvData, line)
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
