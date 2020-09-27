package main

import (
	"fmt"
	"io/ioutil"
)

var (
	charactersList  = 0123456789
	charactersCount = 16
	outputFile      = "output.json"
)

type NumbersReport struct {
	DATE string `json:"date"`
	CVC  string `json:"cvc"`
}

func main() {
	fmt.Println(test)
}
