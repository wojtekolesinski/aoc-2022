package main

import (
	_ "embed"
	"fmt"
	s "strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	lines := s.Split(input, "\n")

	results := []rune{}

	for _, line := range lines {
		compartment1, compartment2 := line[:len(line)/2], line[len(line)/2:]
		set1, set2 := make(map[rune]bool), make(map[rune]bool)
		for _, letter := range []rune(compartment1) {
			set1[letter] = true
		}

		for _, letter := range []rune(compartment2) {
			set2[letter] = true
		}

		for letter := range set1 {
			if _, ok := set2[letter]; ok {
				// fmt.Printf("For compartments %s and %s item %s repeats\n", compartment1, compartment2, string(letter))
				results = append(results, letter)
			}
		}

	}

	sum := 0
	for _, letter := range results {
		sum += int(getPriority(letter))
	}

	return sum
}

func part2() int {
	lines := s.Split(input, "\n")

	results := []rune{}

	for i := 0; i < len(lines); i += 3 {
		elf1, elf2, elf3 := lines[i], lines[i+1], lines[i+2]

		set1, set2, set3 := make(map[rune]bool), make(map[rune]bool), make(map[rune]bool)
		for _, letter := range []rune(elf1) {
			set1[letter] = true
		}
		for _, letter := range []rune(elf2) {
			set2[letter] = true
		}
		for _, letter := range []rune(elf3) {
			set3[letter] = true
		}
		for letter := range set1 {
			if _, ok := set2[letter]; ok {
				if _, ok2 := set3[letter]; ok2 {
					results = append(results, letter)
				}
			}
		}
	}

	sum := 0
	for _, letter := range results {
		sum += int(getPriority(letter))
	}
	return sum
}

func getPriority(item rune) rune {
	if item >= 'a' && item <= 'z' {
		return item - 'a' + 1
	} else {
		return item - 'A' + 1 + 26
	}
}
