package main

import (
	_ "embed"
	"fmt"
	s "strings"
	// "strconv"

	// "github.com/wojtekolesinski/aoc-2022/util"
)

//go:embed input.txt
var input string

func isVisible(grid [][]int, row, col int) bool {
	el := grid[row][col]
	// visible from top
	visible := true
	for i := 0; i < row; i++ {
		if grid[i][col] >= el {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// visible from bottom
	visible = true
	for i := len(grid)-1; i > row; i-- {
		if grid[i][col] >= el {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// visible from right
	visible = true
	for i := len(grid)-1; i > col; i-- {
		if grid[row][i] >= el {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	// visible from left
	visible = true
	for i := 0; i < col; i++ {
		if grid[row][i] >= el {
			visible = false
			break
		}
	}
	return visible
}

func getScenicScore(grid [][]int, row, col int) int {
	el := grid[row][col]
	// score from top
	top := 0
	for i := row-1; i >= 0; i-- {
		top++
		if grid[i][col] >= el {
			break
		}
	}

	// score from bottom
	bottom := 0
	for i := row+1; i < len(grid); i++ {
		bottom++
		if grid[i][col] >= el {
			break
		}
	}

	// score from right
	right := 0
	for i := col+1; i < len(grid[0]); i++ {
		right++
		if grid[row][i] >= el {
			break
		}
	}

	// score from left
	left := 0
	for i := col-1; i >= 0; i-- {
		left++
		if grid[row][i] >= el {
			break
		}
	}
	return top * bottom * left * right
}

func part1() {
	lines := s.Split(input, "\n")

	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}
	for row, line := range lines {
		for column, tree := range line {
			grid[row][column] = int(tree - '0')
		}
	}

	sum := 4 * len(grid) - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			if isVisible(grid, i, j) {
				sum++
			}
		}
	}

	fmt.Println(grid)
	fmt.Println(sum)
}

func part2() {
	lines := s.Split(input, "\n")

	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}
	for row, line := range lines {
		for column, tree := range line {
			grid[row][column] = int(tree - '0')
		}
	}

	maxScore := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			score := getScenicScore(grid, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(grid)
	fmt.Println(maxScore)
}

func main() {
	// part1()
	part2()
}