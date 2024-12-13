package main

import (
	"fmt"
	"os"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day03"
)

func main() {
	input, err := os.ReadFile("cmd/day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day03.PartOne(string(input)))
	fmt.Printf("PartOne: %d\n", day03.PartTwo(string(input)))
}
