package main

import (
	_ "embed"
	"fmt"
	s "strings"
	"strconv"
	"github.com/wojtekolesinski/aoc-2022/util"
)

//go:embed input.txt
var input string

func main() {
	part2()

}

func part1() {
	lines := s.Split(input, "\n")

	piles := make([]util.Stack, 9)

	fmt.Println(piles)
	firstPart := true
	for _, line := range lines {
		if firstPart {
			if s.TrimSpace(line) == "" {
				firstPart = false
				
				for number, pile := range piles {
					fmt.Print(number+1, " ")
					for _, c := range pile {
						fmt.Print(string(c))
					}
					fmt.Print(" ")
				}
				fmt.Println()
				continue
			}

			if line [1] == '1' {
				continue
			}

			for i, j := 1, 0; i < len(line); i+=4 {
				if (line[i] != ' ') {
					piles[j] = append(util.Stack{line[i]}, piles[j]...)
				}
				j++
			}
			fmt.Println()

		} else {
			words := s.Split(line, " ")
			howMany, err := strconv.Atoi(words[1])
			util.CheckError(err)

			from, err := strconv.Atoi(words[3])
			util.CheckError(err)

			to, err := strconv.Atoi(words[5])
			util.CheckError(err)

			fmt.Println(line)

			for i := 0; i < howMany; i++ {
				if el, ok := piles[from-1].Pop(); ok {
					fmt.Println("Pushing from ", from, " to ", to)
					piles[to-1].Push(el)
				}

			}

			for number, pile := range piles {
				fmt.Print(number+1, " ")
				for _, c := range pile {
					fmt.Print(string(c))
				}
				fmt.Print(" ")
			}
			fmt.Println()

			// fmt.Println(line)
		}
	}

	for _, pile := range piles {
		fmt.Print(string(pile[len(pile)-1]))
		// fmt.Print(number+1, " ")
		// for _, crate := range pile {
		// 	fmt.Print( string(crate), " ")
		// }
		// fmt.Println()
	}
	// fmt.Println(piles)
}

func part2() {
	lines := s.Split(input, "\n")

	piles := make([]util.Stack, 9)

	fmt.Println(piles)
	firstPart := true
	for _, line := range lines {
		if firstPart {
			if s.TrimSpace(line) == "" {
				firstPart = false
				
				for number, pile := range piles {
					fmt.Print(number+1, " ")
					for _, c := range pile {
						fmt.Print(string(c))
					}
					fmt.Print(" ")
				}
				fmt.Println()
				continue
			}

			if line [1] == '1' {
				continue
			}

			for i, j := 1, 0; i < len(line); i+=4 {
				if (line[i] != ' ') {
					piles[j] = append(util.Stack{line[i]}, piles[j]...)
				}
				j++
			}
			fmt.Println()

		} else {
			words := s.Split(line, " ")
			howMany, err := strconv.Atoi(words[1])
			util.CheckError(err)

			from, err := strconv.Atoi(words[3])
			util.CheckError(err)

			to, err := strconv.Atoi(words[5])
			util.CheckError(err)

			fmt.Println(line)

			if popped, ok := piles[from-1].PopN(howMany); ok {
				// fmt.Println("Pushing: ", string(popped[:]))
				piles[to-1] = append(piles[to-1], popped...)
			}

			// for i := 0; i < howMany; i++ {

			// 	if el, ok := piles[from-1].Pop(); ok {
			// 		fmt.Println("Pushing from ", from, " to ", to)
			// 		piles[to-1].Push(el)
			// 	}

			// }

			for number, pile := range piles {
				fmt.Print(number+1, " ")
				for _, c := range pile {
					fmt.Print(string(c))
				}
				fmt.Print(" ")
			}
			fmt.Println()

			// fmt.Println(line)
		}
	}

	for _, pile := range piles {
		fmt.Print(string(pile[len(pile)-1]))
		// fmt.Print(number+1, " ")
		// for _, crate := range pile {
		// 	fmt.Print( string(crate), " ")
		// }
		// fmt.Println()
	}
	fmt.Println(piles)
}
