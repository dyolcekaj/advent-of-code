package day3

import (
	"strings"
	"unicode"
)

type token int

const (
	eof token = iota
	mul
	openParen
	closeParen
	comma
	number
	ignore
	do
	dont
)

type lexer struct {
	input string
}

type lval struct {
	token token
	num   int
}

func (l *lexer) nextToken() lval {
	l.input = strings.TrimLeftFunc(l.input, unicode.IsSpace)

	if len(l.input) == 0 {
		return lval{token: eof}
	}

	if strings.HasPrefix(l.input, "mul") {
		l.input = l.input[3:]
		return lval{token: mul}
	}

	if l.input[0] == '(' {
		l.input = l.input[1:]
		return lval{token: openParen}
	}

	if l.input[0] == ')' {
		l.input = l.input[1:]
		return lval{token: closeParen}
	}

	if l.input[0] == ',' {
		l.input = l.input[1:]
		return lval{token: comma}
	}

	num := 0
	nlen := 0
	for i := 0; i < len(l.input); i++ {
		if l.input[i] < '0' || l.input[i] > '9' {
			break
		}

		nlen++
		num *= 10
		num += int(l.input[i] - '0')
	}

	if num > 0 {
		l.input = l.input[nlen:]

		return lval{
			token: number,
			num:   num,
		}
	}

	if strings.HasPrefix(l.input, "don't()") {
		l.input = l.input[7:]
		return lval{token: dont}
	}

	if strings.HasPrefix(l.input, "do()") {
		l.input = l.input[4:]
		return lval{token: do}
	}

	l.input = l.input[1:]
	return lval{token: ignore}
}

func parseMul(l *lexer) (int, bool) {
	if l.nextToken().token != openParen {
		return 0, false
	}

	at := l.nextToken()
	if at.token != number {
		return 0, false
	}

	if l.nextToken().token != comma {
		return 0, false
	}

	bt := l.nextToken()
	if bt.token != number {
		return 0, false
	}

	if l.nextToken().token != closeParen {
		return 0, false
	}

	return at.num * bt.num, true
}

func PartOne(input string) int {
	l := &lexer{input: input}

	sum := 0

	for nt := l.nextToken(); nt.token != eof; nt = l.nextToken() {
		if nt.token == mul {
			result, ok := parseMul(l)
			if ok {
				sum += result
			}
		}
	}

	return sum
}

func PartTwo(input string) int {
	l := &lexer{input: input}

	sum := 0
	enabled := true

	for nt := l.nextToken(); nt.token != eof; nt = l.nextToken() {
		switch nt.token {
		case do:
			enabled = true
		case dont:
			enabled = false
		case mul:
			if !enabled {
				continue
			}

			result, ok := parseMul(l)
			if ok {
				sum += result
			}
		default:
			continue
		}
	}

	return sum
}
