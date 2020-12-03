package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	charactersList  = 1234567890 // characters
	charactersCount = 16         // Card number length
	cvcCharacters   = 3          // CVC kebgth
	cvcCount        = 3          // Num of CVC's needed for each card
	expiryDuration  = 5          // In years
	numOfEntries    = 1          // Num of card entries needed
)

const (
	outputFile = "output.json"
)

type membershipReport struct {
	EXP []string `json:"exp-date"`
	CVC []string `json:"cvc"`
}

// Fetches the random int based on given character length

func randomInt(min, max, source int64) int64 {
	sourceObj := rand.NewSource(source * time.Now().UnixNano())
	randObj := rand.New(sourceObj)
	return min + randObj.Int63n(max-min)
}

// Fetches the random CVC numbers
func getCVC() []string {
	s := make([]string, 0)
	min, max := getMinMax(cvcCount)
	for i := 1; i <= cvcCount; i++ {
		s = append(s, strconv.FormatInt(randomInt(min, max, int64(i)), 10))
	}
	fmt.Println(s)
	return s

}

// Fetches expiry date with all the months
func getExpDate() []string {
	dates := make([]string, 0)
	year := time.Now().Year() + expiryDuration
	for i := 1; i <= 12; i++ {
		dates = append(dates, fmt.Sprintf("%d/%d", i, year))
	}
	return dates
}

// Gets the min and max number for the given length
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

// Checks whether file exist or not
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Write to the file
func writingInFile(b []byte) {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("OS Error: %s", err)
	}
	if _, err = file.Write(b); err != nil {
		fmt.Printf("Error writing to a file %s", err)
	}
	file.Close()
}

// Calculates the random CVC and Year values
func writeToJSON(num int64) {
	// writes data to JSON outputFile
	result := []map[string]*membershipReport{}

	data := make(map[string]*membershipReport)

	key := strconv.FormatInt(num, 10)

	cvc := getCVC()

	expDate := getExpDate()

	data[key] = &membershipReport{expDate, cvc}

	fmt.Println(data)

	if !fileExists(outputFile) {
		//simple append
		result = append(result, data)
		b, err := json.Marshal(&result)
		if err != nil {
			log.Fatalln("Error while marshaling: ", err)
		}
		writingInFile(b)
		return
	}
	//check for duplications
	byteValue, err := ioutil.ReadFile(outputFile)
	if err != nil {
		log.Fatalln("Error while reading: ", err)
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatalln("Error while marshaling: ", err)
	}

	isFound := false
	for _, r := range result {
		if reflect.DeepEqual(r, data) {
			isFound = true
			break
		}
	}

	if isFound {
		fmt.Println("Duplicate Found.")
		return
	}

	//appending
	result = append(result, data)
	b, err := json.Marshal(&result)
	if err != nil {
		log.Fatalln("Error while marshaling: ", err)
	}
	writingInFile(b)
}

func main() {
	for i := 1; i <= numOfEntries; i++ {
		min, max := getMinMax(charactersCount)
		writeToJSON(randomInt(min, max, int64(i)))
	}
}
