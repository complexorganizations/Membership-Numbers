package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
)

var (
	charactersList  = 1234567890
	charactersCount = 16
)

const (
	outputFile = "output.json"
)

type MembershipReport struct {
	EXP    string `json:"exp-date"`
	CVC    string `json:"cvc"`
}

func main() {
	//
}
