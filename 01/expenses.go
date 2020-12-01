package expenses

import "errors"

/**
 * Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.
 * Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
 */

func FixReport(entries []int) int {
	return 0
}

func sumsTo(target int, search []int) ([2]int, error) {
	var result [2]int

	for i := 0; i < len(search); i++ {
		for j := i + 1; j < len(search); j++ {
			if search[i] + search[j] == target {
				result[0] = search[i]
				result[1] = search[j]
				return result, nil
			}
		}
	}

	return result, errors.New("unable to find two numbers that sum to target")
}
