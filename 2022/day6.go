package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func shiftStreamWindow(streamWindow []string, receivedData string) []string {
	streamWindow = append(streamWindow, receivedData)
	streamWindow = streamWindow[1:]
	return streamWindow
}

func checkDuplicates(streamWindow []string) bool {
	seenChar := make(map[string]bool)
	
	for _, char := range streamWindow {
		if char == "" {
			return true
		}
		if _, seen := seenChar[char]; seen {
			return true
		}
		seenChar[char] = true
	}
	return false
}

func main() {
	markerCount, err := strconv.Atoi(os.Args[1])
	checkError(err)

	file, err := os.Open("day6_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	packetMarker := make([]string, markerCount, markerCount)
	count := 1
	for scanner.Scan() {
		packetMarker = shiftStreamWindow(packetMarker, scanner.Text())
		if !checkDuplicates(packetMarker) {
			break
		}
		count++
	}

	fmt.Println("Packet marker:", count)
}
