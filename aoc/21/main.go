package main

import (
	_ "embed"
	"fmt"
	"time"

	"reflect"
	"strings"

	"aoc/util"
)

func parse() map[string]any {
	monkeys := make(map[string]any)

	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, ": ")
		if len(words[1]) < 10 {
			var number int
			fmt.Sscanf(words[1], "%d", &number)
			monkeys[words[0]] = float64(number)
		} else {
			monkeys[words[0]] = words[1]
		}
	}

	return monkeys
}

func compute(name string, monkeys map[string]any) float64 {
	monkey := monkeys[name]

	switch monkey.(type) {
	case float64:
		return monkey.(float64)
	case string:
		words := strings.Split(monkey.(string), " ")
		first, second, operator := words[0], words[2], words[1]
		switch operator {
		case "+":
			monkeys[name] = compute(first, monkeys) + compute(second, monkeys)
		case "*":
			monkeys[name] = compute(first, monkeys) * compute(second, monkeys)
		case "-":
			monkeys[name] = compute(first, monkeys) - compute(second, monkeys)
		case "/":
			monkeys[name] = compute(first, monkeys) / compute(second, monkeys)
		default:
			panic("")
		}
	}
	return monkeys[name].(float64)
}

func part1() {
	monkeys := parse()

	fmt.Println(int(compute("root", monkeys)))
}

func precompute(monkeys map[string]any) {
	done := false
	prev := 0
	for !done {
		for name, monkey := range monkeys {
			switch monkey.(type) {
			case float64:
				continue
			case string:
				if strings.Contains(monkey.(string), "humn") {
					continue
				} else {
					words := strings.Split(monkey.(string), " ")
					first, second, operator := words[0], words[2], words[1]
					if reflect.TypeOf(monkeys[first]) == reflect.TypeOf(0.0) && reflect.TypeOf(monkeys[second]) == reflect.TypeOf(0.0) {
						switch operator {
						case "+":
							monkeys[name] = monkeys[first].(float64) + monkeys[second].(float64)
						case "*":
							monkeys[name] = monkeys[first].(float64) * monkeys[second].(float64)
						case "-":
							monkeys[name] = monkeys[first].(float64) - monkeys[second].(float64)
						case "/":
							monkeys[name] = monkeys[first].(float64) / monkeys[second].(float64)
						}
					}
				}
			}
		}

		done = true
		count := 0
		for _, m := range monkeys {
			switch m.(type) {
			case string:
				if strings.Contains(m.(string), "humn") {
					continue
				} else {
					count++
				}
			}
		}
		if count != 0 {
			if prev == len(monkeys)-count {
				break
			}
			prev = len(monkeys) - count
			done = false

		}
	}
}

func part2() {
	root := strings.Split(parse()["root"].(string), " ")
	first, second := root[0], root[2]
	start, end := 1, 5000000000000
	for {
		humn := (end + start) / 2

		monkeys := parse()
		precompute(monkeys)
		monkeys["humn"] = float64(humn)

		diff := compute(first, monkeys) - compute(second, monkeys)
		fmt.Println("Trying: ", int(humn), " got: ", int(diff))

		if diff == 0 {
			fmt.Println("Found value: ", int(humn))
			break
		} else if diff < 0 {
			end = humn
		} else {
			start = humn + 1
		}
	}

}

//go:embed input.txt
var input string

func main() {
	func() {
		defer util.TimeTrack(time.Now(), "part1")
		part1()
	}()

	func() {
		defer util.TimeTrack(time.Now(), "part2")
		part2()
	}()
}
