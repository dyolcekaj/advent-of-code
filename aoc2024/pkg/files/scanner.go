package files

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadInts(fileName string) ([][]int, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := make([][]int, 0)
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())

		row := make([]int, 0, len(tokens))

		for _, token := range tokens {
			i, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}

			row = append(row, i)
		}

		data = append(data, row)
	}

	return data, scanner.Err()
}

func ReadLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readLines(f)
}

func readLines(reader io.Reader) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadRunes(fileName string) ([][]rune, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadRunesFromReader(f)
}

func ReadRunesFromReader(reader io.Reader) ([][]rune, error) {
	lines, err := readLines(reader)
	if err != nil {
		return nil, err
	}

	data := make([][]rune, 0, len(lines))

	for _, line := range lines {
		data = append(data, []rune(line))
	}

	return data, nil
}

func ForEachLine(fileName string, fn func(line string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fn(scanner.Text())
	}

	return scanner.Err()
}
