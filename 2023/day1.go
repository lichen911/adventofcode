package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("day1_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calibration_sum int = 0
	var first_digit rune
	var second_digit rune

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("line value: %v\n", line)

		for i := 0; i < len(line); i++ {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				first_digit = r
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				second_digit = r
				break
			}
		}

		final_num, err := strconv.Atoi(string(first_digit) + string(second_digit))
		checkError(err)

		calibration_sum += final_num
	}

	fmt.Println("Final calibration value: ", calibration_sum)
}
