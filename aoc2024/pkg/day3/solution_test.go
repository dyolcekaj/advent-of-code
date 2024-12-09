package day3_test

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day3"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "example",
			want:  161,
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day3.PartOne(tc.input)
			if got != tc.want {
				t.Errorf("got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "example",
			want:  48,
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day3.PartTwo(tc.input)
			if got != tc.want {
				t.Errorf("got %d, wanted %d", got, tc.want)
			}
		})
	}
}
