package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	RPSData, err := os.Open("input.txt")
	defer RPSData.Close()

	if err != nil {
		fmt.Println("There was an error: " + err.Error())
		return
	}

	totalScore := 0
	RPSReader := bufio.NewReader(RPSData)

	// I dont like this loop. TODO: Do this another way
	for {
		line, _, err := RPSReader.ReadLine()

		if err == io.EOF {
			break
		}

		totalScore = totalScore + scoreGame(string(line))
	}

	fmt.Printf("The total score following the strategy was: %d\n", totalScore)

}

func scoreGame(gameData string) int {
	// A,X = Rock
	// B,Y = Paper
	// C,Z = Scissors

	opponentRPS := make(map[string]string)
	playerRPS := make(map[string]string)
	ShapeValue := make(map[string]int)

	ShapeValue["Rock"] = 1
	ShapeValue["Paper"] = 2
	ShapeValue["Scissors"] = 3

	opponentRPS["A"] = "Rock"
	opponentRPS["B"] = "Paper"
	opponentRPS["C"] = "Scissors"

	playerRPS["X"] = "Rock"
	playerRPS["Y"] = "Paper"
	playerRPS["Z"] = "Scissors"

	opponent := opponentRPS[string(gameData[0])]
	player := playerRPS[string(gameData[2])]

	score := ShapeValue[player]

	// Draw
	if opponent == player {
		score = score + 3

	} else {

		if doesThePlayerWin(player, opponent) {
			// Win
			score = score + 6
		}
	}

	// fmt.Printf("Opponent: %s, Player: %s Result: %s, Score: %d\n", opponent, player, result, score)

	return score
}

func doesThePlayerWin(player string, opponent string) bool {
	winner := true

	if player == "Rock" && opponent == "Paper" {
		winner = false
	}

	if player == "Paper" && opponent == "Scissors" {
		winner = false
	}

	if player == "Scissors" && opponent == "Rock" {
		winner = false
	}

	return winner
}
