package main

import (
	"fmt"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day4"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

func main() {
	puzzle, err := files.ReadRunes("cmd/day4/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day4.PartOne(puzzle))
	fmt.Printf("PartTwo: %d\n", day4.PartTwo(puzzle))
}
