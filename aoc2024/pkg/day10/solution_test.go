package day10_test

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day10"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/grids"
	"strings"
	"testing"
)

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  36,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			runes, err := files.ReadRunesFromReader(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			grid := grids.FromRunes(runes)

			got := day10.PartOne(grid)

			if got != tc.want {
				t.Errorf("PartOne = %d; want %d", got, tc.want)
			}
		})

	}
}
