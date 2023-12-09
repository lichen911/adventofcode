package main

import (
	"bufio"
	"fmt"
	"os"

	// "strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isWinningNumber(myNum string, winningNumbers []string) bool {
	for _, num := range winningNumbers {
		if myNum == num {
			return true
		}
	}
	return false
}

func getCardPoints(winningNumbers, myNumbers []string) int {
	points := 0
	for _, num := range myNumbers {
		if isWinningNumber(num, winningNumbers) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func main() {
	file, err := os.Open("day4_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()

		winningNumbers := strings.Fields(strings.Split(strings.Split(line, "|")[0], ":")[1])
		myNumbers := strings.Fields(strings.Split(line, "|")[1])

		cardPoints := getCardPoints(winningNumbers, myNumbers)
		fmt.Println("Card points:", cardPoints)
		totalPoints += cardPoints

	}
	fmt.Println("Total points:", totalPoints)
}
