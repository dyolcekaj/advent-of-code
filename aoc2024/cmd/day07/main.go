package main

import (
	"fmt"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day07"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

const input = "cmd/day07/input.txt"

func main() {
	lines, err := files.ReadLines(input)
	if err != nil {
		panic(err)
	}

	results, equations := day07.ParseInput(lines)

	fmt.Printf("PartOne: %d\n", day07.PartOne(results, equations))
	fmt.Printf("PartTwo: %d\n", day07.PartTwo(results, equations))
}
