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

type MapRange struct {
	sourceStart int
	destStart   int
	length      int
}

func loadSeedMap(fileName string) ([]int, [][]MapRange) {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapLoadIndex := -1
	seeds := make([]int, 0)

	// Initialize map
	seedMap := make([][]MapRange, 7)
	for i := range seedMap {
		seedMap[i] = make([]MapRange, 0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			for _, seed := range strings.Fields(line)[1:] {
				seedInt, err := strconv.Atoi(seed)
				checkError(err)
				seeds = append(seeds, seedInt)
			}
			continue
		} else if strings.HasSuffix(line, "map:") {
			mapLoadIndex++
			continue
		}

		destStart, err := strconv.Atoi(strings.Fields(line)[0])
		checkError(err)
		sourceStart, err := strconv.Atoi(strings.Fields(line)[1])
		checkError(err)
		length, err := strconv.Atoi(strings.Fields(line)[2])
		checkError(err)

		newMapRange := MapRange{
			sourceStart: sourceStart,
			destStart:   destStart,
			length:      length,
		}
		seedMap[mapLoadIndex] = append(seedMap[mapLoadIndex], newMapRange)
	}

	return seeds, seedMap
}

func getMapSectionOutput(sourceNum int, mapSection []MapRange) int {
	for _, mapRow := range mapSection {
		if sourceNum >= mapRow.sourceStart &&
			sourceNum < mapRow.sourceStart+mapRow.length {
			if mapRow.destStart > mapRow.sourceStart {
				return sourceNum + (mapRow.destStart - mapRow.sourceStart)
			} else {
				return sourceNum - (mapRow.sourceStart - mapRow.destStart)
			}
		}
	}
	return sourceNum
}

func getLocNum(seed int, seedMap [][]MapRange) int {
	sourceNum, destNum := seed, -1
	for _, mapSection := range seedMap {
		destNum = getMapSectionOutput(sourceNum, mapSection)
		sourceNum = destNum
	}

	return destNum
}

func findLowestLocNum(seeds []int, seedMap [][]MapRange) int {
	lowestLocNum := getLocNum(seeds[0], seedMap)

	for i := 0; i < len(seeds)-2; i += 2 {
		seedRangeStart := seeds[i]
		seedRangeLength := seeds[i+1]

		for seed := seedRangeStart; seed < seedRangeStart+seedRangeLength; seed++ {
			locNum := getLocNum(seed, seedMap)
			if locNum < lowestLocNum {
				lowestLocNum = locNum
			}
		}
	}

	return lowestLocNum
}

func main() {
	seeds, seedMap := loadSeedMap("day5_input.txt")
	fmt.Println("seed:", seeds)
	fmt.Println("seed map:", seedMap)

	lowestLocNum := findLowestLocNum(seeds, seedMap)
	fmt.Println("Lowest location number:", lowestLocNum)
}
