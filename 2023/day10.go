package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Coord struct {
	x, y int
}

type Notes struct {
	diagram  [][]rune
	startPos Coord
}

func NewNotes(input string) Notes {
	rows := strings.Split(input, "\n")
	output := make([][]rune, len(rows))

	var startX, startY int

	for y, row := range rows {
		output[y] = make([]rune, len(row))
		for x, col := range row {
			output[y][x] = col

			if col == 'S' {
				startX = x
				startY = y
			}
		}
	}
	return Notes{
		diagram:  output,
		startPos: Coord{x: startX, y: startY},
	}
}

func (pipeNotes *Notes) PrintNotes() {
	for y := range pipeNotes.diagram {
		for x := range pipeNotes.diagram[y] {
			fmt.Print(string(pipeNotes.diagram[y][x]))
		}
		fmt.Println()
	}
	fmt.Println("\nStart X:", pipeNotes.startPos.x, "Y:", pipeNotes.startPos.y)
}

func (pipeNotes *Notes) GetPipe(coord Coord) (rune, error) {
	if coord.x < 0 || coord.y < 0 ||
		coord.x >= len(pipeNotes.diagram[0]) || coord.y >= len(pipeNotes.diagram) {
		return 0, errors.New("Invalid coordinates")
	}

	return pipeNotes.diagram[coord.y][coord.x], nil
}

func (pipeNotes *Notes) GetLoopPath() []Coord {
	previous := Coord{x: pipeNotes.startPos.x, y: pipeNotes.startPos.y}
	current := getNextPosition(
		previous,
		previous,
		*pipeNotes,
	)

	loopPath := make([]Coord, 0)
	loopPath = append(loopPath, pipeNotes.startPos, current)
	for current != pipeNotes.startPos {
		newCurrent := getNextPosition(
			previous,
			current,
			*pipeNotes,
		)
		previous = current
		current = newCurrent

		if current != pipeNotes.startPos {
			loopPath = append(loopPath, current)
		}
	}
	return loopPath
}

func getNextPosition(previous, current Coord, pipeNotes Notes) Coord {
	/*
		| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
		. is ground; there is no pipe in this tile.
		S is the starting position of the animal
	*/
	var next1, next2 Coord
	switch pipe, _ := pipeNotes.GetPipe(current); pipe {
	case '|':
		next1.x, next2.x = current.x, current.x
		next1.y, next2.y = current.y+1, current.y-1
	case '-':
		next1.x, next2.x = current.x+1, current.x-1
		next1.y, next2.y = current.y, current.y
	case 'L':
		next1.x, next2.x = current.x+1, current.x
		next1.y, next2.y = current.y, current.y-1
	case 'J':
		next1.x, next2.x = current.x-1, current.x
		next1.y, next2.y = current.y, current.y-1
	case '7':
		next1.x, next2.x = current.x-1, current.x
		next1.y, next2.y = current.y, current.y+1
	case 'F':
		next1.x, next2.x = current.x+1, current.x
		next1.y, next2.y = current.y, current.y+1
	case 'S':
		northCoord := Coord{x: current.x, y: current.y - 1}
		northPipe, err := pipeNotes.GetPipe(northCoord)
		if err == nil && (northPipe == 'F' || northPipe == '7' || northPipe == '|') {
			return northCoord
		}

		eastCoord := Coord{x: current.x + 1, y: current.y}
		eastPipe, err := pipeNotes.GetPipe(eastCoord)
		if err == nil && (eastPipe == 'J' || eastPipe == '7' || eastPipe == '-') {
			return eastCoord
		}

		southCoord := Coord{x: current.x, y: current.y + 1}
		southPipe, err := pipeNotes.GetPipe(southCoord)
		if err == nil && (southPipe == 'L' || southPipe == 'J' || southPipe == '|') {
			return southCoord
		}

		westCoord := Coord{x: current.x - 1, y: current.y}
		westPipe, err := pipeNotes.GetPipe(westCoord)
		if err == nil && (westPipe == 'L' || westPipe == 'F' || westPipe == '-') {
			return westCoord
		}
	}

	if previous.x == next1.x && previous.y == next1.y {
		return next2
	} else {
		return next1
	}
}

func main() {
	raw_input, err := os.ReadFile("day10_input.txt")
	checkError(err)

	pipeNotes := NewNotes(string(raw_input))
	// pipeNotes.PrintNotes()

	loopPath := pipeNotes.GetLoopPath()
	// fmt.Println("Loop path:", loopPath)

	loopMiddle := len(loopPath) / 2
	fmt.Println("Loop middle:", loopMiddle)
}
