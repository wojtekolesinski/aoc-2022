package main

import (
	_ "embed"
	"fmt"
	s "strings"
)

//go:embed input.txt
var input string

type Register struct {
	cycle int
	value int
}

func (r *Register) noop() {
	r.cycle++
}

func (r Register) checkCycle() int {
	if r.cycle == 20 || (r.cycle - 20) % 40 == 0 {
		fmt.Println("cycle: ", r.cycle, " value: ", r.value,  " -> ", r.value * r.cycle)
		return r.value * r.cycle
	}
	return 0
}

func part1() {
	lines := s.Split(input, "\n")

	register := Register{value: 1, cycle: 1}
	sum := 0
	for _, line := range lines {
		var command string
		var value int
		fmt.Sscanf(line, "%s %d", &command, &value)

		switch command {
		case "noop":
			register.noop()
			sum += register.checkCycle()
		case "addx":
			register.noop()
			sum += register.checkCycle()
			register.noop()
			register.value += value
			sum += register.checkCycle()
		default:
			panic("SOMETHING WENT WRONG")
		}
	}

	fmt.Println(sum)
}

func main() {
	part1()
}