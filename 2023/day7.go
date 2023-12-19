package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
} 

type Card struct {
	value rune
}

type Hand struct {
	cards []Card
	bid int

	// TODO: Add attribute to classify what type of hand, use enum struct
}

func (h *Hand) beats(otherHand Hand) bool {
	// Add logic to use type of hand to initially determine if one hand beats another,
	// then if two hands are equal look at individual cards for tie breakers.

	return false
}

func NewHand(cardsStr string, bid int) Hand {
	cards := make([]Card, 0)
	for _, card := range cardsStr {
		cards = append(cards, Card{value: card})
	}

	// TODO: Add logic to populate type of hand

	return Hand{
		cards: cards,
		bid: bid,
	}
}

func main() {
	file, err := os.Open("day7_input_sample.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make([]Hand, 0)


	for scanner.Scan() {
		line := scanner.Text()

		cards := strings.Fields(line)[0]
		bid, err := strconv.Atoi(strings.Fields(line)[1])
		checkError(err)

		hands = append(hands, NewHand(cards, bid))
	}

	fmt.Println(hands)
}