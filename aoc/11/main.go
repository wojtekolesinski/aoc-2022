package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	s "strings"
)

//go:embed input.txt
var input string

type Monkey struct {
	items 		[]int
	testNumber 	int
	ifTrue		int
	ifFalse		int
	operation	func(old int) int
	inspections	int
}

func (m *Monkey) makeTurn(first, second *Monkey) {
	itemsLen := len(m.items)
	for i := 0; i < itemsLen; i++ {
		item := m.items[0]
		m.items = m.items[1:]

		item = m.operation(item)

		if item % m.testNumber == 0 {
			first.items = append(first.items, item)
		} else {
			second.items = append(second.items, item)
		}
		m.inspections++
	}
}

func (m *Monkey) makeTurnV2(first, second *Monkey, denominator int) {
	itemsLen := len(m.items)
	for i := 0; i < itemsLen; i++ {
		item := m.items[0]
		m.items = m.items[1:]

		item = m.operation(item)
		item %= denominator

		if item % m.testNumber == 0 {
			first.items = append(first.items, item)
		} else {
			second.items = append(second.items, item)
		}
		m.inspections++
	}
}

func parseOperation(op string) func(old int) int {
	words := s.Split(op, " ")
	operator := words[1]
	right := words[2]

	rightIsNum := right != "old"
	var rightNum int;
	if rightIsNum {
		rightNum, _ = strconv.Atoi(right)
	}

	return func(old int) int {
		switch operator {
		case "+":
			if rightIsNum {
				return old + rightNum
			} else {
				return old + old
			}
		case "*":
			if rightIsNum {
				return old * rightNum
			} else {
				return old * old
			}
		default: return old
		}
	}
}

func parseInput() []*Monkey {
	lines := s.Split(input, "\n")

	monkeys := make([]*Monkey, 0)
	fmt.Println(monkeys)

	for i := 0; i < len(lines); i+= 7 {
		var itemsStr, operationStr string
		var testNumber, ifTrue, ifFalse int
		itemsStr = s.Split(lines[i+1], ": ")[1]
		operationStr = s.Split(lines[i+2], " = ")[1]
		fmt.Sscanf(lines[i+3], "  Test: divisible by %d", &testNumber)
		fmt.Sscanf(lines[i+4], "    If true: throw to monkey %d", &ifTrue)
		fmt.Sscanf(lines[i+5], "    If false: throw to monkey %d", &ifFalse)
		operation := parseOperation(operationStr)

		items := make([]int, 0)
		for _, item := range s.Split(itemsStr, ", ") {
			parsedItem, _ := strconv.Atoi(item)
			items = append(items, parsedItem)
		}

		m := Monkey{
			testNumber: testNumber,
			items: items,
			ifTrue: ifTrue,
			ifFalse: ifFalse,
			operation: operation,
		}

		monkeys = append(monkeys, &m)
	}
	return monkeys	
}

func printMonkeys(monkeys []*Monkey) {
	for i := range monkeys {
		fmt.Print(*monkeys[i], " ")
	}
	fmt.Println()
}

func part1() {
	monkeys := parseInput()
	printMonkeys(monkeys)
	
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.makeTurn(monkeys[monkey.ifTrue], monkeys[monkey.ifFalse])
			// printMonkeys(monkeys)
		}
		printMonkeys(monkeys)
	}

	inspections := []int{}
	for i := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, monkeys[i].inspections)
		inspections = append(inspections, monkeys[i].inspections)
	}

	sort.Ints(inspections)
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0]*inspections[1])

}

func part2() {
	monkeys := parseInput()
	printMonkeys(monkeys)

	denominator := 1
	for _, m := range monkeys {
		denominator *= m.testNumber
	}
	
	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			monkey.makeTurnV2(monkeys[monkey.ifTrue], monkeys[monkey.ifFalse], denominator)
			// printMonkeys(monkeys)
		}
		// printMonkeys(monkeys)
	}

	inspections := []int{}
	for i := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, monkeys[i].inspections)
		inspections = append(inspections, monkeys[i].inspections)
	}

	sort.Ints(inspections)
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0]*inspections[1])
}

func main() {
	// part1()
	part2()
}