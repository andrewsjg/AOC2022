package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	sectionIDUpperBound int
	sectionIDLowerBound int
}

func main() {
	assignments, err := makeAssignments("input.txt")

	if err != nil {
		fmt.Println("There was an error: " + err.Error())
	}

	containedCount := 0
	overlapCount := 0

	for _, assignmentPair := range assignments {
		if checkContained(assignmentPair) {
			containedCount++
		}

		if checkOverlap(assignmentPair) {
			overlapCount++
		}
	}

	fmt.Printf("The total contained assignments is: %d Total overlapping assignments is: %d\n", containedCount, overlapCount)
}

// Return an array of assignment pairs from the input
func makeAssignments(input string) ([][]assignment, error) {
	assignments := [][]assignment{}

	assignmentData, err := os.Open(input)
	defer assignmentData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return assignments, err

	}

	assignmentReader := bufio.NewReader(assignmentData)

	for {
		assignmentPairBytes, _, err := assignmentReader.ReadLine()

		if err == io.EOF {
			break
		}

		aPair := string(assignmentPairBytes)

		if len(assignmentPairBytes) > 0 {
			assignmentPairs := []assignment{}

			pairs := strings.Split(aPair, ",")

			p1 := strings.Split(pairs[0], "-")
			p2 := strings.Split(pairs[1], "-")

			p1Upper, p1Lower := getUpperandLowerBound(p1)
			p2Upper, p2Lower := getUpperandLowerBound(p2)

			pair1Assignment := assignment{sectionIDUpperBound: p1Upper, sectionIDLowerBound: p1Lower}
			pair2Assignment := assignment{sectionIDUpperBound: p2Upper, sectionIDLowerBound: p2Lower}

			assignmentPairs = append(assignmentPairs, pair1Assignment)
			assignmentPairs = append(assignmentPairs, pair2Assignment)

			assignments = append(assignments, assignmentPairs)
		}

	}

	return assignments, nil
}

// Check if either assignment contains the other
func checkContained(assignments []assignment) bool {
	contained := false

	if len(assignments) > 0 {
		assignment1 := assignments[0]
		assignment2 := assignments[1]

		// assignment 1 is contained by assignment 2
		if assignment1.sectionIDLowerBound >= assignment2.sectionIDLowerBound && assignment1.sectionIDUpperBound <= assignment2.sectionIDUpperBound {
			contained = true
		}

		// assignment 2 is contained by assignment 1
		if assignment2.sectionIDLowerBound >= assignment1.sectionIDLowerBound && assignment2.sectionIDUpperBound <= assignment1.sectionIDUpperBound {
			contained = true
		}

	}

	return contained
}

// Check if an assignment pair overlaps
func checkOverlap(assignments []assignment) bool {
	overlap := false

	if len(assignments) > 0 {
		assignment1 := assignments[0]
		assignment2 := assignments[1]

		if checkContained(assignments) {
			overlap = true
		}

		if assignment1.sectionIDLowerBound >= assignment2.sectionIDLowerBound && assignment1.sectionIDLowerBound <= assignment2.sectionIDUpperBound {

			overlap = true
		}

		// Is this check redundant?
		if assignment2.sectionIDLowerBound >= assignment1.sectionIDLowerBound && assignment2.sectionIDLowerBound <= assignment1.sectionIDUpperBound {

			overlap = true
		}

		if assignment1.sectionIDUpperBound <= assignment2.sectionIDUpperBound && assignment1.sectionIDUpperBound >= assignment2.sectionIDLowerBound {

			overlap = true
		}

		// Is this check redundant?
		if assignment2.sectionIDUpperBound <= assignment1.sectionIDUpperBound && assignment2.sectionIDUpperBound >= assignment1.sectionIDLowerBound {

			overlap = true
		}
	}

	return overlap
}

// Convienience function to return the upper and lower bound as integers
// based on the string data from the input
func getUpperandLowerBound(assignment []string) (int, int) {
	upper := 0
	lower := 0

	lower, err := strconv.Atoi(assignment[0])
	if err != nil {
		return 0, 0
	}

	upper, err = strconv.Atoi(assignment[1])
	if err != nil {
		return 0, 0
	}

	return upper, lower

}
