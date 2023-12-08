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

func updateGearMap(gear rune, gearLocX, gearLocY int, gearMap *map[[2]int][]int, partNum int) {
	if string(gear) == "*" {
		gearMapKey := [2]int{gearLocX, gearLocY}
		if _, ok := (*gearMap)[gearMapKey]; ok {
			(*gearMap)[gearMapKey] = append((*gearMap)[gearMapKey], partNum)
		} else {
			(*gearMap)[gearMapKey] = []int{partNum}
		}
	}
}

func isPartNumber(schematic [][]rune, x, y, length int, gearMap *map[[2]int][]int, partNum int) bool {
	var xStart, xEnd int
	var curX, curY int

	// test top - y-1, x-1 to x+length+1
	curY = y-1
	if curY >= 0 {
		if x-1 < 0 {
			xStart = 0
		} else {
			xStart = x-1
		}

		if x+length > len(schematic[curY])-1 {
			xEnd = len(schematic[curY])-1
		} else {
			xEnd = x+length
		}

		for curX, r := range schematic[curY][xStart:xEnd+1] {
			updateGearMap(r, curX+xStart, curY, gearMap, partNum)
			if string(r) != "." {
				return true
			}
		}
	}

	// test sides - y, x-1 and x+length
	curX = x-1
	if curX >= 0 {
		r := schematic[y][curX]
		updateGearMap(r, curX, y, gearMap, partNum)
		if string(r) != "." {
			return true
		}
	}

	curX = x+length
	if curX <= len(schematic[y])-1 {
		r := schematic[y][curX]
		updateGearMap(r, curX, y, gearMap, partNum)
		if string(r) != "." {
			return true
		}
	}

	// test bottom y+1, x-1 to x+length+1
	curY = y+1
	if curY <= len(schematic)-1 {
		if x-1 < 0 {
			xStart = 0
		} else {
			xStart = x-1
		}

		if x+length > len(schematic[curY])-1 {
			xEnd = len(schematic[curY])-1
		} else {
			xEnd = x+length
		}

		for curX, r := range schematic[curY][xStart:xEnd+1] {
			updateGearMap(r, curX+xStart, curY, gearMap, partNum)
			if string(r) != "." {
				return true
			}
		}
	}

	return false
}


func findPartNumbers(schematic [][]rune, gearMap *map[[2]int][]int) int {
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

				partNumInt, err := strconv.Atoi(partNum)
				checkError(err)
				if isPartNumber(schematic, x, y, len(partNum), gearMap, partNumInt) {
					partNumSum += partNumInt
				}

				x = numEnd
			}

			x++
		}
	}

	return partNumSum
}

func getGearRatioSum(gearMap map[[2]int][]int) int {
	gearRatioSum := 0
	for _, partNum := range gearMap {
		if len(partNum) == 2 {
			gearRatioSum += partNum[0] * partNum[1]
		}
	}

	return gearRatioSum
}

func main() {
	raw_schematic, err := os.ReadFile("day3_input.txt")
	checkError(err)

	schematic := getSchematic(string(raw_schematic))
	gearMap := make(map[[2]int][]int)

	// printSchematic(schematic)
	partNumSum := findPartNumbers(schematic, &gearMap)
	fmt.Println("Part num sum:", partNumSum)

	// fmt.Println("Gear map:", gearMap)
	gearRatioSum := getGearRatioSum(gearMap)
	fmt.Println("Gear ratio sum:", gearRatioSum)
}
