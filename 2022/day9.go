package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type RopeEnd struct {
	x int
	y int
}

type Rope struct {
	head   RopeEnd
	tail   RopeEnd
	length int
	visitedTailSpaces map[[2]int]int
}

func (r *Rope) moveHead(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "U":
			r.head.y++
		case "D":
			r.head.y--
		case "L":
			r.head.x--
		case "R":
			r.head.x++
		}

		r.moveTail()

		fmt.Println(r.getLocation())
	}
}

func (r *Rope) moveTail() {
	if r.getDistance() > r.length {
		if r.head.x == r.tail.x {
			if r.head.y > r.tail.y {
				r.tail.y++
			} else {
				r.tail.y--
			}
		} else if r.head.y == r.tail.y {
			if r.head.x > r.tail.x {
				r.tail.x++
			} else {
				r.tail.x--
			}
		} else {
			if r.head.x > r.tail.x {
				r.tail.x++
			} else {
				r.tail.x--
			}
			if r.head.y > r.tail.y {
				r.tail.y++
			} else {
				r.tail.y--
			}
		}
		r.visitedTailSpaces[[2]int{r.tail.x, r.tail.y}]++
	}
}

func (r Rope) getVisitedTailSpaces() int {
	return len(r.visitedTailSpaces)
}

func (r Rope) getDistance() int {
	return int(math.Sqrt(math.Pow(float64(r.head.x-r.tail.x), 2) + math.Pow(float64(r.head.y-r.tail.y), 2)))
}

func (r Rope) getLocation() string {
	return fmt.Sprintf("Head: (%d, %d) Tail: (%d, %d)", r.head.x, r.head.y, r.tail.x, r.tail.y)
}

func NewRope(length int) Rope {
	rope := Rope{RopeEnd{0, 0}, RopeEnd{0, 0}, length, make(map[[2]int]int)}
	rope.visitedTailSpaces[[2]int{0, 0}]++
	return rope
}

func main() {
	file, err := os.Open("day9_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rope := NewRope(1)
	fmt.Println(rope.getLocation())

	for scanner.Scan() {
		line := scanner.Text()
		direction := strings.Split(line, " ")[0]
		distance, err := strconv.Atoi(strings.Split(line, " ")[1])
		checkError(err)

		fmt.Println("direction: ", direction, "distance: ", distance)
		rope.moveHead(direction, distance)

		// fmt.Println("Distance between head and tail: ", rope.getDistance())
	}
	fmt.Println("Visited tail spaces: ", rope.getVisitedTailSpaces())
}
