package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func parse() []int {
	numbers := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		var x int
		fmt.Sscanf(line, "%d", &x)
		numbers = append(numbers, x)
	}
	return numbers
}

func search(arr []int, el int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			return i, true
		}
	}
	return -1, false
}

func mix(numbers []int, times int) []int {
	tracker := []int{}
	for i := range numbers {
		tracker = append(tracker, i)
	}

	for i := 1; i <= times; i++ {
		for idx, el := range numbers {
			if el == 0 {
				continue
			}

			index, _ := search(tracker, idx)
			target := (index + el) % (len(numbers) - 1)
			if target <= 0 {
				target = len(numbers) + target - 1
			}

			// remove the index
			tracker = append(tracker[:index], tracker[index+1:]...)

			// add it back
			tracker = append(tracker[:target], append([]int{idx}, tracker[target:]...)...)

		}
	}

	result := []int{}
	for _, el := range tracker {
		result = append(result, numbers[el])
	}
	return result
}

func groveCoords(numbers []int) int {
	index, _ := search(numbers, 0)

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += numbers[(index+i*1000)%len(numbers)]
	}
	return sum
}

func part1() {
	numbers := parse()

	mixed := mix(numbers, 1)
	coords := groveCoords(mixed)
	fmt.Println("Part 1: ", coords)
}

func part2() {
	decryption_key := 811589153
	numbers := parse()
	for i := range numbers {
		numbers[i] *= decryption_key
	}

	mixed := mix(numbers, 10)
	coords := groveCoords(mixed)
	fmt.Println("Part 2: ", coords)
}

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}
