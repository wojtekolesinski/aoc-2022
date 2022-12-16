package main

import (
	_ "embed"
	"fmt"
	s "strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part2())
}

func part1() int {
	moves := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	win := 6
	draw := 3
	lose := 0

	debug := map[string] string {
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock (1)",
		"Y": "paper (2)",
		"Z": "scissors(3)",
	}

	scores := map[int] []int {
		1: []int{0, draw, lose, win}, // rock
		2: []int{0, win, draw, lose}, // paper
		3: []int{0, lose, win, draw}, // scissors
	}

	lines := s.Split(input, "\n")
	score := 0
	for _, line := range lines {
		values := s.Split(line, " ")
		opponent := values[0]
		player := values[1]
		this := getScoreForPart1(moves[opponent], moves[player], scores)
		fmt.Println(debug[opponent], debug[player], this)
		// fmt.Println(this)
		score += this

	}
	return score
}

func getScoreForPart1(opponent, player int, scores map[int][]int) int {
	return player + scores[player][opponent]
}

func part2() int {
	rock, paper, scissors := 1, 2, 3
	win, lose, draw := 6, 0, 3

	moves := map[string]int{
		"A X": lose + scissors,
		"A Y": draw + rock,
		"A Z": win + paper,
		"B X": lose + rock,
		"B Y": draw + paper,
		"B Z": win + scissors,
		"C X": lose + paper,
		"C Y": draw + scissors,
		"C Z": win + rock,
	}

	// points := []int {0, 3, 6}

	// // correct choices to lose, draw or win
	// correctChoices := map[int] []int {
	// 	rock: {paper, rock, scissors},
	// 	paper: {scissors, paper, rock},
	// 	scissors: {rock, scissors, paper},
	// }

	lines := s.Split(input, "\n")
	score := 0
	for _, line := range lines {
		
		
		score += moves[line]
	}

	return score
}
