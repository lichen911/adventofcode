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

func isPartNumber(schematic [][]rune, x, y, length int) bool {
	var xStart, xEnd int

	// test top - y-1, x-1 to x+length+1
	if y-1 >= 0 {
		if x-1 < 0 {
			xStart = 0
		} else {
			xStart = x-1
		}

		if x+length > len(schematic[y-1])-1 {
			xEnd = len(schematic[y-1])-1
		} else {
			xEnd = x+length
		}

		for _, r := range schematic[y-1][xStart:xEnd+1] {
			if string(r) != "." {
				return true
			}
		}
	}

	// test sides - y, x-1 and x+length
	if x-1 >= 0 {
		if string(schematic[y][x-1]) != "." {
			return true
		}
	}
	if x+length <= len(schematic[y])-1 {
		if string(schematic[y][x+length]) != "." {
			return true
		}
	}

	// test bottom y+1, x-1 to x+length+1
	if y+1 <= len(schematic)-1 {
		if x-1 < 0 {
			xStart = 0
		} else {
			xStart = x-1
		}

		if x+length > len(schematic[y+1])-1 {
			xEnd = len(schematic[y+1])-1
		} else {
			xEnd = x+length
		}

		for _, r := range schematic[y+1][xStart:xEnd+1] {
			if string(r) != "." {
				return true
			}
		}
	}

	return false
}


func findPartNumbers(schematic [][]rune) int {
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

				if isPartNumber(schematic, x, y, len(partNum)) {
					partNumInt, err := strconv.Atoi(partNum)
					checkError(err)

					partNumSum += partNumInt
				}

				x = numEnd
			}

			x++
		}
	}

	return partNumSum
}

func main() {
	raw_schematic, err := os.ReadFile("day3_input.txt")
	checkError(err)

	schematic := getSchematic(string(raw_schematic))
	// printSchematic(schematic)
	partNumSum := findPartNumbers(schematic)
	fmt.Println("Part num sum:", partNumSum)
}
