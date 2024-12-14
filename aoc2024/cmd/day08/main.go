package main

import (
	"fmt"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day08"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

const input = "cmd/day08/input.txt"

func main() {
	grid, err := files.ReadRunes(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d\n", day08.PartOne(grid))
	fmt.Printf("Part two: %d\n", day08.PartTwo(grid))
}
