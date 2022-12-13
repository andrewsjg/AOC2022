package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type elfGroup struct {
	rucksacks []string
}

func main() {
	// Part 1
	fmt.Printf("The sum of all shared item priorities is: %d\n", sumPriorties())

	// Part 2
	groups := getElfGroups()
	sum := 0

	for _, group := range groups {
		sum = sum + priorty(group.Badge())
	}

	fmt.Printf("The sum of all the badge priorities : %d\n", sum)
}

func sumPriorties() int {
	rucksackData, err := os.Open("input.txt")
	defer rucksackData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return 0
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

		value := priorty(match)

		//fmt.Printf("Rucksack: %s, Compartment1: %s, Compartment2: %s Shared Item: %s Ascii: %d, Value: %d\n", rucksack, compartment1, compartment2, string(match), int(match), value)
		sum = sum + value
	}

	return sum
}

func getElfGroups() []elfGroup {

	elfGroups := []elfGroup{}

	rucksackData, err := os.Open("input.txt")
	defer rucksackData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return elfGroups
	}

	rucksackReader := bufio.NewReader(rucksackData)

	rsCount := 0
	CurentElfGroup := elfGroup{}

	for {
		rucksack, _, err := rucksackReader.ReadLine()

		if err == io.EOF {
			break
		}

		CurentElfGroup.rucksacks = append(CurentElfGroup.rucksacks, string(rucksack))
		rsCount = rsCount + 1

		// New Group
		if rsCount == 3 {
			elfGroups = append(elfGroups, CurentElfGroup)
			CurentElfGroup = elfGroup{}
			rsCount = 0
		}
	}

	return elfGroups
}

func (g elfGroup) Badge() rune {

	badge := ' '

	rs1 := g.rucksacks[0]
	rs2 := g.rucksacks[1]
	rs3 := g.rucksacks[2]

	for _, char := range rs1 {
		if strings.ContainsRune(rs2, char) && strings.ContainsRune(rs3, char) {
			badge = char
		}
	}

	return badge
}

// Utility Functions

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

func priorty(item rune) int {
	value := int(item)

	if value < 97 {
		value = value - 38
	} else {
		value = value - 96
	}

	return value

}
