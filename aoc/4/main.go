package main

import (
	_ "embed"
	"fmt"
	"strconv"
	s "strings"

	"github.com/wojtekolesinski/aoc-2022/util"
)

//go:embed input.txt
var input string

type cleaningRange struct {
	start int
	end   int
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	lines := s.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		r1, r2 := parseInput(line)
		overlap := checkForFullOverlap(r1, r2)
		if overlap {
			sum++
		}
		fmt.Println(line, r1, r2, checkForFullOverlap(r1, r2))

	}

	return sum
}

func part2() int {
	lines := s.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		r1, r2 := parseInput(line)
		overlap := checkForAnyOverlap(r1, r2)
		if overlap {
			sum++
		}
		fmt.Println(line, r1, r2, checkForAnyOverlap(r1, r2))

	}

	return sum
}

func parseInput(line string) (cleaningRange, cleaningRange) {
	elfs := s.Split(line, ",")
	elf1 := s.Split(elfs[0], "-")
	elf2 := s.Split(elfs[1], "-")

	start1, err := strconv.Atoi(elf1[0])
	util.CheckError(err)

	end1, err := strconv.Atoi(elf1[1])
	util.CheckError(err)

	start2, err := strconv.Atoi(elf2[0])
	util.CheckError(err)

	end2, err := strconv.Atoi(elf2[1])
	util.CheckError(err)

	return cleaningRange{start: start1, end: end1}, cleaningRange{start: start2, end: end2}
}

func checkForFullOverlap(r1, r2 cleaningRange) bool {
	if r1.start >= r2.start && r1.end <= r2.end ||
		r2.start >= r1.start && r2.end <= r1.end {
		return true
	}

	return false
}

func checkForAnyOverlap(r1, r2 cleaningRange) bool {
	if r2.end >= r1.start && r2.start <= r1.start ||
		r2.start <= r1.end && r2.start >= r1.start ||
		checkForFullOverlap(r1, r2) {
		return true
	}

	return false
}
