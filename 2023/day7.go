package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Card struct {
	value rune
}

func CardCmp(a, b Card) int {
	// Compare two Cards in the same manner as cmp.Compare
	cardFaces := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	aIndex := slices.Index(cardFaces, a.value)
	bIndex := slices.Index(cardFaces, b.value)

	return bIndex - aIndex
}

type Hand struct {
	cards    []Card
	bid      int
	handType HandType
}

func (h *Hand) setHandType() {
	// Set the handType field of the Hand struct
	cardCountMap := make(map[rune]int)

	// Populate a map containing the counts of each card
	for _, card := range h.cards {
		if _, ok := cardCountMap[card.value]; ok {
			cardCountMap[card.value]++
		} else {
			cardCountMap[card.value] = 1
		}
	}

	// Create a slice of all card counts and then sort in reverse
	cardCounts := make([]int, 0)
	for _, count := range cardCountMap {
		cardCounts = append(cardCounts, count)
	}
	slices.Sort(cardCounts)
	slices.Reverse(cardCounts)

	// Use the card counts to determine the hand type
	if cardCounts[0] == 5 {
		h.handType = FiveOfAKind
	} else if cardCounts[0] == 4 {
		h.handType = FourOfAKind
	} else if cardCounts[0] == 3 && cardCounts[1] == 2 {
		h.handType = FullHouse
	} else if cardCounts[0] == 3 {
		h.handType = ThreeOfAKind
	} else if cardCounts[0] == 2 && cardCounts[1] == 2 {
		h.handType = TwoPair
	} else if cardCounts[0] == 2 {
		h.handType = OnePair
	} else {
		h.handType = HighCard
	}
}

func NewHand(cardsStr string, bid int) Hand {
	// Initalize a new Hand
	cards := make([]Card, 0)
	for _, card := range cardsStr {
		cards = append(cards, Card{value: card})
	}

	hand := Hand{
		cards: cards,
		bid:   bid,
	}
	hand.setHandType()

	return hand
}

func HandCmp(a, b Hand) int {
	// Compare two hands following the cmp.Compare method
	if a.handType != b.handType {
		return int(a.handType - b.handType)
	} else {
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] == b.cards[i] {
				continue
			} else if true {
				return CardCmp(a.cards[i], b.cards[i])
			}
		}
	}

	return 0
}

func GetTotalWinnings(hands []Hand) int {
	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	return totalWinnings
}

func main() {
	file, err := os.Open("day7_input.txt")
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

	// fmt.Println("Pre-sorted hands:", hands)
	slices.SortFunc(hands, HandCmp)
	// fmt.Println("Post-sorted hands:", hands)

	totalwinnings := GetTotalWinnings(hands)
	fmt.Println("Total winnings:", totalwinnings)

}
