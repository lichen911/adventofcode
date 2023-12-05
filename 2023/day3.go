package main

import (
	// "bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getSchematic(raw_schematic string) [][]string {
	rows := strings.Split(raw_schematic, "\n")
	schematic := make([][]string, len(rows))

	for y, row := range rows {
		schematic[y] = make([]string, len(row))
		for x, col := range row {
			schematic[y][x] = string(col)
		}
	}
	return schematic
}

func main() {
	raw_schematic, err := os.ReadFile("day3_input_sample.txt")
	checkError(err)
	
	schematic := getSchematic(string(raw_schematic))
	// _ = schematic
	fmt.Println(schematic[0][0])
	
}