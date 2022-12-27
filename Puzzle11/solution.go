package Puzzle11

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int

	operation      string
	divisor        int
	throwToOnTrue  int
	throwToOnFalse int
	numInspections int
}

func RunSolution(inputFileName string) {

	monkies := parseInput(inputFileName)
	numrounds := 20

	for i := 0; i < numrounds; i++ {

		// Do a round
		for monkeyNum := range monkies {
			doMonkeyBusiness(monkeyNum, &monkies)
		}
	}

	inspections := []int{}
	for monkeyNum, monkey := range monkies {
		inspections = append(inspections, monkey.numInspections)
		fmt.Printf("Monkey %d inspected %d items\n", monkeyNum, monkey.numInspections)
	}

	sort.Ints(inspections)

	fmt.Printf("Part 1 - The sum of the two most active monkies is %d\n", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}

func parseInput(inputFileName string) []Monkey {
	monkies := []Monkey{}
	monkey := Monkey{}

	monkey.numInspections = 0

	monkeyData, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(("There was an error reading input: ") + err.Error())
		return monkies
	}
	defer monkeyData.Close()

	monkeyReader := bufio.NewReader(monkeyData)

	for {

		line, _, err := monkeyReader.ReadLine()

		monkeyLine := string(line)
		if err == io.EOF {
			// append the last monkey
			monkies = append(monkies, monkey)
			break
		}

		if monkeyLine == "" {
			//Each monkey is separated by one blank line
			monkies = append(monkies, monkey)

			// Make a new monkey
			monkey = Monkey{}
		}

		// Parse Items
		if strings.Contains(monkeyLine, "Starting items") {

			tmpArr := strings.Split(monkeyLine[strings.Index(monkeyLine, ":")+1:], ",")
			items := []int{}

			for _, item := range tmpArr {
				iItem, err := strconv.Atoi(strings.TrimSpace(item))

				if err != nil {
					fmt.Println("There was an error parsing items: " + err.Error())
					return monkies
				}

				items = append(items, iItem)
			}

			monkey.items = items

		}

		// Add operation to the monkey
		if strings.Contains(monkeyLine, "Operation") {
			monkey.operation = strings.TrimSpace(monkeyLine[strings.Index(monkeyLine, ":")+1:])
		}

		// Add divisor
		if strings.Contains(monkeyLine, "divisible") {
			div, err := strconv.Atoi(strings.TrimSpace(monkeyLine[strings.LastIndex(monkeyLine, " "):]))
			if err != nil {
				fmt.Println("There was an error parsing the divisor:  " + err.Error())
			}
			monkey.divisor = div
		}

		// Add if true action
		if strings.Contains(monkeyLine, "If true") {
			action, err := strconv.Atoi(monkeyLine[len(monkeyLine)-1:])

			if err != nil {
				fmt.Println("There was an error parsing the true action: " + err.Error())
			}
			monkey.throwToOnTrue = action
		}

		// Add if false action
		if strings.Contains(monkeyLine, "If false") {
			action, err := strconv.Atoi(monkeyLine[len(monkeyLine)-1:])

			if err != nil {
				fmt.Println("There was an error parsing the false action: " + err.Error())
			}
			monkey.throwToOnFalse = action
		}
	}

	return monkies
}

// this alters the monkies array
func doMonkeyBusiness(monkeyNum int, monkies *[]Monkey) {

	for _, item := range (*monkies)[monkeyNum].items {

		(*monkies)[monkeyNum].numInspections++
		// Perform the operation
		worryScore := (performOperation(item, (*monkies)[monkeyNum].operation) / 3)

		// Perform the test
		test := performP1Test(worryScore, (*monkies)[monkeyNum].divisor)

		destinationMonkey := -1
		if test {
			destinationMonkey = (*monkies)[monkeyNum].throwToOnTrue

		} else {
			destinationMonkey = (*monkies)[monkeyNum].throwToOnFalse
		}

		(*monkies)[destinationMonkey].items = append((*monkies)[destinationMonkey].items, worryScore)

	}

	// Monkey will have tossed every item at the end of its turn. So empty the items array
	(*monkies)[monkeyNum].items = []int{}
}

func performP1Test(worryScore int, divisor int) bool {
	result := false

	if worryScore%divisor == 0 {
		result = true
	}

	return result

}

// Parse the operation and perform it on the item
func performOperation(item int, operation string) int {
	new := 0

	rightVal := item
	var err error

	op := strings.TrimSpace(operation[strings.Index(operation, "=")+1:])
	operator := string(strings.TrimSpace(op[strings.Index(op, " "):])[0])
	right := strings.TrimSpace(op[strings.LastIndex(op, " "):])

	if right != "old" {
		rightVal, err = strconv.Atoi(right)

		if err != nil {
			fmt.Println("Error parsing right operand: " + err.Error())
		}
	}

	switch operator {
	case "+":
		new = item + rightVal

	case "*":
		new = item * rightVal

	case "-":
		new = item - rightVal

	case "/":
		new = item / rightVal

	}

	return new
}
