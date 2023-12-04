package main

import (
	"bufio"
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

var textDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	file, err := os.Open("day1_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calibration_sum int = 0
	var firstDigit rune
	var secondDigit rune

	for scanner.Scan() {
		line := scanner.Text()

	outerloop1:
		for i := 0; i < len(line); i++ {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				firstDigit = r
				break
			} else {
				for k, v := range textDigits {
					if strings.HasPrefix(line[i:], k) {
						firstDigit = rune(v[0])
						break outerloop1
					}
				}
			}
		}

	outerloop2:
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				secondDigit = r
				break
			} else {
				for k, v := range textDigits {
					if strings.HasSuffix(line[:i+1], k) {
						secondDigit = rune(v[0])
						break outerloop2
					}
				}
			}
		}

		final_num, err := strconv.Atoi(string(firstDigit) + string(secondDigit))
		checkError(err)

		calibration_sum += final_num
	}

	fmt.Println("Final calibration value: ", calibration_sum)
}
