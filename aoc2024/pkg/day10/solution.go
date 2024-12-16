package day10

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/grids"
)

func PartOne(grid [][]int) int {
	sum := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				sum += dfs(grid, make(map[grids.Point2]struct{}), 0, x, y)
			}
		}
	}

	return sum
}

func PartTwo(grid [][]int) int {
	memoize := make([][]int, len(grid))

	for x := 0; x < len(grid); x++ {
		memoize[x] = make([]int, len(grid[0]))

		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 9 {
				memoize[x][y] = 1
			}
		}
	}

	sum := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 0 {
				sum += countPaths(grid, memoize, 0, x, y)
			}
		}
	}

	return sum
}

func dfs(grid [][]int, seen map[grids.Point2]struct{}, value, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 0
	}

	if value > 9 {
		return 0
	}

	if grid[x][y] != value {
		return 0
	}

	if value == 9 && grid[x][y] == 9 {
		p := grids.P2(x, y)

		if _, ok := seen[p]; ok {
			return 0
		}

		seen[p] = struct{}{}

		return 1
	}

	return dfs(grid, seen, value+1, x-1, y) +
		dfs(grid, seen, value+1, x+1, y) +
		dfs(grid, seen, value+1, x, y-1) +
		dfs(grid, seen, value+1, x, y+1)
}

func countPaths(grid [][]int, memoize [][]int, value, x, y int) int {
	if value > 9 {
		return 0
	}

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 0
	}

	if grid[x][y] == value && memoize[x][y] != 0 {
		return memoize[x][y]
	}

	if grid[x][y] != value {
		return 0
	}

	memoize[x][y] = countPaths(grid, memoize, value+1, x-1, y) +
		countPaths(grid, memoize, value+1, x+1, y) +
		countPaths(grid, memoize, value+1, x, y-1) +
		countPaths(grid, memoize, value+1, x, y+1)

	return memoize[x][y]
}
