package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

var symbolPoints = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

type symbolCombination struct {
	player1 string
	player2 string
}

var handPoints = map[symbolCombination]int {
	// Tie
	{"A", "X"}: 3,
	{"B", "Y"}: 3,
	{"C", "Z"}: 3,

	// Player 1 wins
	{"A", "Z"}: 0,
	{"B", "X"}: 0,
	{"C", "Y"}: 0,

	// Player 2 wins
	{"A", "Y"}: 6,
	{"B", "Z"}: 6,
	{"C", "X"}: 6,
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getScore(player1 string, player2 string) int {
	score := 0

	score += symbolPoints[player2]
	score += handPoints[symbolCombination{player1, player2}]

	return score
} 

func main() {
	file, err := os.Open("day2_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		strategy := strings.Split(scanner.Text(), " ")
		// fmt.Println(strategy)

		score := getScore(strategy[0], strategy[1])
		totalScore += score
	}

	fmt.Println(totalScore)
}