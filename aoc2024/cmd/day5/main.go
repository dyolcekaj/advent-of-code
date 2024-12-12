package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/day5"
	"github.com/dyolcekaj/advent-of-code/aoc2024/pkg/files"
)

const (
	fileName = "cmd/day5/input.txt"
)

func main() {
	rules, pages := parseInput()

	fmt.Printf("PartOne: %d\n", day5.PartOne(rules, pages))
	fmt.Printf("PartTwo: %d\n", day5.PartTwo(rules, pages))
}

func parseInput() ([][]int, [][]int) {
	lines, err := files.ReadLines(fileName)
	if err != nil {
		panic(err)
	}

	rules := make([][]int, 0)
	pages := make([][]int, 0)

	i := 0
	for ; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			i++
			break
		}

		rule := make([]int, 0, 2)
		for _, n := range strings.SplitN(lines[i], "|", 2) {
			if len(n) != 2 {
				panic("invalid rule: " + lines[i])
			}

			x, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}

			rule = append(rule, x)
		}

		rules = append(rules, rule)
	}

	for ; i < len(lines); i++ {
		tokens := strings.Split(lines[i], ",")
		page := make([]int, 0, len(tokens))

		for _, token := range tokens {
			n, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			page = append(page, n)
		}

		pages = append(pages, page)
	}

	return rules, pages
}
