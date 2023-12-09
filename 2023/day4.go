package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type card struct {
	winningNumbers []string
	myNumbers      []string
	copies         int
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

func getWinningCardCount(winningNumbers, myNumbers []string) int {
	cardCount := 0
	for _, num := range myNumbers {
		if isWinningNumber(num, winningNumbers) {
			cardCount++
		}
	}
	return cardCount
}

func calcCardCopies(cardStack *[]card) {
	for i, currentCard := range *cardStack {
		winningCardCount := getWinningCardCount(
			currentCard.winningNumbers,
			currentCard.myNumbers,
		)

		for cardMultiple := 0; cardMultiple < currentCard.copies; cardMultiple++ {
			for j := i + 1; j <= i+winningCardCount; j++ {
				(*cardStack)[j].copies++
			}
		}
	}
}

func getTotalCardCount(cardStack []card) int {
	cardTotal := 0
	for _, currentCard := range cardStack {
		cardTotal += currentCard.copies
	}
	return cardTotal
}

func printCards(cardStack []card) {
	for i, currentCard := range cardStack {
		fmt.Println(
			"Card: ", i+1,
			"Winning Numbers:", currentCard.winningNumbers,
			"My numbers:", currentCard.myNumbers,
			"Copies:", currentCard.copies,
		)
	}
}

func main() {
	file, err := os.Open("day4_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0
	var cardStack []card

	for scanner.Scan() {
		line := scanner.Text()

		winningNumbers := strings.Fields(strings.Split(strings.Split(line, "|")[0], ":")[1])
		myNumbers := strings.Fields(strings.Split(line, "|")[1])

		cardPoints := getCardPoints(winningNumbers, myNumbers)
		totalPoints += cardPoints

		cardStack = append(cardStack,
			card{
				winningNumbers: winningNumbers,
				myNumbers:      myNumbers,
				copies:         1,
			})
	}
	fmt.Println("Total points:", totalPoints)

	calcCardCopies(&cardStack)
	printCards(cardStack)
	totalCardCount := getTotalCardCount(cardStack)
	fmt.Println("Total card count:", totalCardCount)
}
