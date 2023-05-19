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

type Knot struct {
	x int
	y int
}

type Rope struct {
	knots             []*Knot
	visitedTailSpaces map[[2]int]int
}

func (r *Rope) moveHeadKnot(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "U":
			r.knots[0].y++
		case "D":
			r.knots[0].y--
		case "L":
			r.knots[0].x--
		case "R":
			r.knots[0].x++
		}

		r.moveTailKnots()

		// fmt.Println(r.getLocation())
	}
}

func (r *Rope) moveTailKnots() {
	for i := 1; i < len(r.knots); i++ {
		r.moveKnot(r.knots[i-1], r.knots[i])
	}

	tail_x := r.knots[len(r.knots)-1].x
	tail_y := r.knots[len(r.knots)-1].y
	r.visitedTailSpaces[[2]int{tail_x, tail_y}]++
}

func (r *Rope) moveKnot(prev *Knot, cur *Knot) {
	if r.getDistance(prev, cur) > 1 {
		if prev.x == cur.x {
			if prev.y > cur.y {
				cur.y++
			} else {
				cur.y--
			}
		} else if prev.y == cur.y {
			if prev.x > cur.x {
				cur.x++
			} else {
				cur.x--
			}
		} else {
			if prev.x > cur.x {
				cur.x++
			} else {
				cur.x--
			}
			if prev.y > cur.y {
				cur.y++
			} else {
				cur.y--
			}
		}
	}
}

func (r Rope) getVisitedTailSpaces() int {
	return len(r.visitedTailSpaces)
}

func (r Rope) getKnotCount() int {
	return len(r.knots)
}

func (r Rope) getDistance(cur *Knot, prev *Knot) int {
	return int(math.Sqrt(math.Pow(float64(prev.x-cur.x), 2) + math.Pow(float64(prev.y-cur.y), 2)))
}

func (r Rope) getLocation(cur *Knot, prev *Knot) string {
	return fmt.Sprintf("Head: (%d, %d) Tail: (%d, %d)", prev.x, prev.y, cur.x, cur.y)
}

func NewRope(knotCount int) Rope {
	knots := []*Knot{}
	for i := 0; i < knotCount; i++ {
		knots = append(knots, &Knot{0, 0})
	}

	rope := Rope{knots, make(map[[2]int]int)}
	rope.visitedTailSpaces[[2]int{0, 0}]++
	return rope
}

func main() {
	file, err := os.Open("day9_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rope := NewRope(10)
	fmt.Println("Knot count:", rope.getKnotCount())

	for scanner.Scan() {
		line := scanner.Text()
		direction := strings.Split(line, " ")[0]
		distance, err := strconv.Atoi(strings.Split(line, " ")[1])
		checkError(err)

		fmt.Println("direction: ", direction, "distance: ", distance)
		rope.moveHeadKnot(direction, distance)
	}
	fmt.Println("Visited tail spaces: ", rope.getVisitedTailSpaces())
}
