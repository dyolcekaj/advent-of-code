package main

import (
	"fmt"
	"os"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day3"
)

func main() {
	input, err := os.ReadFile("cmd/day3/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("PartOne: %d\n", day3.PartOne(string(input)))
	fmt.Printf("PartOne: %d\n", day3.PartTwo(string(input)))
}
