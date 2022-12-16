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

type Sensor struct {
	*Point
	beacon       Point
	distToBeacon int
	leftEdges    [][]Point
	rightEdges   [][]Point
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func NewSensor(x, y, bcnX, bcnY int) *Sensor {
	s := Sensor{
		Point:  NewPoint(x, y),
		beacon: *NewPoint(bcnX, bcnY),
	}
	s.distToBeacon = ManhattanDistance(*s.Point, s.beacon)
	s.ComputeEdges()
	return &s
}

func (s *Sensor) ComputeEdges() {
	offset := s.distToBeacon + 1
	top := NewPoint(s.x, s.y+offset)
	right := NewPoint(s.x+offset, s.y)
	bottom := NewPoint(s.x, s.y-offset)
	left := NewPoint(s.x-offset, s.y)

	s.leftEdges = append(s.leftEdges, []Point{*bottom, *left}, []Point{*right, *top})
	s.rightEdges = append(s.rightEdges, []Point{*left, *top}, []Point{*bottom, *right})
}

func (s Sensor) String() string {
	return fmt.Sprintf("{{x: %d y: %d} bcn:{x: %d y: %d} dist: %d lE: %s, rE: %s}",
		s.x,
		s.y,
		s.beacon.x,
		s.beacon.y,
		s.distToBeacon,
		fmt.Sprint(s.leftEdges),
		fmt.Sprint(s.rightEdges),
	)
}

func ManhattanDistance(first, second Point) int {
	return IntAbs(first.x-second.x) + IntAbs(first.y-second.y)
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func parseSensors() []Sensor {
	lines := s.Split(input, "\n")
	sensors := []Sensor{}

	for _, line := range lines {
		var x, y, bcnX, bcnY int
		words := s.Split(line, " ")
		fmt.Sscanf(words[2], "x=%d", &x)
		fmt.Sscanf(words[3], "y=%d", &y)
		fmt.Sscanf(words[8], "x=%d", &bcnX)
		fmt.Sscanf(words[9], "y=%d", &bcnY)

		sensors = append(sensors, *NewSensor(x, y, bcnX, bcnY))
	}
	return sensors
}


func CheckIntersection(line1, line2 []Point) (bool, Point) {
	leftX := IntMax(IntMin(line1[0].x, line1[1].x), IntMin(line2[0].x, line2[1].x))
	rightX := IntMin(IntMax(line1[0].x, line1[1].x), IntMax(line2[0].x, line2[1].x))
	bottomY := IntMax(IntMin(line1[0].y, line1[1].y), IntMin(line2[0].y, line2[1].y)) 
	topY := IntMin(IntMax(line1[0].y, line1[1].y), IntMax(line2[0].y, line2[1].y))

	if rightX - leftX < 0 || topY - bottomY < 0 {
		return false, Point{}
	}

	makeF := func (x1, x2 Point) func (int) int {
		a := (x2.y - x1.y) / (x2.x - x1.x)
		b := x1.y - a * x1.x
		// fmt.Println(x1, x2, fmt.Sprintf("%dx + %d", a, b))
		return func (x int) int {
			return a * x + b
		}
	}

	f1 := makeF(line1[0], line1[1])
	f2 := makeF(line2[0], line2[1])

	for i := leftX; i <= rightX; i++ {
		if f1(i) == f2(i) {
			return true, *NewPoint(i, f1(i))
		}
	}

	return false, Point{}
}

func part1() {
	sensors := parseSensors()
	rowNo := 2000000
	impossiblePoints := map[Point]bool{}
	for _, sensor := range sensors {
		if dy := IntAbs(sensor.y - rowNo); dy < sensor.distToBeacon {
			fmt.Println(sensor)
			dx := sensor.distToBeacon - dy
			for x := -dx; x <= dx; x++ {
				point := *NewPoint(sensor.x+x, rowNo)
				if sensor.beacon != point {
					impossiblePoints[point] = true
				}
			}
		}
	}

	fmt.Println(len(impossiblePoints))
}

func checkIfCovered(point Point, sensors []Sensor) bool {
	for _, s := range sensors {
		if ManhattanDistance(point, *s.Point) <= s.distToBeacon {
			return true
		}
	}
	return false
}

func part2() {
	sensors := parseSensors()

	// for i := range sensors {
	// 	fmt.Println(sensors[i])
	// }
	min, max := 0, 4000000
	intersections := []Point{}
	for i, s1 := range sensors {
		for j, s2 := range sensors {
			if i == j {
				continue
			}
			for _, e1 := range s1.leftEdges {
				for _, e2 := range s2.rightEdges {
					if ok, point := CheckIntersection(e1, e2); ok {
						if point.x >= min && point.x <= max && point.y >= min && point.y <= max {
							
							if !checkIfCovered(point, sensors) {
								fmt.Println(point.x, point.y, point.x * 4000000 + point.y)
								return
							}
						}
					} 
				}
			}
		}
	}
	fmt.Println(intersections)
}

func main() {
	// part1()
	part2()
}
