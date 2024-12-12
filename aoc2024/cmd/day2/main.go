package main

import (
	"fmt"
	"time"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

const fileName = "cmd/day2/input.txt"

func main() {
	lines, err := files.ReadInts(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d\n", partOne(lines))
	fmt.Printf("Part two: %d\n", partTwo(lines))
}

func partTwo(data [][]int) int {
	start := time.Now()

	safe := 0

	for _, levels := range data {
		if isSafeDampenedReport(levels) {
			safe++
		}
	}

	fmt.Println(time.Since(start))

	return safe
}

func partOne(data [][]int) int {
	start := time.Now()

	safe := 0

	for _, levels := range data {
		if _, ok := isSafeReport(levels); ok {
			safe++
		}
	}

	fmt.Println(time.Since(start))

	return safe
}

func isSafeDampenedReport(levels []int) bool {
	i, ok := isSafeReport(levels)
	if ok {
		return true
	}

	// special case, sign flips in first two steps
	if (levels[1]-levels[0])*(levels[2]-levels[1]) < 0 {
		if _, ok2 := isSafeReport(levels[1:]); ok2 {
			return true
		}
	}

	// try removing both levels in the broken step
	for j := i; j >= 0 && i-j < 2; j-- {
		clone := make([]int, len(levels)-1)

		copy(clone, levels[:j])
		copy(clone[j:], levels[j+1:])

		if _, ok := isSafeReport(clone); ok {
			return true
		}
	}

	return false
}

func isSafeReport(levels []int) (int, bool) {
	prev := levels[0]
	prevDiff := 0

	for i := 1; i < len(levels); i++ {
		curr := levels[i]
		diff := curr - prev

		if !isSafeChange(diff, prevDiff) {
			return i, false
		}

		prev = curr
		prevDiff = diff
	}

	return 0, true
}

func isSafeChange(diff, prev int) bool {
	if diff*prev < 0 {
		return false
	}

	a := abs(diff)

	if a > 3 || a < 1 {
		return false
	}

	return true
}

func abs(diff int) int {
	if diff > 0 {
		return diff
	}

	return -diff
}
