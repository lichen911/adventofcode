package main

import (
	"bufio"
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

func ConvertStrings(input []string) []int {
	output := make([]int, 0)
	for _, item := range input {
		itemInt, err := strconv.Atoi(item)
		checkError(err)
		output = append(output, itemInt)
	}
	return output
}

func getHistDiffs(input []int) []int {
	output := make([]int, 0)
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		output = append(output, diff)
	}
	return output
}

func checkAllZero(input []int) bool {
	for _, val := range input {
		if val != 0 {
			return false
		}
	}
	return true
}

func addNextValDiff(input *[][]int) int {
	nextValSum := 0
	lastRowIdx := len(*input) - 1
	(*input)[lastRowIdx] = append((*input)[lastRowIdx], 0)

	for i := lastRowIdx - 1; i >= 0; i-- {
		lastVal := (*input)[i][len((*input)[i])-1]
		nextVal := lastVal + (*input)[i+1][len((*input)[i+1])-1]
		(*input)[i] = append((*input)[i], nextVal)

		if i == 0 {
			nextValSum += nextVal
		}
	}
	return nextValSum
}

func addPrevValDiff(input *[][]int) int {
	prevValSum := 0
	lastRowIdx := len(*input) - 1
	(*input)[lastRowIdx] = append([]int{0}, (*input)[lastRowIdx]...)

	for i := lastRowIdx - 1; i >= 0; i-- {
		firstVal := (*input)[i][0]
		prevVal := firstVal - (*input)[i+1][0]
		(*input)[i] = append([]int{prevVal}, (*input)[i]...)

		if i == 0 {
			prevValSum += prevVal
		}
	}
	return prevValSum
}

func getHistVals(histSeq []int) (int, int) {
	histValDiffs := make([][]int, 0)
	histValDiffs = append(histValDiffs, histSeq)

	var curHistdiffs []int

	curHistdiffs = getHistDiffs(histSeq)
	histValDiffs = append(histValDiffs, curHistdiffs)

	for !checkAllZero(curHistdiffs) {
		curHistdiffs = getHistDiffs(curHistdiffs)
		histValDiffs = append(histValDiffs, curHistdiffs)
	}

	histValSum := addNextValDiff(&histValDiffs)
	prevHistValSum := addPrevValDiff(&histValDiffs)
	// fmt.Println("histValDiffs:", histValDiffs, "histValSum:", histValSum)
	return histValSum, prevHistValSum
}

func main() {
	file, err := os.Open("day9_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	histValAccum, prevHistValAccum := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		histSeqStr := strings.Fields(line)
		histSeq := ConvertStrings(histSeqStr)
		histValSum, prevHistValSum := getHistVals(histSeq)
		histValAccum += histValSum
		prevHistValAccum += prevHistValSum
	}

	fmt.Println("History value sum:", histValAccum)
	fmt.Println("Previous history value sum:", prevHistValAccum)
}
