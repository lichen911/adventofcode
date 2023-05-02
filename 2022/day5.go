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

type CrateStack struct {
	crates []string
}

func (cs *CrateStack) isEmpty() bool {
	return len(cs.crates) == 0
}

func (cs *CrateStack) push(crate string) {
	cs.crates = append(cs.crates, crate)
}

func (cs *CrateStack) pop() string {
	if cs.isEmpty() {
		return ""
	}
	crate := cs.crates[len(cs.crates)-1]
	cs.crates = cs.crates[:len(cs.crates)-1]
	return crate
}

func (cs *CrateStack) peek() string {
	if cs.isEmpty() {
		return ""
	}
	return cs.crates[len(cs.crates)-1]
}

type CrateGroup struct {
	totalStacks int
	crates      []CrateStack
}

func (cs *CrateGroup) create(crate_input string) {
	crate_lines := strings.Split(crate_input, "\n")
	stack_numbers := strings.Split(crate_lines[len(crate_lines)-2], " ")

	var err error
	cs.totalStacks, err = strconv.Atoi(stack_numbers[len(stack_numbers)-2])
	checkError(err)

	cs.crates = make([]CrateStack, cs.totalStacks, cs.totalStacks)

	for i := len(crate_lines) - 3; i >= 0; i-- {
		current_line := crate_lines[i]
		for j := 0; j < len(current_line); j += 4 {
			current_crate := current_line[j : j+3]
			if current_crate != "   " {
				current_crate = strings.ReplaceAll(current_crate, "[", "")
				current_crate = strings.ReplaceAll(current_crate, "]", "")
				cs.crates[j/4].push(current_crate)
			}
		}
	}
}


func (cs *CrateGroup) getTopCrates() string {
	topCrates := ""
	for i := 0; i < cs.totalStacks; i++ {
		topCrates += cs.crates[i].peek()
	}

	return topCrates
}

type CrateMover9000 struct {
	CrateGroup
}

func (cm *CrateMover9000) move(quantity int, source int, destination int) {
	for i := 0; i < quantity; i++ {
		crate := cm.crates[source-1].pop()
		cm.crates[destination-1].push(crate)
	}
}

type CrateMover9001 struct {
	CrateGroup
}

func (cm *CrateMover9001) move(quantity int, source int, destination int) {
	move_crates := make([]string, 0, quantity)
	for i := 0; i < quantity; i++ {
		crate := cm.crates[source-1].pop()
		move_crates = append(move_crates, crate)
	}

	for i := len(move_crates) - 1; i >= 0; i-- {
		cm.crates[destination-1].push(move_crates[i])
	}
}

func main() {
	file, err := os.Open("day5_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var crate_desc string
	scanned_crate_desc := false
	crateMover9000 := CrateMover9000{}
	crateMover9001 := CrateMover9001{}

	for scanner.Scan() {
		current_line := scanner.Text()
		if !scanned_crate_desc {
			// Read crate description
			if current_line == "" {
				scanned_crate_desc = true
				crateMover9000.create(crate_desc)
				crateMover9001.create(crate_desc)
			} else {
				crate_desc += current_line + "\n"
			}
		} else {
			// Read movement steps
			quantity, err := strconv.Atoi(strings.Split(current_line, " ")[1])
			checkError(err)

			source, err := strconv.Atoi(strings.Split(current_line, " ")[3])
			checkError(err)

			destination, err := strconv.Atoi(strings.Split(current_line, " ")[5])
			checkError(err)

			crateMover9000.move(quantity, source, destination)
			crateMover9001.move(quantity, source, destination)
		}
	}

	fmt.Println("CrateMover 9000:", crateMover9000.getTopCrates())
	fmt.Println("CrateMover 9001:", crateMover9001.getTopCrates())
}
