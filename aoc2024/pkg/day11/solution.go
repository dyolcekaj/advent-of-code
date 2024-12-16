package day11

import (
	"math"
)

func PartOne(blinks int, stones []int) int {
	q := queue(make([]int, len(stones)))
	copy(q, stones)
	q = q.Push(-1) // sentinel

	i := 0
	for i < blinks {
		var value int

		switch value, q = q.Pop(); value {
		case -1:
			q = q.Push(-1)
			i++
		case 0:
			q = q.Push(1)
		default:
			nums := splitOrMultiply(value)

			for _, n := range nums {
				q = q.Push(n)
			}
		}
	}

	return q.Len() - 1 // remove sentinel
}

func PartTwo(blinks int, stones []int) int {
	memo := make(map[int]map[int]int)

	sum := 0

	for _, stone := range stones {
		sum += blink(stone, blinks, memo)
	}

	return sum
}

func blink(stone, times int, memo map[int]map[int]int) int {
	if times == 0 {
		return 1
	}

	if val, ok := memo[stone][times]; ok {
		return val
	}

	count := 0

	if stone == 0 {
		count = blink(1, times-1, memo)
	} else if digits := numDigits(stone); digits%2 == 0 {
		e := int(math.Pow(10, float64(digits/2)))

		count = blink(stone/e, times-1, memo) + blink(stone%e, times-1, memo)
	} else {
		count = blink(stone*2024, times-1, memo)
	}

	if _, ok := memo[stone]; !ok {
		memo[stone] = make(map[int]int)
	}
	memo[stone][times] = count

	return count
}

type queue []int

func (q queue) Len() int {
	return len(q)
}

func (q queue) Push(n int) queue {
	return append(q, n)
}

func (q queue) Pop() (int, queue) {
	return q[0], q[1:]
}

func (q queue) Peek() int {
	return q[0]
}

func splitOrMultiply(n int) []int {
	digits := numDigits(n)
	if digits%2 != 0 {
		return []int{n * 2024}
	}

	e := int(math.Pow(10, float64(digits/2)))

	return []int{n / e, n % e}
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0

	for n != 0 {
		n /= 10
		count++
	}

	return count
}
