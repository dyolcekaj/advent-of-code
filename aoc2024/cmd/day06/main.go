package main

import (
	"fmt"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day06"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

func main() {
	grid, err := files.ReadRunes("cmd/day06/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day06.PartOne(grid))
	fmt.Printf("PartTwo: %d\n", day06.PartTwo(grid))
}
