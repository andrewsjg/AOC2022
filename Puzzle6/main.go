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

// 1953
func processDataStream(data string) {

	firstMarker := 0

	for i := 0; i < len(data); i++ {

		checkStr := ""
		if i+3 < len(data) {
			checkStr = string(data[i]) + string(data[i+1]) + string(data[i+2]) + string(data[i+3])
		}

		if !hasDups(checkStr) {
			//fmt.Printf("Marker: %s at %d\n", checkStr, (strings.Index(data, checkStr))+4)

			if firstMarker == 0 {
				firstMarker = (strings.Index(data, checkStr)) + 4
			}
		}

	}

	fmt.Printf("First marker found at: %d\n", firstMarker)
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
