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

var currentBag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("day2_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var gameSum int = 0
	var gamePowerSum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		invalidGame := false

		maxColorCounts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		gameId, err := strconv.Atoi(strings.Fields(strings.Split(line, ":")[0])[1])
		checkError(err)

		cubeSubsets := strings.Split(strings.Split(line, ":")[1], ";")

		for _, subset := range cubeSubsets {
			colors := strings.Split(subset, ",")
			for _, colorCount := range colors {
				count, err := strconv.Atoi(strings.Fields(colorCount)[0])
				checkError(err)
				color := strings.Fields(colorCount)[1]

				if count > currentBag[color] {
					invalidGame = true
				}

				if maxColorCounts[color] < count {
					maxColorCounts[color] = count
				}

			}
		}

		if !invalidGame {
			gameSum += gameId
		}
		invalidGame = false

		power := 1
		for _, v := range maxColorCounts {
			power *= v
		}
		gamePowerSum += power

	}

	fmt.Println("Game sum: ", gameSum)
	fmt.Println("Game power sum: ", gamePowerSum)
}
