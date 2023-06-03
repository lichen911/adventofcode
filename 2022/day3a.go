package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func findCommonItem(string1, string2 string) rune {
	// Find the first common item between two strings
	for _, char1 := range string1 {
		for _, char2 := range string2 {
			if char1 == char2 {
				return char1
			}
		}
	}
	return rune(0)
}

func findBadge(badgeGroup []string) rune {
	// Find the first common item between three strings
	for _, char1 := range badgeGroup[0] {
		for _, char2 := range badgeGroup[1] {
			for _, char3 := range badgeGroup[2] {
				if char1 == char2 && char2 == char3 {
					return char1
				}
			}
		}
	}
	return rune(0)
}

func getPriority(char rune) int {
	// Lowercase letters have priority 1-26, uppercase letters have priority 27-52
	var priority int
	if unicode.IsLower(char) {
		priority = int(char) - 96
	} else {
		priority = int(char) - 38
	}
	return priority
}

func main() {
	file, err := os.Open("day3_input.txt")
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPriority := 0
	badgePriority := 0
	lineCount := 0
	badgeGroup := make([]string, 3, 3)
	for scanner.Scan() {
		line := scanner.Text()

		part1 := line[0 : len(line)/2]
		part2 := line[len(line)/2:]
		totalPriority += getPriority(findCommonItem(part1, part2))

		if lineCount < 2 {
			badgeGroup[lineCount] = line
		} else {
			badgeGroup[lineCount] = line
			badgePriority += getPriority(findBadge(badgeGroup))
			lineCount = 0
			continue
		}
		lineCount++
	}

	fmt.Printf("Sum of all priority values: %v\n", totalPriority)
	fmt.Printf("Sum of all badge priority values: %v\n", badgePriority)
}
