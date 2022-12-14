package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	processDataStream(string(data))
}

func processDataStream(data string) {

	firstMarker := 0
	firstSoM := 0

	for i := 0; i < len(data); i++ {

		checkMarker := ""
		checkSoM := ""

		if i+3 < len(data) {
			checkMarker = string(data[i]) + string(data[i+1]) + string(data[i+2]) + string(data[i+3])
		}

		if i+13 < len(data) {
			checkSoM = string(data[i]) + string(data[i+1]) + string(data[i+2]) + string(data[i+3]) + string(data[i+4]) +
				string(data[i+5]) + string(data[i+6]) + string(data[i+7]) + string(data[i+8]) + string(data[i+9]) +
				string(data[i+10]) + string(data[i+11]) + string(data[i+12]) + string(data[i+13])
		}

		if !hasDups(checkMarker) {

			if firstMarker == 0 {
				firstMarker = (strings.Index(data, checkMarker)) + 4
			}
		}

		if !hasDups(checkSoM) {

			if firstSoM == 0 {
				firstSoM = (strings.Index(data, checkMarker)) + 14
			}
		}

	}

	fmt.Printf("First marker found at: %d\n", firstMarker)
	fmt.Printf("First start of message found at: %d\n", firstSoM)
}

func hasDups(checkStr string) bool {

	dup := false

	for _, char := range checkStr {
		if strings.Count(checkStr, string(char)) > 1 {
			dup = true
		}
	}

	return dup
}
