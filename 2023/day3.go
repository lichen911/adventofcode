package main

import (
	// "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getSchematic(raw_schematic string) [][]rune {
	rows := strings.Split(raw_schematic, "\n")
	schematic := make([][]rune, len(rows))

	for y, row := range rows {
		schematic[y] = make([]rune, len(row))
		for x, col := range row {
			schematic[y][x] = col
		}
	}
	return schematic
}

func printSchematic(schematic [][]rune) {
	for y := range schematic {
		for x := range schematic[y] {
			fmt.Print(string(schematic[y][x]))
		}
		fmt.Println()
	}
}

func isPartNumber(schematic [][]rune, x, y, langth int) bool {
	return true
}

func findPartNumbers(schematic [][]rune) {
	partNumSum := 0
	for y := range schematic {
		x := 0
		for x < len(schematic[y]) {
			numEnd := x
			partNum := ""
			if unicode.IsDigit(schematic[y][x]) {
				onDigit := true
				partNum = string(schematic[y][x])

				for onDigit {
					if numEnd+1 < len(schematic[y]) {
						if unicode.IsDigit(schematic[y][numEnd+1]) {
							numEnd += 1
							partNum += string(schematic[y][numEnd])
						} else {
							onDigit = false
						}
					} else {
						onDigit = false
					}
				}

				fmt.Println("Part num: ", partNum)
				x = numEnd
			}

			if isPartNumber(schematic, x, y, len(partNum)) {
				var err error
				partNumSum, err = strconv.Atoi(partNum)
				checkError(err)
			}

			x++
		}
	}

	fmt.Println(partNumSum)
}

func main() {
	raw_schematic, err := os.ReadFile("day3_input_sample.txt")
	checkError(err)

	schematic := getSchematic(string(raw_schematic))
	printSchematic(schematic)
	findPartNumbers(schematic)
}
