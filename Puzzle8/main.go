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

	visibleCount, maxScore := visibleCountandScore(trees)
	// Part 1
	fmt.Printf("Part 1 - Number of visible trees: %d\n", visibleCount) //1708

	// Part 2
	fmt.Printf("Part 2 - Max tree score: %d\n", maxScore) // 1179360 to high // 1051392 WRONG

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

func visibleCountandScore(trees [][]int) (int, int) {
	visibleCount := 0
	maxTreeScore := 0

	for treeRowNum, treeRow := range trees {

		for treeColNum, tree := range treeRow {

			treeCol := getColumn(trees, treeColNum)

			// Because the arrays are zero indexed
			colTop, colTail := splitArray(treeCol, treeRowNum+1)
			rowLeft, rowRight := splitArray(treeRow, treeColNum+1)

			if biggestTree(tree, colTop) || biggestTree(tree, colTail) || biggestTree(tree, rowLeft) || biggestTree(tree, rowRight) {
				visibleCount++
			}

			currenttreeScore := scoreTree(tree, rowLeft, rowRight, colTop, colTail)
			if currenttreeScore > maxTreeScore {
				maxTreeScore = currenttreeScore
			}

		}
	}

	return visibleCount, maxTreeScore
}

func scoreTree(tree int, left []int, right []int, up []int, down []int) int {
	score := 0

	score = viewDistance(tree, reverse(left)) * viewDistance(tree, right) * viewDistance(tree, reverse(up)) * viewDistance(tree, down)

	return score
}

func viewDistance(tree int, trees []int) int {
	score := 0

	for _, t := range trees {
		if t < tree {
			score++
		} else {
			score++
			break
		}
	}

	return score
}

func reverse[T any](arr []T) []T {

	// first make a copy of the input slice.
	// For this function We want to return a NEW
	// slice that is reversed, not modify the source slice.
	// Slices make use of pointers for the values, so need to copy
	// the source slice into a new slice before doing
	// the reverse operation

	a := make([]T, len(arr))
	copy(a, arr)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
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
