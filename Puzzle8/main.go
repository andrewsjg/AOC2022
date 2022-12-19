package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	trees := readTrees("input.txt")

	// Part 1
	fmt.Printf("Part 1 - Number of visible trees: %d\n", visibleCount(trees))
}

// This seems unecessary, but my brain hurts working out a better way at
// the moment
func getColumn(trees [][]int, columnNumber int) []int {
	treeCol := []int{}

	for _, treeRow := range trees {
		for rowPos, tree := range treeRow {
			if rowPos == columnNumber {
				treeCol = append(treeCol, tree)
			}
		}
	}
	return treeCol
}

func visibleCount(trees [][]int) int {
	visibleCount := 0

	for treeRowNum, treeRow := range trees {

		for treeColNum, tree := range treeRow {

			treeCol := getColumn(trees, treeColNum)

			// Because the arrays are zero indexed
			colTop, colTail := splitArray(treeCol, treeRowNum+1)
			rowLeft, rowRight := splitArray(treeRow, treeColNum+1)

			//fmt.Printf("biggest Tree ColTop: %t biggest Tree ColTail: %t biggest Tree Rowleft: %t biggest Tree RowRight: %t\n", biggestTree(tree, colTop), biggestTree(tree, colTail), biggestTree(tree, rowLeft), biggestTree(tree, rowRight))

			if biggestTree(tree, colTop) || biggestTree(tree, colTail) || biggestTree(tree, rowLeft) || biggestTree(tree, rowRight) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

// Why not make it generic!
func splitArray[T any](arr []T, splitAt int) ([]T, []T) {
	part1 := []T{}
	part2 := []T{}

	part1 = arr[:splitAt]
	if len(part1) > 0 {
		part1 = part1[:len(part1)-1]
	}
	part2 = arr[splitAt:]

	return part1, part2
}

func biggestTree(treeToCheck int, treeRowCol []int) bool {
	biggest := true

	for _, tree := range treeRowCol {
		if tree >= treeToCheck {
			biggest = false
		}
	}

	return biggest
}

func readTrees(input string) [][]int {

	trees := [][]int{}

	// TODO: Really should just make this block part of a utility package

	treeData, err := os.Open(input)
	defer treeData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return trees

	}

	treeReader := bufio.NewReader(treeData)

	for {

		treeLine, _, err := treeReader.ReadLine()
		treeRow := []int{}

		if err == io.EOF {
			break
		}

		for _, treeByte := range treeLine {

			tree, err := strconv.Atoi(string(treeByte))

			if err != nil {
				fmt.Println("There was an error converting the tree string to an int: " + err.Error())
				return trees
			}
			treeRow = append(treeRow, tree)

		}

		trees = append(trees, treeRow)

	}
	return trees
}
