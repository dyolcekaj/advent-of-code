package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day11"
)

const (
	input = "30 71441 3784 580926 2 8122942 0 291"
)

func main() {
	fields := strings.Fields(input)

	nums := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	fmt.Printf("PartOne: %d\n", day11.PartOne(25, nums))
	fmt.Printf("PartTwo: %d\n", day11.PartTwo(75, nums))
}
