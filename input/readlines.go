package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntLines(fname string) ([]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		textLine := scanner.Text()

		n, err := strconv.Atoi(textLine)
		if err != nil {
			return nil, err
		}

		lines = append(lines, n)
	}

	return lines, scanner.Err()
}