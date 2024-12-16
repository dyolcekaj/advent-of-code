package day11_test

import (
	"testing"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day11"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		stones []int
		blinks int
		want   int
	}{
		{
			name:   "example",
			stones: []int{0, 1, 10, 99, 999},
			blinks: 1,
			want:   7,
		},
		{
			name:   "my input",
			stones: []int{30, 71441, 3784, 580926, 2, 8122942, 0, 291},
			blinks: 25,
			want:   191690,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day11.PartOne(tc.blinks, tc.stones)

			if got != tc.want {
				t.Errorf("PartOne = %d; want %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		stones []int
		blinks int
		want   int
	}{
		{
			name:   "example",
			stones: []int{0, 1, 10, 99, 999},
			blinks: 1,
			want:   7,
		},
		{
			name:   "my input",
			stones: []int{30, 71441, 3784, 580926, 2, 8122942, 0, 291},
			blinks: 25,
			want:   191690,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day11.PartTwo(tc.blinks, tc.stones)

			if got != tc.want {
				t.Errorf("PartTwo = %d; want %d", got, tc.want)
			}
		})
	}
}
