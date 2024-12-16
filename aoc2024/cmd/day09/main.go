package main

import (
	"fmt"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day09"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

const input = "cmd/day09/input.txt"

func main() {
	lines, err := files.ReadRunes(input)
	if err != nil {
		panic(err)
	}

	if len(lines) != 1 {
		panic("weird input")
	}

	diskMap := make([]int, len(lines[0]))
	for i, r := range lines[0] {
		diskMap[i] = int(r - '0')
	}

	fmt.Printf("Part one: %d\n", day09.PartOne(diskMap))
	fmt.Printf("Part two: %d\n", day09.PartTwo(diskMap))
}
