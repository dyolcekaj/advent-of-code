package day05

import (
	"sort"

	"slices"
)

func PartOne(rules [][]int, pages [][]int) int {
	lessThan := buildRuleSet(rules)

	sum := 0
	for _, page := range pages {
		if isOrderedManual(page, lessThan) {
			sum += page[len(page)/2]
		}
	}

	return sum
}

func PartTwo(rules [][]int, pages [][]int) int {
	lessThan := buildRuleSet(rules)

	sum := 0
	for _, page := range pages {
		if isOrderedManual(page, lessThan) {
			continue
		}

		ordered := reorderManual(page, lessThan)

		sum += ordered[len(ordered)/2]
	}

	return sum
}

func buildRuleSet(rules [][]int) map[int][]int {
	lessThan := map[int][]int{}

	for _, rule := range rules {
		if _, ok := lessThan[rule[0]]; !ok {
			lessThan[rule[0]] = []int{}
		}

		lessThan[rule[0]] = append(lessThan[rule[0]], rule[1])
	}

	return lessThan
}

func reorderManual(pages []int, rules map[int][]int) []int {
	clone := make([]int, len(pages))
	copy(clone, pages)

	sort.SliceStable(clone, func(i, j int) bool {
		return slices.Contains(rules[clone[i]], clone[j])
	})

	return clone
}

func isOrderedManual(pages []int, rules map[int][]int) bool {
	seen := map[int]struct{}{}

	for _, page := range pages {
		for _, less := range rules[page] {
			if _, ok := seen[less]; ok {
				return false
			}
		}

		seen[page] = struct{}{}
	}

	return true
}
