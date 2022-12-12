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

	mostCalories := findMostCalories(elfExpedition)

	fmt.Printf("The most calories carried is: %d\n", mostCalories)

}

func findMostCalories(elfExpedition []Elf) int {
	most := 0

	for _, elfRecord := range elfExpedition {
		elfTotalCals := elfRecord.TotalCalories()

		if elfTotalCals > most {
			most = elfTotalCals
		}
	}

	return most
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
