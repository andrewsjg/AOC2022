package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Star struct {
	Calories int
}

type Elf struct {
	Name      string
	Inventory []Star
}

func main() {

	elfExpedition, err := makeElfExpedition("input.txt")

	if err != nil {
		fmt.Println("There was an error: " + err.Error())
	}

	mostCalories, _ := findMostCalories(elfExpedition)

	fmt.Printf("The most calories carried is: %d\n", mostCalories)
	fmt.Printf("The top 3 elves are carrying %d calories\n", totalTopThree(elfExpedition))

}

func findMostCalories(elfExpedition []Elf) (mostCals int, elfNum int) {

	for idx, elfRecord := range elfExpedition {
		elfTotalCals := elfRecord.TotalCalories()

		if elfTotalCals > mostCals {
			mostCals = elfTotalCals
			elfNum = idx
		}
	}

	return mostCals, elfNum
}

func totalTopThree(elfExpedition []Elf) int {
	total := 0

	for i := 0; i < 3; i++ {
		elfTotal, elfNum := findMostCalories(elfExpedition)

		total = total + elfTotal

		elfExpedition = RemoveIndex(elfExpedition, elfNum)
	}

	return total
}

func RemoveIndex(s []Elf, index int) []Elf {
	ret := make([]Elf, 0)
	ret = append(ret, s[:index]...)

	return append(ret, s[index+1:]...)
}

func makeElfExpedition(filename string) ([]Elf, error) {
	elfExpedition := []Elf{}
	elfData, err := os.Open(filename)
	defer elfData.Close()

	if err != nil {
		return elfExpedition, err
	}

	elfreader := bufio.NewReader(elfData)

	anElf := Elf{}
	elfNum := 1

	anElf.Name = "Elf" + strconv.Itoa(elfNum)
	for {
		line, _, err := elfreader.ReadLine()

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			elfExpedition = append(elfExpedition, anElf)
			anElf = Elf{}
			elfNum = elfNum + 1
			anElf.Name = "Elf" + strconv.Itoa(elfNum)
		} else {

			calories, err := strconv.Atoi(string(line))
			if err != nil {
				calories = 0
			}

			aStar := Star{Calories: calories}
			anElf.Inventory = append(anElf.Inventory, aStar)

		}
	}

	return elfExpedition, nil
}

func (e *Elf) TotalCalories() int {
	totalCals := 0

	for _, cals := range e.Inventory {
		totalCals = totalCals + cals.Calories
	}

	return totalCals
}
