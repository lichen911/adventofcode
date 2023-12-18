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

func getRaceMultiple(raceTimes, raceDistances []int) int {
	var distance int
	raceMultiple := 1

	for i, raceTime := range raceTimes {
		winCount := 0
		raceDistance := raceDistances[i]
		// fmt.Println("race time:", raceTime, "race distance:", raceDistance)

		for holdTime := 1; holdTime < raceTime; holdTime++ {
			movingTime := raceTime - holdTime
			distance = holdTime * movingTime

			if distance > raceDistance {
				winCount++
			}
			// fmt.Println(
			// 	"hold time:", holdTime,
			// 	"moving time:", movingTime,
			// 	"distance:", distance,
			// 	"win count:", winCount,
			// )

		}

		raceMultiple *= winCount
	}

	return raceMultiple
}

func main() {
	file, err := os.Open("day6_input_sample.txt")
	checkError(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	raceTimes := make([]int, 0)
	raceDistances := make([]int, 0)

	line, err := reader.ReadString('\n')
	checkError(err)
	for _, raceTime := range strings.Fields(line)[1:] {
		raceTimeInt, err := strconv.Atoi(raceTime)
		checkError(err)
		raceTimes = append(raceTimes, raceTimeInt)
	}

	line, err = reader.ReadString('\n')
	checkError(err)
	for _, raceDistance := range strings.Fields(line)[1:] {
		raceDistanceInt, err := strconv.Atoi(raceDistance)
		checkError(err)
		raceDistances = append(raceDistances, raceDistanceInt)
	}

	fmt.Println("raceTimes:", raceTimes)
	fmt.Println("raceDistances:", raceDistances)

	raceMultiple := getRaceMultiple(raceTimes, raceDistances)
	fmt.Println("raceMultiple:", raceMultiple)

}
