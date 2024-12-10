package day4_test

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day4"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		want      int
		inputFile string
	}{
		{
			name:      "example",
			want:      18,
			inputFile: "testdata/example.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input, err := files.ReadRunes(tc.inputFile)
			if err != nil {
				t.Error(err)
			}

			got := day4.PartOne(input)

			if got != tc.want {
				t.Errorf("got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		want      int
		inputFile string
	}{
		{
			name:      "example",
			want:      9,
			inputFile: "testdata/example.txt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input, err := files.ReadRunes(tc.inputFile)
			if err != nil {
				t.Error(err)
			}

			got := day4.PartTwo(input)

			if got != tc.want {
				t.Errorf("got %d, wanted %d", got, tc.want)
			}
		})
	}
}
