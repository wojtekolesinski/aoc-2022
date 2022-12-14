package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	s "strings"
)

//go:embed input.txt
var input string

func compare(first, second any) int {
	firstIsFloat, secondIsFloat := false, false
	var firstArr, secondArr  []any

	switch first.(type) {
	case float64:
		firstIsFloat = true
		firstArr = []any{first}
	case []any, []float64:
		firstArr = first.([]any)
	}

	switch second.(type) {
	case float64:
		secondIsFloat = true
		secondArr = []any{second}
	case []any, []float64:
		secondArr = second.([]any)
	}

	if firstIsFloat && secondIsFloat {
		return int(firstArr[0].(float64) - secondArr[0].(float64))
	}

	for i := 0; i < len(firstArr) && i < len(secondArr); i++ {
		if result := compare(firstArr[i], secondArr[i]); result != 0 {
			return result
		}
	}

	return len(firstArr) - len(secondArr)
}

func part1() {
	lines := s.Split(input, "\n")

	sum := 0

	for i := 0; i < len(lines); i += 3 {
		var first, second []any
		json.Unmarshal([]byte(lines[i]), &first)
		json.Unmarshal([]byte(lines[i+1]), &second)

		if res := compare(first, second); res < 0 {
			sum += i/3+1
		}

	}
	fmt.Println(sum)
}

func part2() {
	lines := s.Split(input, "\n")

	packets := []any{
		[]any{[]any{6.0}},
		[]any{[]any{2.0}},
	}

	for i := 0; i < len(lines); i += 3 {
		var first, second []any
		json.Unmarshal([]byte(lines[i]), &first)
		json.Unmarshal([]byte(lines[i+1]), &second)
		packets = append(packets, first, second)
	}

	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })

	result := 1

	for index, packet := range packets {
		if p := fmt.Sprint(packet); p == "[[6]]" || p == "[[2]]" {
			result *= index+1
		}
	}

	fmt.Println(result)
}

func main() {
	part2()
}
