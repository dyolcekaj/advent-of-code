package main

import (
	"fmt"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day04"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

func main() {
	puzzle, err := files.ReadRunes("cmd/day04/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day04.PartOne(puzzle))
	fmt.Printf("PartTwo: %d\n", day04.PartTwo(puzzle))
}
