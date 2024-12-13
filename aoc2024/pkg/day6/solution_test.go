package day6_test

import (
	"strings"
	"testing"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day6"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

var exampleGrid = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "example",
			want:  41,
			input: exampleGrid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input, err := files.ReadRunesFromReader(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			got := day6.PartOne(input)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

/*
6,3
7,6
7,7
8,1
8,3
9,7
*/

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "example",
			want:  6,
			input: exampleGrid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input, err := files.ReadRunesFromReader(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			got := day6.PartTwo(input)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
