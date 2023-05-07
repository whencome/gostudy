package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	// open file
	path := "../../Resources/iris_data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("open file fail: %v", err)
	}
	defer file.Close()

	// create a dataframe
	irisDF := dataframe.ReadCSV(file)
	fmt.Println(irisDF)

	// create a filter
	filter := dataframe.F{
		Colname:    "Species",
		Comparator: "==",
		Comparando: "versicolor",
	}
	// 检查过滤器是否正确
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal((versicolorDF.Err))
	}

	// 数据处理
	filteredDF := irisDF.Filter(filter).Select([]string{"Sepal.Width", "Species"})
	fmt.Println(filteredDF)

}
