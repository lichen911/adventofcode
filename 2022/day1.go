package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func main() {
	file, err := os.Open("day1_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var all_elves []int
	current_elf := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			checkError(err)
			
			current_elf += calories
		} else {
			all_elves = append(all_elves, current_elf)
			current_elf = 0
		}
	}
	all_elves = append(all_elves, current_elf)

	sort.Ints(all_elves)
	top_three := all_elves[len(all_elves)-3:]
	fmt.Println(top_three)
	fmt.Println(sum(top_three))
}
