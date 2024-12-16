package day08

import "github.com/dyolcekaj/advent-of-code/aoc2024/pkg/grids"

func PartOne(grid [][]rune) int {
	signals := make(map[rune][]grids.Point2)
	antinodes := make(map[grids.Point2]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			curr := grids.P2(i, j)

			if points, ok := signals[grid[i][j]]; ok {
				for _, point := range points {
					dx := curr.X - point.X
					dy := curr.Y - point.Y

					if valid(curr.X+dx, curr.Y+dy, grid) {
						antinodes[grids.P2(curr.X+dx, curr.Y+dy)] = struct{}{}
					}
					if valid(point.X-dx, point.Y-dy, grid) {
						antinodes[grids.P2(point.X-dx, point.Y-dy)] = struct{}{}
					}
				}
			}

			signals[grid[i][j]] = append(signals[grid[i][j]], grids.P2(i, j))
		}
	}

	return len(antinodes)
}

func valid(x, y int, grid [][]rune) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x])
}

func PartTwo(grid [][]rune) int {
	signals := make(map[rune][]grids.Point2)
	antinodes := make(map[grids.Point2]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			curr := grids.P2(i, j)

			if points, ok := signals[grid[i][j]]; ok {
				for _, point := range points {
					dx := curr.X - point.X
					dy := curr.Y - point.Y

					for x, y := curr.X, curr.Y; valid(x, y, grid); x, y = x+dx, y+dy {
						antinodes[grids.P2(x, y)] = struct{}{}
					}

					for x, y := point.X, point.Y; valid(x, y, grid); x, y = x-dx, y-dy {
						antinodes[grids.P2(x, y)] = struct{}{}
					}
				}
			}

			signals[grid[i][j]] = append(signals[grid[i][j]], grids.P2(i, j))
		}
	}

	return len(antinodes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
