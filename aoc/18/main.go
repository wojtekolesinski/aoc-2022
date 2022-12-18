package main

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/util"
	. "aoc/util/math"
)

//go:embed input.txt
var input string

type Point3D struct {
	x, y, z int
}

type Cube struct {
	Point3D
	sides int
}

func NewPoint(x, y, z int) *Point3D {
	return &Point3D{x, y, z}
}

func NewCube(x, y, z int) *Cube {
	return &Cube{
		Point3D: *NewPoint(x, y, z),
		sides:   6,
	}
}

func (p Cube) IsNextTo(other Cube) bool {
	return p.x == other.x && p.y == other.y && IntAbs(p.z-other.z) == 1 ||
		p.x == other.x && IntAbs(p.y-other.y) == 1 && p.z == other.z ||
		IntAbs(p.x-other.x) == 1 && p.y == other.y && p.z == other.z
}

func parse() []Cube {
	cubes := make([]Cube, 0)

	for _, line := range strings.Split(input, "\n") {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		// shift the cubes so they're not on the edge
		cubes = append(cubes, *NewCube(x+1, y+1, z+1))
	}

	return cubes
}

func part1() {
	cubes := parse()

	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			if cubes[i].IsNextTo(cubes[j]) {
				cubes[i].sides--
				cubes[j].sides--
			}
		}
	}

	sum := 0
	for _, cube := range cubes {
		sum += cube.sides
	}
	fmt.Println(sum)
}

func part2() {
	cubes := parse()

	size := 30
	grid := make([][][]Cube, size)

	for x := 0; x < len(grid); x++ {
		grid[x] = make([][]Cube, size)

		for y := 0; y < len(grid[x]); y++ {
			grid[x][y] = make([]Cube, size)

			for z := 0; z < len(grid[x][y]); z++ {
				grid[x][y][z] = Cube{Point3D: *NewPoint(x, y, z)}
			}
		}
	}

	for _, cube := range cubes {
		grid[cube.x][cube.y][cube.z] = cube
	}

	queue := util.Queue[Cube]{}
	queue.Add(grid[0][0][0])

	visited := make(map[Cube]bool)
	sides := 0
	for !queue.IsEmpty() {
		current:= queue.Pop()

		// check if visited already
		if _, ok := visited[current]; ok {
			continue
		}

		visited[current] = true

		x, y, z := current.x, current.y, current.z
		neighbours := make([]Cube, 0)
		if x != 0 {
			neighbours = append(neighbours, grid[x-1][y][z])
		}
		if x != size-1 {
			neighbours = append(neighbours, grid[x+1][y][z])
		}
		if y != 0 {
			neighbours = append(neighbours, grid[x][y-1][z])
		}
		if y != size-1 {
			neighbours = append(neighbours, grid[x][y+1][z])
		}
		if z != 0 {
			neighbours = append(neighbours, grid[x][y][z-1])
		}
		if z != size-1 {
			neighbours = append(neighbours, grid[x][y][z+1])
		}

		for _, neighbour := range neighbours {
			if neighbour.sides != 0 {
				sides++
			} else {
				queue.Add(neighbour)
			}
		}
	}

	fmt.Println(sides)

}

func main() {
	part1()
	part2()
}
