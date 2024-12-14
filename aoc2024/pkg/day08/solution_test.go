package day08_test

import (
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day08"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
	"strings"
	"testing"
)

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPartOne(t *testing.T) {
	runes, err := files.ReadRunesFromReader(strings.NewReader(example))
	if err != nil {
		t.Fatal(err)
	}

	got := day08.PartOne(runes)

	if got != 14 {
		t.Errorf("PartOne = %d; want 14", got)
	}
}
