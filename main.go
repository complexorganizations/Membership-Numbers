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
	NUMBER string `json:"md5"`
	CVC    string `json:"sha-1"`
	EXP    string `json:"sha-256"`
}

func RandomString() {
	// generate all the random strings here.
}

func file() {
	// write to file and appened to file goes on here.
}

func main() {
	//
}
