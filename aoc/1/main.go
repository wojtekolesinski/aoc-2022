package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"aoc/util"
)

func main() {

	file, err := os.Open("aoc/1/input.txt")
	util.CheckError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var sums []int
	var sum int = 0

	for _, el := range input {
		if strings.TrimSpace(el) == "" {
			sums = append(sums, sum)
			sum = 0
		} else {
			intVal, err := strconv.Atoi(el)
			util.CheckError(err)
			sum += intVal
		}
	}

	max := findTopNSum(sums, 3)

	fmt.Println(max)
}

// part 1
func findMax(sums []int) int {
	max := 0
	for _, el := range sums {
		if el > max {
			max = el
		}
	}
	return max
}

// part 2
func findTopNSum(sums []int, n int) int {
	topN := make([]int, n)

	for _, el := range sums {
		if el < topN[0] {
			continue
		}
		for i := n - 1; i >= 0; i-- {
			if el > topN[i] {
				for j := 0; j < i; j++ {
					topN[j] = topN[j+1]
				}
				topN[i] = el
				break
			}
		}
	}

	sum := 0
	for _, el := range topN {
		sum += el
	}
	return sum
}
