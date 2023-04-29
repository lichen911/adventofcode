package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(num string) int {
	integer, err := strconv.Atoi(num)
	checkError(err)
	return integer
}

func checkFullyContains(sections []string) bool {
	sec1_start := toInt(strings.Split(sections[0], "-")[0])
	sec1_end := toInt(strings.Split(sections[0], "-")[1])

	sec2_start := toInt(strings.Split(sections[1], "-")[0])
	sec2_end := toInt(strings.Split(sections[1], "-")[1])

	if sec1_start <= sec2_start && sec1_end >= sec2_end {
		return true
	} else if sec2_start <= sec1_start && sec2_end >= sec1_end {
		return true
	} else {
		return false
	}
}

func checkPartialContains(sections []string) bool {
	sec1_start := toInt(strings.Split(sections[0], "-")[0])
	sec1_end := toInt(strings.Split(sections[0], "-")[1])

	sec2_start := toInt(strings.Split(sections[1], "-")[0])
	sec2_end := toInt(strings.Split(sections[1], "-")[1])

	// The following block tests to see if there is any overlap between ranges sec1_start - sec1_end and sec2_start - sec2_end.
	if sec1_start <= sec2_start && sec1_end >= sec2_start {
		return true
	} else if sec2_start <= sec1_start && sec2_end >= sec1_start {	
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open("day4_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fullyContainCount := 0
	partialContainCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		
		assignedSections := strings.Split(line, ",")
		if checkFullyContains(assignedSections) {
			fullyContainCount++
		}

		if checkPartialContains(assignedSections) {
			partialContainCount++
		}
	}
	fmt.Println("Sections contained fully by other sections:", fullyContainCount)
	fmt.Println("Sections contained partially by other sections:", partialContainCount)
}