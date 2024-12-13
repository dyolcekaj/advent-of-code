package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("cmd/day01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	as, bs := make([]int, 0, 2), make([]int, 0, 2)
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.SplitN(line, "   ", 2)

		a, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		as = append(as, a)
		bs = append(bs, b)
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d\n", partOne(as, bs))
	fmt.Printf("Part two: %d\n", partTwo(as, bs))
}

func partTwo(as []int, bs []int) int {
	count := make(map[int]int)

	for _, i := range bs {
		count[i]++
	}

	sim := 0
	for _, i := range as {
		sim += i * count[i]
	}

	return sim
}

func partOne(as, bs []int) int {
	diff := func(a, b int) int {
		if a > b {
			return a - b
		}

		return b - a
	}

	sort.Ints(as)
	sort.Ints(bs)

	sum := 0
	for i := 0; i < len(as); i++ {
		sum += diff(as[i], bs[i])
	}

	return sum
}
