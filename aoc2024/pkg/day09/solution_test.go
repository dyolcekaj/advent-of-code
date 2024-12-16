package day09_test

import (
	"testing"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day09"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		input []int
	}{
		{
			name:  "example",
			want:  60,
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "example plus one",
			want:  60,
			input: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "long example",
			want:  1928,
			input: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day09.PartOne(tc.input)

			if got != tc.want {
				t.Errorf("PartOne = %d; want %d", got, tc.want)
			}
		})
	}
}

// 00992111777.44.333....5555.6666.....8888..
// 00992111777.44.333....5555.6666.....8888..
func TestPartTwo(t *testing.T) {
	t.Parallel()

	// 12345
	// 0 12 345 6789 10 11 12 13 14
	// 0 .. 111 .... 2 2 2 2 2
	testCases := []struct {
		name  string
		want  int
		input []int
	}{
		//{
		//	name:  "example",
		//	want:  132,
		//	input: []int{1, 2, 3, 4, 5},
		//},
		{
			name:  "long example",
			want:  2858,
			input: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
		},
	}

	// i00 a998 i111 a888 i2 a777 i333 a6 i44 a6 i5555 m66
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day09.PartTwo(tc.input)

			if got != tc.want {
				t.Errorf("PartTwo = %d; want %d", got, tc.want)
			}
		})
	}
}
