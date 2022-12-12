package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	rucksackData, err := os.Open("input.txt")
	defer rucksackData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return
	}

	rucksackReader := bufio.NewReader(rucksackData)
	sum := 0

	for {
		rucksack, _, err := rucksackReader.ReadLine()

		if err == io.EOF {
			break
		}

		compartment1, compartment2 := getCompartments(string(rucksack))
		match := findMatch(compartment1, compartment2)

		value := int(match)

		if value < 97 {
			value = value - 38
		} else {
			value = value - 96
		}

		fmt.Printf("Rucksack: %s, Compartment1: %s, Compartment2: %s Shared Item: %s Ascii: %d, Value: %d\n", rucksack, compartment1, compartment2, string(match), int(match), value)
		sum = sum + value
	}

	fmt.Printf("The sum of all shared item priorities is: %d\n", sum)
}

func getCompartments(rucksack string) (compartment1 string, compartment2 string) {

	compartment1 = rucksack[:len(rucksack)/2]
	compartment2 = rucksack[len(rucksack)/2:]

	return compartment1, compartment2
}

func findMatch(compartment1 string, compartment2 string) (match rune) {

	for _, char := range compartment1 {
		if strings.ContainsRune(compartment2, char) {
			match = char
		}
	}

	return match
}
