package day07

import (
	"strconv"
	"strings"
)

func PartOne(results []int, equations [][]int) int {
	result := 0

	for i := 0; i < len(results); i++ {
		if PartOneSingle(results[i], equations[i]) {
			result += results[i]
		}
	}

	return result
}

func PartOneSingle(result int, equation []int) bool {
	return backtrack(result, equation[0], 1, equation, []op{mul, add})
}

func PartTwo(results []int, equations [][]int) int {
	result := 0

	for i := 0; i < len(results); i++ {
		if PartTwoSingle(results[i], equations[i]) {
			result += results[i]
		}
	}

	return result
}

func PartTwoSingle(result int, equation []int) bool {
	return backtrack(result, equation[0], 1, equation, []op{mul, add, concat})
}

func backtrack(result, curr, idx int, equation []int, ops []op) bool {
	if idx >= len(equation) {
		return curr == result
	}

	for _, operation := range ops {
		if backtrack(result, operation(curr, equation[idx]), idx+1, equation, ops) {
			return true
		}
	}

	return false
}

type op = func(a, b int) int

func mul(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func concat(a, b int) int {
	return a*padding(b) + b
}

func padding(i int) int {
	p := 10

	for p <= i {
		p *= 10
	}

	return p
}

func ParseInput(lines []string) ([]int, [][]int) {
	if len(lines) == 0 {
		return nil, nil
	}

	results := make([]int, 0, len(lines))
	equations := make([][]int, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Fields(line)

		result, err := strconv.Atoi(strings.TrimRightFunc(tokens[0], func(r rune) bool {
			return r == ':'
		}))
		if err != nil {
			panic(err)
		}

		equation := make([]int, 0, len(tokens)-1)
		for _, token := range tokens[1:] {
			i, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			equation = append(equation, i)
		}

		results = append(results, result)
		equations = append(equations, equation)
	}

	return results, equations
}
