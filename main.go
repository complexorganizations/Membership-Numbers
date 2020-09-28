package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

var (
	charactersList      = 1234567890
	cardCharactersCount = 16
	expCharactersCount  = 4
	csvCharactersCount  = 4
	outputFile          = "output.json"
)

type MembershipReport struct {
	// NUMBER string `json:"card-number"`
	EXP    string `json:"exp-date"`
	CVC    string `json:"cvc"`
}

func RandomString() int {
	// use a random seed
	rand.Seed(time.Now().UnixNano())
	// generate all the random strings here.
	randInt := rand.Intn(10)
	// return value
	return randInt
}

func file() {
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
		// loop everything here.
	    }
}
