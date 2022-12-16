package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Create custom types to which we can add methods like push and pop
type stack string
type containerStack map[int]stack

func main() {

	// Part 1
	stacks := makeStacks("input.txt")
	stacks.executeMoves("input.txt", true)

	topCrates := ""
	for i := 1; i <= len(stacks); i++ {
		topCrates = topCrates + string(stacks[i][0])
	}
	fmt.Println("Part 1 - Top Crates: " + topCrates) //SVFDLGLWV

	// Part 2
	stacks = makeStacks("input.txt")
	stacks.executeMoves("input.txt", false)

	topCrates = ""
	for i := 1; i <= len(stacks); i++ {
		topCrates = topCrates + string(stacks[i][0])
	}
	fmt.Println("Part 2 - Top Crates: " + topCrates) //SVFDLGLWV
}

/*

NOTE - Detail from the input for reference
-----

Stack alignment in output:

stack 1 = 1
stack 2 = 5
stack 3 = 9
stack 4 = 13
stack 5 = 17
stack 6 = 21
stack 7 = 26
stack 8 = 29
stack 9 = 33

Stacks from the input file:

[N]         [C]     [Z]
[Q] [G]     [V]     [S]         [V]
[L] [C]     [M]     [T]     [W] [L]
[S] [H]     [L]     [C] [D] [H] [S]
[C] [V] [F] [D]     [D] [B] [Q] [F]
[Z] [T] [Z] [T] [C] [J] [G] [S] [Q]
[P] [P] [C] [W] [W] [F] [W] [J] [C]
[T] [L] [D] [G] [P] [P] [V] [N] [R]
 1   2   3   4   5   6   7   8   9

*/

// Read the input and parse the cargo into a map type
func makeStacks(input string) containerStack {
	stacks := make(map[int]stack)

	conatinerData, err := os.Open(input)
	defer conatinerData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return stacks

	}

	containerReader := bufio.NewReader(conatinerData)

	for {
		containerInfo, _, err := containerReader.ReadLine()

		if err == io.EOF {
			break
		}

		if len(containerInfo) == 0 {
			break
		}

		containerLine := string(containerInfo)

		stackCount := 1
		stackIncrement := 1
		// Only parse lines that look like containers. This ignores the last line of the container
		// input which represents the stack number
		if strings.Contains(containerLine, "[") {
			for stackIncrement < len(containerLine) {
				cargo := string(containerLine[stackIncrement])

				// Dont add empty slots into the map
				if cargo != " " {
					// The cargo value has to be cast into a stack to allow the append to work
					stacks[stackCount] = stacks[stackCount] + stack(cargo)
				}

				// The cargo ID's are at 4 character intervals
				stackIncrement = stackIncrement + 4
				stackCount++
			}
		}
	}

	return stacks
}

// Execute the moves encoded in the input on the containerStack
func (cs *containerStack) executeMoves(input string, partOne bool) {
	moveData, err := os.Open(input)
	defer moveData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return

	}

	moveReader := bufio.NewReader(moveData)

	for {
		bMove, _, err := moveReader.ReadLine()

		if err == io.EOF {
			break
		}

		// Ignore anything that doesnt contain a move instruction
		move := string(bMove)
		if strings.Contains(move, "move") {
			amount, from, to := parseMove(move)
			//fmt.Printf("Moving %d containers from %d to %d\n", amount, from, to)

			if partOne {
				for i := 1; i <= amount; i++ {
					cs.moveContainer(from, to)
				}
			} else {
				cs.moveMultipleContainers(amount, from, to)
			}

		}
	}
}

// Read a move string and return a set of integers for the amount to move,
// the source stack and the destination stack
func parseMove(moveInstruction string) (amountToMove int, from int, to int) {

	// Need to check for double digit number of containers to move
	amountStr1 := string(moveInstruction[5])
	amountStr2 := string(moveInstruction[6])

	_, err := strconv.Atoi(amountStr2)

	if err == nil {
		amountToMove, err = strconv.Atoi(amountStr1 + amountStr2)
		if err != nil {
			fmt.Println("AMOUNT ERROR")
			amountToMove = 0
		}
	} else {
		amountToMove, err = strconv.Atoi(string(moveInstruction[5]))

		if err != nil {
			fmt.Println("AMOUNT ERROR")
			amountToMove = 0
		}
	}

	indexofFromValue := strings.Index(moveInstruction, "from") + 5
	indexofToValue := strings.Index(moveInstruction, "to") + 3

	to, err = strconv.Atoi(string(moveInstruction[indexofToValue]))

	// TODO: Should really propagate an error here
	if err != nil {
		fmt.Println("TO ERROR")
		to = 0
	}

	from, err = strconv.Atoi(string(moveInstruction[indexofFromValue]))

	// TODO: Should really propagate an error here
	if err != nil {
		fmt.Println("FROM ERROR")
		from = 0
	}
	//fmt.Printf("Amount to move: %d from: %d to: %d\n", amountToMove, from, to)
	return amountToMove, from, to
}

// Move a container from one stack to another
func (cs *containerStack) moveContainer(fromStack int, toStack int) {
	tmpStack := *cs

	from := tmpStack[fromStack]
	to := tmpStack[toStack]

	to.push(from.pop())

	tmpStack[fromStack] = from
	tmpStack[toStack] = to

	*cs = tmpStack
}

// Move a container from one stack to another
func (cs *containerStack) moveMultipleContainers(amount, fromStack int, toStack int) {
	tmpStack := *cs

	from := tmpStack[fromStack]
	to := tmpStack[toStack]

	to.push(from.popMultiple(amount))

	tmpStack[fromStack] = from
	tmpStack[toStack] = to

	*cs = tmpStack
}

// Stack operations

// Pop a value off the stack (as represented by a string).
// Note this method mutates the stack
func (s *stack) pop() string {

	popped := ""

	if len(*s) > 0 {
		popped = string((*s)[0])
		*s = (*s)[1:]
	}

	return popped
}

// Push a value onto the stack (as represented by a string).
// Note this method mutates the stack
// Note push works for one item or multiple!
func (s *stack) push(item string) {
	*s = stack(item) + (*s)
}

func (s *stack) popMultiple(noOfItemsToPop int) string {

	popped := ""

	if len(*s) > 0 {
		popped = string((*s)[:noOfItemsToPop])

		*s = (*s)[noOfItemsToPop:]

	}

	return popped
}
