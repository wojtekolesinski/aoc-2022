package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"time"

	s "strings"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

type Tetromino struct {
	points []*Point
	top    int
}

func (t *Tetromino) MoveDown(board [][]bool) bool {
	for _, p := range t.points {
		if p.y == 0 || board[p.y-1][p.x] {
			return false
		}
	}

	for _, p := range t.points {
		p.y--
	}
	t.top--
	return true
}

func (t *Tetromino) MoveSideways(dx int, board [][]bool) {
	for _, p := range t.points {
		if x := p.x + dx; x < 0 || x > len(board[0])-1 || board[p.y][x] {
			return
		}
	}

	for _, p := range t.points {
		p.x += dx
	}

}

func TetrominoGenerator() func(Point) Tetromino {
	step := 0
	return func(start Point) Tetromino {
		t := Tetromino{}
		x, y := start.x, start.y
		switch step % 5 {
		case 0:
			t.points = []*Point{NewPoint(x, y), NewPoint(x+1, y), NewPoint(x+2, y), NewPoint(x+3, y)}
			t.top = y
		case 1:
			t.points = []*Point{NewPoint(x+1, y), NewPoint(x+1, y+1), NewPoint(x, y+1), NewPoint(x+2, y+1), NewPoint(x+1, y+2)}
			t.top = y + 2
		case 2:
			t.points = []*Point{NewPoint(x, y), NewPoint(x+1, y), NewPoint(x+2, y), NewPoint(x+2, y+1), NewPoint(x+2, y+2)}
			t.top = y + 2
		case 3:
			t.points = []*Point{NewPoint(x, y), NewPoint(x, y+1), NewPoint(x, y+2), NewPoint(x, y+3)}
			t.top = y + 3
		case 4:
			t.points = []*Point{NewPoint(x, y), NewPoint(x+1, y), NewPoint(x, y+1), NewPoint(x+1, y+1)}
			t.top = y + 1
		}
		step++
		return t
	}
}

func printBoard(board [][]bool, block Tetromino) {
	for y := len(board) - 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < len(board[y]); x++ {
			b := false
			for _, p := range block.points {
				if p.x == x && p.y == y {
					b = true
					break
				}
			}
			if board[y][x] {
				fmt.Print("#")
			} else if b {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func makeAnimation(display bool, frame time.Duration) func ([][]bool ,Tetromino){
	return func(board [][]bool, block Tetromino) {
		if display {
			clearScreen()
			printBoard(board, block)
			time.Sleep(time.Millisecond * frame)
		}
	}
}

func part1() {
	steps := parse()

	board := make([][]bool, 20)

	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, 7)
	}

	generator := TetrominoGenerator()
	blocks := 2022
	top := -1
	step := 0
	frame := time.Duration(50)

	animate := makeAnimation(false, frame)

	for i := 0; i < blocks; i++ {
		animate(board, Tetromino{})

		block := generator(*NewPoint(2, top+4))

		animate(board, block)

		isMoving := true
		for isMoving {
			block.MoveSideways(steps[step], board)
			step++
			step %= len(steps)

			animate(board, block)

			isMoving = block.MoveDown(board)

			animate(board, block)
		}

		for _, p := range block.points {
			board[p.y][p.x] = true
		}
		if block.top > top {
			top = block.top
		}

		if top > len(board) - 8 {
			board = append(board, make([]bool, 7))
			board = append(board, make([]bool, 7))
			board = append(board, make([]bool, 7))
			board = append(board, make([]bool, 7))
			board = append(board, make([]bool, 7))
			board = append(board, make([]bool, 7))
		}
	}

	fmt.Println(top+1)

}

func part2() {}

//go:embed input.txt
var input string

func parse() []int {
	lines := s.Split(input, "")
	directions := make([]int, 0)

	for _, line := range lines {
		switch line {
		case ">":
			directions = append(directions, 1)
		case "<":
			directions = append(directions, -1)
		default:
			panic("Wrong input")
		}
	}

	return directions
}

func main() {
	part1()
	part2()
}
