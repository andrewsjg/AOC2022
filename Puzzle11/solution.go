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

	monkeys := parseInput(inputFileName)
	numrounds := 10000

	for i := 0; i < numrounds; i++ {

		// Do a round
		for monkeyNum := range monkeys {
			doMonkeyBusiness(monkeyNum, &monkeys, true)
		}
	}

	inspections := []int{}
	for monkeyNum, monkey := range monkeys {
		inspections = append(inspections, monkey.numInspections)
		fmt.Printf("Monkey %d inspected %d items\n", monkeyNum, monkey.numInspections)
	}

	sort.Ints(inspections)

	fmt.Printf("Part 1 - The sum of the two most active monkeys is %d\n", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}

func parseInput(inputFileName string) []Monkey {
	monkeys := []Monkey{}
	monkey := Monkey{}

	monkey.numInspections = 0

	monkeyData, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(("There was an error reading input: ") + err.Error())
		return monkeys
	}
	defer monkeyData.Close()

	monkeyReader := bufio.NewReader(monkeyData)

	for {

		line, _, err := monkeyReader.ReadLine()

		monkeyLine := string(line)
		if err == io.EOF {
			// append the last monkey
			monkeys = append(monkeys, monkey)
			break
		}

		if monkeyLine == "" {
			//Each monkey is separated by one blank line
			monkeys = append(monkeys, monkey)

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
					return monkeys
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

	return monkeys
}

// this alters the monkeys array
func doMonkeyBusiness(monkeyNum int, monkeys *[]Monkey, part2 bool) {

	// This is horribly ineffiecient to do on every round. But it was the easiest
	// way to hack into my P1 solution and I couldnt be bothered refactoring

	// Need a common multiple. Since the divisors all appear to be primne
	// The product of all divisors is always equal to the least common multiple
	// needed Reddit's help for this one.

	lcm := 1
	for _, monkey := range *monkeys {
		lcm *= monkey.divisor
	}

	for _, item := range (*monkeys)[monkeyNum].items {

		(*monkeys)[monkeyNum].numInspections++

		worryScore := performOperation(item, (*monkeys)[monkeyNum].operation) / 3

		if part2 {
			worryScore = (performOperation(item, (*monkeys)[monkeyNum].operation)) % lcm
		}

		destinationMonkey := -1
		if worryScore%(*monkeys)[monkeyNum].divisor == 0 {
			destinationMonkey = (*monkeys)[monkeyNum].throwToOnTrue

		} else {
			destinationMonkey = (*monkeys)[monkeyNum].throwToOnFalse
		}

		(*monkeys)[destinationMonkey].items = append((*monkeys)[destinationMonkey].items, worryScore)

	}

	// Monkey will have tossed every item at the end of its turn. So empty the items array
	(*monkeys)[monkeyNum].items = []int{}
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
