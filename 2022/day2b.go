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

var handPoints = map[string]int {
	"X": 0, // Lose
	"Y": 3, // Tie
	"Z": 6, // Win
}

type symbolCombination struct {
	player1 string
	player2 string
}

var strategyGuide = map[symbolCombination]int {
	{"A", "X"}: symbolPoints["Z"],
	{"A", "Y"}: symbolPoints["X"],
	{"A", "Z"}: symbolPoints["Y"],
	
	{"B", "X"}: symbolPoints["X"],
	{"B", "Y"}: symbolPoints["Y"],
	{"B", "Z"}: symbolPoints["Z"],

	{"C", "X"}: symbolPoints["Y"],
	{"C", "Y"}: symbolPoints["Z"],
	{"C", "Z"}: symbolPoints["X"],
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getScore(player1 string, player2 string) int {
	score := 0

	score += handPoints[player2]
	score += strategyGuide[symbolCombination{player1, player2}]

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