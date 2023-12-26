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

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmm(args ...int) int {
	// Recursively iterate through pairs of arguments
	// i.e. lcm(args[0], lcm(args[1], lcm(args[2], args[3])))

	if len(args) == 2 {
		return lcm(args[0], args[1])
	} else {
		return lcm(args[0], lcmm(args[1:]...))
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

func getStartNodes(nodeMap map[string]Node) []string {
	startNodes := make([]string, 0)
	for node := range nodeMap {
		if string(node[2]) == "A" {
			startNodes = append(startNodes, node)
		}
	}
	return startNodes
}

func TraverseNetwork(nodeMap map[string]Node, instructions string) int {
	stepCount := 0
	stepCounts := make([]int, 0)
	startNodes := getStartNodes(nodeMap)

	for _, currentNode := range startNodes {
		currentIns, currentInsIdx := getNextIns(instructions, -1)

		for !strings.HasSuffix(currentNode, "Z") {
			if currentIns == "L" {
				currentNode = nodeMap[currentNode].left
			} else {
				currentNode = nodeMap[currentNode].right
			}

			currentIns, currentInsIdx = getNextIns(instructions, currentInsIdx)
			stepCount++
		}
		stepCounts = append(stepCounts, stepCount)
		stepCount = 0
	}

	return lcmm(stepCounts...)
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

	stepCount := TraverseNetwork(nodeMap, instructions)
	fmt.Println("Step count:", stepCount)
}
