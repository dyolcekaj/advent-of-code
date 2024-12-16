package grids

import (
	"fmt"
	"strings"
)

func PrettyPrint[T any](grid [][]T) {
	var sb strings.Builder

	for _, row := range grid {
		for _, cell := range row {
			sb.WriteString(fmt.Sprintf("%v ", cell))
		}

		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}

func FromLines(lines []string) [][]int {
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = runesToInts([]rune(line))
	}

	return grid
}

func FromRunes(runes [][]rune) [][]int {
	grid := make([][]int, len(runes))

	for i, line := range runes {
		grid[i] = runesToInts(line)
	}

	return grid
}

func runesToInts(runes []rune) []int {
	ints := make([]int, len(runes))

	for i, r := range runes {
		if r > '9' || r < '0' {
			panic("invalid rune: " + string(r))
		}

		ints[i] = int(r - '0')
	}

	return ints
}
