package day05_test

import (
	"testing"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day05"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		rules [][]int
		pages [][]int
	}{
		{
			name: "example",
			want: 143,
			rules: [][]int{
				{47, 53}, {97, 13}, {97, 61}, {97, 47},
				{75, 29}, {61, 13}, {75, 53}, {29, 13},
				{97, 29}, {53, 29}, {61, 53}, {97, 53},
				{61, 29}, {47, 13}, {75, 47}, {97, 75},
				{47, 61}, {75, 61}, {47, 29}, {75, 13},
				{53, 13},
			},
			pages: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day05.PartOne(tc.rules, tc.pages)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  int
		rules [][]int
		pages [][]int
	}{
		{
			name: "example",
			want: 123,
			rules: [][]int{
				{47, 53}, {97, 13}, {97, 61}, {97, 47},
				{75, 29}, {61, 13}, {75, 53}, {29, 13},
				{97, 29}, {53, 29}, {61, 53}, {97, 53},
				{61, 29}, {47, 13}, {75, 47}, {97, 75},
				{47, 61}, {75, 61}, {47, 29}, {75, 13},
				{53, 13},
			},
			pages: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day05.PartTwo(tc.rules, tc.pages)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}
