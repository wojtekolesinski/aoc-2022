package main

import (
	_ "embed"
	"fmt"
	// s "strings"
)

//go:embed input.txt
var input string

func main() {
	part2()
}

func part1() {
	for i := 3; i < len(input); i++ {
		codeMap := map[byte]bool{}
		for j := 0; j < 4; j++ {
			codeMap[input[i-j]] = true
		}
		// fmt.Println(codeMap)
		x := len(codeMap) == 4
		fmt.Print(x, i+1, " ")
		for code := range codeMap {
			fmt.Print(string(code))
		}
		fmt.Println()

		if x {
			return
		}
	}
}

func part2() {
	for i := 13; i < len(input); i++ {
		codeMap := map[byte]bool{}
		for j := 0; j < 14; j++ {
			codeMap[input[i-j]] = true
		}
		// fmt.Println(codeMap)
		x := len(codeMap) == 14
		fmt.Print(x, i+1, " ")
		for code := range codeMap {
			fmt.Print(string(code))
		}
		fmt.Println()

		if x {
			return
		}
	}
}

