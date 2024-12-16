package main

import (
	"fmt"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day10"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/grids"
)

const input = "cmd/day10/input.txt"

func main() {
	runes, err := files.ReadRunes(input)
	if err != nil {
		panic(err)
	}

	grid := grids.FromRunes(runes)

	fmt.Printf("Part one: %d\n", day10.PartOne(grid))
	fmt.Printf("Part two: %d\n", day10.PartTwo(grid))
}
