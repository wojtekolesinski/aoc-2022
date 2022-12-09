package main

import (
	_ "embed"
	"fmt"
	s "strings"
	"strconv"
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

func (r *RopeNode) determineTailMove(head RopeNode) {
	// x, y := r.x, r.y

	if head.x-r.x > 0 {
		r.x++
	} else if head.x-r.x < 0 {
		r.x--
	}

	if head.y-r.y > 0 {
		r.y++
	} else if head.y-r.y < 0 {
		r.y--
	}

	// return x, y

	// if head.x == tail.x {
	// 	return tail.x, (head.y + tail.y) / 2
	// }

	// if head.y == tail.y {
	// 	return (head.x + tail.x) / 2, tail.y
	// }

	// dx, dy := head.x-tail.x, head.y-tail.y
	// dx = dx / int(math.Abs(float64(dx)))
	// dy = dy / int(math.Abs(float64(dy)))

	// return tail.x + dx, tail.y + dy
}

func part1() {
	lines := s.Split(input, "\n")

	head, tail := RopeNode{x: 0, y: 0}, RopeNode{x: 0, y: 0}

	tailLocations := map[RopeNode]bool{}
	tailLocations[tail] = true
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

			if tailMustMove(head, tail) {
				tail.determineTailMove(head)
				tailLocations[tail] = true
			}
		}

	}
	fmt.Println(len(tailLocations))
}

func main() {
	part1()
}
