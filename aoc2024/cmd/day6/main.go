package main

import (
	"fmt"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day6"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

func main() {
	grid, err := files.ReadRunes("cmd/day6/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day6.PartOne(grid))
	fmt.Printf("PartTwo: %d\n", day6.PartTwo(grid))
}
