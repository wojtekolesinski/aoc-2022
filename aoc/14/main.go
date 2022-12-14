package main

import (
	_ "embed"
	"fmt"
	s "strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func parsePoint(s string) *Point {
	var x, y int
	fmt.Sscanf(s, "%d,%d", &x, &y)
	return &Point{x: x, y: y}
}

func addRocks(cave [][]string, paths [][]Point) [][]string {
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			start, end := path[i], path[i+1]

			if start.x == end.x {
				var starty, endy int
				if start.y < end.y {
					starty, endy = start.y, end.y
				} else {
					starty, endy = end.y, start.y
				}
				for j := starty; j <= endy; j++ {
					cave[j][start.x] = "#"
				}
			} else {
				var startx, endx int
				if start.x < end.x {
					startx, endx = start.x, end.x
				} else {
					startx, endx = end.x, start.x
				}
				for j := startx; j <= endx; j++ {
					cave[start.y][j] = "#"
				}
			}
		}
	}
	return cave
}

func printCave(cave [][]string) {
	for i, level := range cave {
		fmt.Printf("%3d %s\n", i, s.Join(level, ""))
	}
}

func part1() {
	lines := s.Split(input, "\n")

	min, max := Point{1000, 1000}, Point{0, 0}

	paths := [][]Point{}

	for _, line := range lines {
		coords := s.Split(line, " -> ")
		points := []Point{}
		for _, coord := range coords {
			newPoint := parsePoint(coord)

			if newPoint.x > max.x {
				max.x = newPoint.x
			} else if newPoint.x < min.x {
				min.x = newPoint.x
			}

			if newPoint.y > max.y {
				max.y = newPoint.y
			} else if newPoint.y < min.y {
				min.y = newPoint.y
			}

			points = append(points, *newPoint)
		}
		paths = append(paths, points)

	}

	cave := [][]string{}
	for i := 0; i <= max.y+1; i++ {
		cave = append(cave, s.Split(s.Repeat(".", max.x+max.y), ""))
	}
	cave = append(cave, s.Split(s.Repeat("#", max.x+max.y), ""))

	cave = addRocks(cave, paths)

	cnt := 0
	for  {
		currPos := Point{500, 0}
		leftEdge, rightEdge := false, false
		for {
			// if currPos.y == len(cave) -1 {
			// 	// end = true
			// 	break
			// }
			leftEdge = currPos.x == 0
			rightEdge = currPos.x == len(cave[0])-1

			if cave[currPos.y+1][currPos.x] == "." {
				currPos.y++
			} else if !leftEdge && cave[currPos.y+1][currPos.x-1] == "." {
				currPos.y++
				currPos.x--
			} else if !rightEdge && cave[currPos.y+1][currPos.x+1] == "." {
				currPos.y++
				currPos.x++
			} else {
				break
			}
		}

		if currPos.y == 0 {
			break
		}

		cave[currPos.y][currPos.x] = "o"
		cnt++
	}
	printCave(cave)
	fmt.Println(min, max, cnt)
}


func main() {
	part1()
}
