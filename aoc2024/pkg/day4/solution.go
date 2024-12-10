package day4

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/timer"
)

var (
	mas = [3]rune{'M', 'A', 'S'}

	directions = func() [][2]int {
		d := make([][2]int, 0, 8)

		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}

				d = append(d, [2]int{i, j})
			}
		}

		return d
	}()
)

func PartOne(puzzle [][]rune) int {
	defer timer.Time("day4.PartOne")()

	count := 0

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == 'X' {
				if num := countXMAS(puzzle, i, j); num > 0 {
					count += num
				}
			}
		}
	}

	return count
}

func countXMAS(puzzle [][]rune, i int, j int) int {
	count := 0

	for _, d := range directions {
		if spellsMAS(puzzle, i+d[0], j+d[1], d[0], d[1]) {
			count++
		}
	}

	return count
}

func PartTwo(puzzle [][]rune) int {
	defer timer.Time("day4.PartTwo")()

	count := 0

	for i := 1; i < len(puzzle)-1; i++ {
		for j := 1; j < len(puzzle[i])-1; j++ {
			if puzzle[i][j] == 'A' {
				if isXShaped(puzzle, i, j) {
					count++
				}
			}
		}
	}

	return count
}

func isXShaped(puzzle [][]rune, i, j int) bool {
	for _, d := range directions {
		// not an X
		if d[0] == 0 || d[1] == 0 {
			continue
		}

		if !spellsMAS(puzzle, i-d[0], j-d[1], d[0], d[1]) {
			continue
		}

		if spellsMAS(puzzle, i+d[0], j-d[1], -d[0], d[1]) {
			return true
		}

		if spellsMAS(puzzle, i-d[0], j+d[1], d[0], -d[1]) {
			return true
		}
	}

	return false
}

func spellsMAS(puzzle [][]rune, i, j, di, dj int) bool {
	if i < 0 || i >= len(puzzle) {
		return false
	}

	if j < 0 || j >= len(puzzle[i]) {
		return false
	}

	if minmax := i + 2*di; minmax < 0 || minmax >= len(puzzle) {
		return false
	}

	if minmax := j + 2*dj; minmax < 0 || minmax >= len(puzzle[i]) {
		return false
	}

	word := [3]rune{
		puzzle[i][j],
		puzzle[i+di][j+dj],
		puzzle[i+2*di][j+2*dj],
	}

	return mas == word
}
