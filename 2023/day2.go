package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


var currentBag = map[string]int {
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	file, err := os.Open("day2_input_sample.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// var gameSum int = 0

	for scanner.Scan() {
		line := scanner.Text()

		gameId, err := strconv.Atoi(strings.Fields(strings.Split(line, ":")[0])[1])
		checkError(err)

		fmt.Println(gameId)

		cubeSubsets := strings.Split(strings.Split(line, ":")[1], ";")
		for _, subset := range cubeSubsets {
			colors := strings.Split(subset, ",")
			for _, color := range colors {
				fmt.Println(color)
			}

		}
	}

}