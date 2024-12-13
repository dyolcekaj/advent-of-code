package day6

import (
	"slices"
)

type (
	coord  = struct{ x, y int }
	vector = struct{ x, y, dx, dy int }
)

func PartOne(grid [][]rune) int {
	path, ok := getPathFromStart(grid)
	if !ok {
		panic("loop")
	}

	moves := 0

	seen := make(map[coord]struct{})
	for _, v := range path {
		c := coord{v.x, v.y}
		if _, ok := seen[c]; !ok {
			moves++
		}
		seen[c] = struct{}{}
	}

	return moves
}

func PartTwo(grid [][]rune) int {
	path, ok := getPathFromStart(grid)

	if !ok {
		panic("loop")
	}

	seen := make(map[coord]struct{})
	loops := 0

	for i := 0; i < len(path)-1; i++ {
		v := path[i]

		s := step(v)
		if _, ok := seen[coord{s.x, s.y}]; ok {
			continue
		}

		if grid[s.x][s.y] == '.' {
			grid[s.x][s.y] = '#'

			_, ok := getPath(v, grid)
			if !ok {
				loops++
			}

			grid[s.x][s.y] = '.'

			seen[coord{s.x, s.y}] = struct{}{}
		}
	}

	return loops
}

func getPathFromStart(grid [][]rune) ([]vector, bool) {
	return getPath(findStartPosition(grid), grid)
}

func getPath(v vector, grid [][]rune) ([]vector, bool) {
	path := make([]vector, 0, 2)

	for valid(v.x, v.y, grid) {
		if slices.Contains(path, v) {
			return nil, false
		}

		path = append(path, v)

		s := step(v)
		if valid(s.x, s.y, grid) && grid[s.x][s.y] == '#' {
			v = rotate(v)
		} else {
			v = s
		}
	}

	return path, true
}

func step(v vector) vector {
	return vector{v.x + v.dx, v.y + v.dy, v.dx, v.dy}
}

func rotate(v vector) vector {
	return vector{v.x, v.y, v.dy, -v.dx}
}

func valid(x, y int, grid [][]rune) bool {
	return bc(x, len(grid)) && bc(y, len(grid[x]))
}

func bc(i, max int) bool {
	return i >= 0 && i < max
}

func findStartPosition(grid [][]rune) vector {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case '<':
				return vector{i, j, 0, -1}
			case '>':
				return vector{i, j, 0, 1}
			case 'v':
				return vector{i, j, 1, 0}
			case '^':
				return vector{i, j, -1, 0}
			default:
				continue
			}
		}
	}

	return vector{-1, -1, 0, 0}
}
