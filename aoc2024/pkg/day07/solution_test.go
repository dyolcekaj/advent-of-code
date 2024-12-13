package day07_test

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day07"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
	"strconv"
	"testing"
)

type testCase struct {
	name     string
	want     bool
	result   int
	equation []int
}

func TestHasSolution(t *testing.T) {
	t.Parallel()

	shouldBeTrue := []int{0, 1, 8}
	testCases := setupTestCases(t)

	for _, i := range shouldBeTrue {
		testCases[i].want = true
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day07.PartOneSingle(tc.result, tc.equation)
			if got != tc.want {
				t.Errorf("hasSolution(%d, %v) = %t; want %t", tc.result, tc.equation, got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	lines, err := files.ReadLines("testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	results, equations := day07.ParseInput(lines)

	got := day07.PartTwo(results, equations)

	if got != 11387 {
		t.Errorf("PartTwo = %d; want 11387", got)
	}
}

func TestPartTwoHasSolution(t *testing.T) {
	t.Parallel()

	shouldBeTrue := []int{0, 1, 3, 4, 6, 8}
	testCases := setupTestCases(t)

	for _, i := range shouldBeTrue {
		testCases[i].want = true
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := day07.PartTwoSingle(tc.result, tc.equation)
			if got != tc.want {
				t.Errorf("hasSolution(%d, %v) = %t; want %t", tc.result, tc.equation, got, tc.want)
			}
		})
	}
}

func setupTestCases(t *testing.T) []testCase {
	t.Helper()

	lines, err := files.ReadLines("testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	results, equations := day07.ParseInput(lines)

	testCases := make([]testCase, 0, len(equations))

	for i := 0; i < len(results); i++ {
		testCases = append(testCases, testCase{
			name:     "example" + strconv.Itoa(i),
			result:   results[i],
			equation: equations[i],
		})
	}

	return testCases
}
