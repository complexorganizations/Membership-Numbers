package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	charactersList  = 1234567890
	charactersCount = 16
	characterLimit  = 2
)

const (
	outputFile = "output.json"
)

type fields struct {
	//empty
}

func getMinMax(length int) (low int64, high int64) {
	//returns min and max values of given character length
	min := "1"
	max := "9"

	for i := 1; i < length; i++ {
		min = min + "0"
		max = max + "9"
	}

	low, _ = strconv.ParseInt(min, 10, 64)
	high, _ = strconv.ParseInt(max, 10, 64)

	return
}

func writeToJSON(low int64, high int64) {
	// writes data to JSON outputFile

	f, err := os.Create(outputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := make(map[string]fields)

	for i := low; i <= high; i++ {
		key := strconv.FormatInt(i, 10)
		data[key] = fields{}
	}

	fmt.Println(data)

	jsonString, _ := json.Marshal(data)

	_, err2 := f.Write(jsonString)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func main() {
	low, high := getMinMax(characterLimit)
	writeToJSON(low, high)
}
