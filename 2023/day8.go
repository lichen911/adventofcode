package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	left, right string
}

func getNextIns(instructions string, insIndex int) (string, int) {
	newIndex := insIndex + 1
	if newIndex >= len(instructions) || insIndex == -1 {
		return string(instructions[0]), 0
	} else {
		return string(instructions[newIndex]), newIndex
	}
}

func TraverseNetwork(nodeMap map[string]Node, instructions string) int {
	stepCount := 0
	currentNode := "AAA"
	currentIns, currentInsIdx := getNextIns(instructions, -1)

	for currentNode != "ZZZ" {
		if currentIns == "L" {
			currentNode = nodeMap[currentNode].left
		} else {
			currentNode = nodeMap[currentNode].right
		}

		currentIns, currentInsIdx = getNextIns(instructions, currentInsIdx)
		stepCount++
	}
	return stepCount
}

func main() {
	file, err := os.Open("day8_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodeMap := make(map[string]Node)
	readNodes := false
	var instructions string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readNodes = true
			continue
		}

		if !readNodes {
			instructions = line
			continue
		}

		if readNodes {
			nodeName := strings.Fields(line)[0]
			nodeLeft := strings.Fields(line)[2][1:4]
			nodeRight := strings.Fields(line)[3][:3]
			nodeMap[nodeName] = Node{left: nodeLeft, right: nodeRight}
		}
	}

	// fmt.Println("ins:", instructions)
	// fmt.Println("nodeMap:", nodeMap)

	stepCount := TraverseNetwork(nodeMap, instructions)
	fmt.Println("Step count:", stepCount)
}
