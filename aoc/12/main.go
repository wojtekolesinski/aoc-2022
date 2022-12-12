package main

import (
	_ "embed"
	"fmt"
	s "strings"

	"github.com/wojtekolesinski/aoc-2022/util"
)

//go:embed input.txt
var input string

type Point struct {
	x, y, height int
}

func newPoint(x, y, height int) *Point {
	if height == 'S' {
		height = 'a' - 1
	}
	if height == 'E' {
		height = 'z' + 1
	}

	return &Point{
		x:      x,
		y:      y,
		height: height,
	}
}

type Node struct {
	*Point
	previous *Node
}

func newNode(point Point, prev *Node) *Node {
	return &Node{
		Point:    &point,
		previous: prev,
	}
}

func (n Node) String() string {
	return fmt.Sprintf("[x=%d y=%d height=%c]", n.x, n.y, n.height)
}

func (n *Node) getNeighbours(grid [][]rune) []Node {
	neighbours := []Node{}
	if n.x != 0 {
		new := *newNode(*newPoint(n.x-1, n.y, int(grid[n.y][n.x-1])), n)
		if new.height-n.height <= 1 {
			neighbours = append(neighbours, new)
		}
	}

	if n.x != len(grid[0])-1 {
		new := *newNode(*newPoint(n.x+1, n.y, int(grid[n.y][n.x+1])), n)
		if new.height-n.height <= 1 {
			neighbours = append(neighbours, new)
		}
	}

	if n.y != 0 {
		new := *newNode(*newPoint(n.x, n.y-1, int(grid[n.y-1][n.x])), n)
		if new.height-n.height <= 1 {
			neighbours = append(neighbours, new)
		}
	}

	if n.y != len(grid)-1 {
		new := *newNode(*newPoint(n.x, n.y+1, int(grid[n.y+1][n.x])), n)
		if new.height-n.height <= 1 {
			neighbours = append(neighbours, new)
		}
	}

	return neighbours
}

func parseInput() [][]rune {
	lines := s.Split(input, "\n")
	terrainMap := make([][]rune, 0)

	for _, line := range lines {
		terrainMap = append(terrainMap, []rune(line))
	}

	return terrainMap
}

func search(grid [][]rune, x, y int) Node {
	q := util.Queue[Node]{}
	q.Add(*newNode(*newPoint(x, y, int(grid[y][x])), nil))

	visited := map[Point]bool{}

	for !q.IsEmpty() {
		curr := q.Pop()
		if _, exists := visited[*curr.Point]; exists {
			continue
		}

		if curr.height == 'z'+1 {
			return curr
		}

		q.Add(curr.getNeighbours(grid)...)
		visited[*curr.Point] = true
	}

	return Node{}
}

func part1() {
	grid := parseInput()

	result := search(grid, 0, 0)

	curr := result

	steps := 0
	for curr.previous != nil {
		steps++
		curr = *curr.previous

	}
	fmt.Println(steps)
}

func part2() {
	grid := parseInput()

	startingPoints := [][]int{}

	for y, row := range grid {
		for x, el := range row {
			if el == 'a' || el == 'S' {
				startingPoints = append(startingPoints, []int{x, y})
			}
		}
	}

	results := []Node{}
	for _, point := range startingPoints {
		results = append(results, search(grid, point[0], point[1]))
	}

	min_steps := 99999
	for _, result := range results {
		if result.previous == nil {
			continue
		}

		curr := result

		steps := 0
		for curr.previous != nil {
			steps++
			curr = *curr.previous

		}
		if steps < min_steps {
			min_steps = steps
		}
	}
	fmt.Println(min_steps)
}

func main() {
	part1()
	part2()
}
