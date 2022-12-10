package main

import (
	_ "embed"
	"fmt"
	"strconv"
	s "strings"
)

//go:embed input.txt
var input string

type RopeNode struct {
	x int
	y int
}

func (r RopeNode) String() string {
	return fmt.Sprintf("(%d,%d)", r.x, r.y)
}

func tailMustMove(head, tail RopeNode) bool {
	return abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (r *RopeNode) determineMove(prev RopeNode) {

	// x, y := r.x, r.y
	if !tailMustMove(prev, *r) {
		return
	}
	if prev.x-r.x > 0 {
		r.x++
	} else if prev.x-r.x < 0 {
		r.x--
	}

	if prev.y-r.y > 0 {
		r.y++
	} else if prev.y-r.y < 0 {
		r.y--
	}
}

func part1() {
	lines := s.Split(input, "\n")

	// head, tail := RopeNode{x: 0, y: 0}, RopeNode{x: 0, y: 0}

	rope := make([]RopeNode, 10)
	var head *RopeNode = &rope[0]
	var tail *RopeNode = &rope[9]

	visited := map[RopeNode]bool{}
	// visited[*head] = true
	for _, line := range lines {
		direction := line[0]
		distance, _ := strconv.Atoi(line[2:])

		for i := 0; i < distance; i++ {
			switch direction {
			case 'U':
				head.y++
			case 'L':
				head.x--
			case 'D':
				head.y--
			case 'R':
				head.x++
			default:
				panic("SOMETHING WENT WRONG")
			}

			for j := 1; j < len(rope); j++ {
				rope[j].determineMove(rope[j-1])
			}
			// if tailMustMove(head, tail) {
			// 	tail.determineMove(head)
			// 	visited[tail] = true
			// }
			visited[*tail] = true
			fmt.Println(rope)
		}

	}
	fmt.Println(len(visited))
	// for key := range visited {
	// 	fmt.Println(key)
	// }
	// fmt.Println(visited)
	// arr := make([]RopeNode, 10)
	// fmt.Println(arr)
}

func main() {
	part1()
}
