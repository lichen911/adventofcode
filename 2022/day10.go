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

type Instruction struct {
	op  string
	arg int
}

type Processor struct {
	x             int
	cycle         int
	instructions  []Instruction
	monitorCycles map[int]int
	crt 		  [6][40]bool
}

func (p *Processor) incCycle() {
	fmt.Println("Cycle:", p.cycle, "Register:", p.x)

	_, ok := p.monitorCycles[p.cycle]

	if ok {
		p.monitorCycles[p.cycle] = p.cycle * p.x
	}

	p.setPixel()
	p.cycle++
	
}

func (p *Processor) setPixel() {
	row := (p.cycle - 1) / 40
	col := (p.cycle - 1) % 40

	if p.x-1 <= col && col <= p.x+1 {
		p.crt[row][col] = true
	}
}

func (p *Processor) printCrt() {
	for row := 0; row < len(p.crt); row++ {
		for col := 0; col < len(p.crt[0]); col++ {
			if p.crt[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			
		}
		fmt.Println()
	}
}

func (p *Processor) run() {
	for i := 0; i < len(p.instructions); i++ {
		instruction := p.instructions[i]
		// fmt.Println("Cpu pre-op: Cycle:", p.cycle, "Reg:", p.x, "Op:", instruction.op, "Arg:", instruction.arg)
		switch instruction.op {
		case "addx":
			p.addx(instruction)
		case "noop":
			p.noop(instruction)
		}
	}
}

func (p *Processor) addx(instruction Instruction) {
	// addx takes two cycle to complete
	p.incCycle()
	p.incCycle()
	p.x += instruction.arg
}

func (p *Processor) noop(instruction Instruction) {
	// noop takes one cycle to complete
	p.incCycle()
}

func (p Processor) sumSignalStrength() int {
	total := 0
	for _, signalStrength := range p.monitorCycles {
		total += signalStrength
	}
	return total
}

func (p *Processor) addInstruction(text string) {
	textArgs := strings.Split(text, " ")
	op := textArgs[0]

	var arg int
	var err error
	if len(textArgs) == 2 {
		arg, err = strconv.Atoi(textArgs[1])
		checkError(err)
	}
	p.instructions = append(p.instructions, Instruction{op: op, arg: arg})
}

func NewProcessor(monitorCycles []int) *Processor {
	monitorCyclesLookup := map[int]int{}
	for i := 0; i < len(monitorCycles); i++ {
		monitorCyclesLookup[monitorCycles[i]] = 0
	}

	return &Processor{
		x:             1,
		cycle:         1,
		instructions:  []Instruction{},
		monitorCycles: monitorCyclesLookup,
		crt: [6][40]bool{},
	}
}

func main() {
	file, err := os.Open("day10_input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpu := NewProcessor([]int{20, 60, 100, 140, 180, 220})

	for scanner.Scan() {
		text := scanner.Text()
		cpu.addInstruction(text)
	}

	cpu.run()

	fmt.Println("Signal strength:", cpu.sumSignalStrength())
	cpu.printCrt()
}
