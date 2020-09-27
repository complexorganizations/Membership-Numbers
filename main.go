package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	charactersList  = 0123456789
	charactersCount = 16
	outputFile      = "output.json"
)

type NumbersReport struct {
	DATE string `json:"date"`
	CVC  int8 `json:"cvc"`
}

func writingInFile(b []byte) {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Println(err)
	}
	if _, err = file.Write(b); err != nil {
		fmt.Printf("Error writing to a file %s", err)
	}
	file.Close()
}

func main() {
	for {
		fmt.Println(outputFile)
	}
}
