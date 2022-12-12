package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	RPSData, err := os.Open("input.txt")
	defer RPSData.Close()

	if err != nil {
		fmt.Println("There was an error: " + err.Error())
		return
	}

	totalScorePart1 := 0
	totalScorePart2 := 0
	RPSReader := bufio.NewReader(RPSData)

	// I dont like this loop. TODO: Do this another way
	for {
		line, _, err := RPSReader.ReadLine()

		if err == io.EOF {
			break
		}

		totalScorePart1 = totalScorePart1 + scoreGame(string(line), 1)
		totalScorePart2 = totalScorePart2 + scoreGame(string(line), 2)
	}

	fmt.Printf("\nThe total score following the strategy for part 1 was: %d\n", totalScorePart1) // 12586

	fmt.Printf("The total score following the strategy for part 2 was: %d\n", totalScorePart2)

}

func scoreGame(gameData string, mode int) int {
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
	player := ""

	if mode == 1 {
		player = playerRPS[string(gameData[2])]

	} else {
		// X = Lose
		// Y = Draw
		// Z = Win

		// Make a Draw
		if string(gameData[2]) == "Y" {
			player = opponent
			fmt.Printf("Opponent: %s, Player: %s. Should be a DRAW\n", opponent, player)
		}

		// Make a Loss
		if string(gameData[2]) == "X" {

			player = chooseStrat(opponent, 0)
			fmt.Printf("Opponent: %s, Player: %s. Should be a LOSS\n", opponent, player)
		}

		// Make a Win
		if string(gameData[2]) == "Z" {
			player = chooseStrat(opponent, 1)
			fmt.Printf("Opponent: %s, Player: %s. Should be a WIN\n", opponent, player)
		}

	}

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

// 0 to pick a losing shape, anything else to pick a winning shape
func chooseStrat(opponentShape string, winlose int) string {
	shape := ""

	// This is super ugly! TODO: Fix!
	if strings.Compare(opponentShape, "Rock") == 0 && winlose == 0 {
		shape = "Scissors"
	} else if strings.Compare(opponentShape, "Rock") == 0 && winlose != 0 {
		shape = "Paper"
	}

	if strings.Compare(opponentShape, "Paper") == 0 && winlose == 0 {
		shape = "Rock"
	} else if strings.Compare(opponentShape, "Paper") == 0 && winlose != 0 {
		shape = "Scissors"
	}

	if strings.Compare(opponentShape, "Scissors") == 0 && winlose == 0 {
		shape = "Paper"
	} else if strings.Compare(opponentShape, "Scissors") == 0 && winlose != 0 {
		shape = "Rock"
	}

	return shape
}
