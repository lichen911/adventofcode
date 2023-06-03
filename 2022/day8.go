package main

import (
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

func multiply(vals ...int) int {
	total := 1
	for _, val := range vals {
		total *= val
	}
	return total
}

type Tree struct {
	height int
}

type Forest struct {
	trees [][]Tree
	// visibleTrees []Tree
}

func (f *Forest) grow(text string) {
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		trees := []Tree{}
		for _, char := range line {
			height, err := strconv.Atoi(string(char))
			checkError(err)
			trees = append(trees, Tree{height: height})
		}
		f.trees = append(f.trees, trees)
	}
}

func (f *Forest) isTreeVisible(row, col int) bool {
	height := f.trees[row][col].height
	// fmt.Println("Row:", row, "Col:", col, "Height:", height)

	if row == 0 || col == 0 || row == len(f.trees)-1 || col == len(f.trees[row])-1 {
		// fmt.Println("Is an edge, visible")
		return true
	}

	isVisible := false

	// check if tree is visible from top
	for i := row - 1; i >= 0; i-- {
		if f.trees[i][col].height >= height {
			break
		}
		if i == 0 {
			isVisible = true
		}
	}

	// check if tree is visible from bottom
	for i := row + 1; i < len(f.trees); i++ {
		if f.trees[i][col].height >= height {
			break
		}
		if i == len(f.trees)-1 {
			isVisible = true
		}
	}

	// check if tree is visible from left
	for i := col - 1; i >= 0; i-- {
		if f.trees[row][i].height >= height {
			break
		}
		if i == 0 {
			isVisible = true
		}
	}

	// check if tree is visible from right
	for i := col + 1; i < len(f.trees[row]); i++ {
		if f.trees[row][i].height >= height {
			break
		}
		if i == len(f.trees[row])-1 {
			isVisible = true
		}
	}

	return isVisible
}

func (f *Forest) countVisibleTrees() int {
	count := 0
	for row := 0; row < len(f.trees); row++ {
		for col := 0; col < len(f.trees[row]); col++ {
			if f.isTreeVisible(row, col) {
				count++
			}
		}
	}
	return count
}

func (f *Forest) getScenicScore(row, col int) int {
	height := f.trees[row][col].height
	visibleCount := []int{0, 0, 0, 0}

	// check scenery from top
	for i := row - 1; i >= 0; i-- {
		if i < 0 {
			break
		}
		visibleCount[0]++
		if f.trees[i][col].height >= height {
			break
		}
	}

	// check scenery from bottom
	for i := row + 1; i < len(f.trees); i++ {
		if i >= len(f.trees) {
			break
		}
		visibleCount[1]++
		if f.trees[i][col].height >= height {
			break
		}
	}

	// check scenery from left
	for i := col - 1; i >= 0; i-- {
		if i < 0 {
			break
		}
		visibleCount[2]++
		if f.trees[row][i].height >= height {
			break
		}
	}

	// check scenery from right
	for i := col + 1; i < len(f.trees[row]); i++ {
		if i >= len(f.trees[row]) {
			break
		}
		visibleCount[3]++
		if f.trees[row][i].height >= height {
			break
		}
	}

	return multiply(visibleCount...)
}

func (f *Forest) getHighestScenicScore() int {
	highestScore := 0
	for row := 0; row < len(f.trees); row++ {
		for col := 0; col < len(f.trees[row]); col++ {
			score := f.getScenicScore(row, col)
			if score > highestScore {
				highestScore = score
			}
		}
	}
	return highestScore
}

func (f *Forest) getTreeVisibility() string {
	visible_trees := []string{}
	for row := 0; row < len(f.trees); row++ {
		visible_trees = append(visible_trees, "")
		for col := 0; col < len(f.trees[row]); col++ {
			if f.isTreeVisible(row, col) {
				visible_trees[row] += "X"
			} else {
				visible_trees[row] += "O"
			}
		}
	}

	return strings.Join(visible_trees, "\n")
}

func NewForest(text string) *Forest {
	forest := Forest{}
	forest.grow(text)
	return &forest
}

func main() {
	content, err := os.ReadFile("day8_input.txt")
	checkError(err)

	forest := NewForest(string(content))

	// fmt.Println(forest.getTreeVisibility())
	fmt.Println("Visible tree count:", forest.countVisibleTrees())
	fmt.Println("Highest scenic score:", forest.getHighestScenicScore())
}
