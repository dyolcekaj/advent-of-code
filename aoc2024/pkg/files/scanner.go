package files

import (
	"bufio"
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
	var lines []string

	f, err := os.Open(fileName)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
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
