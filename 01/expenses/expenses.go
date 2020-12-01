package expenses

import (
	"errors"
)

/**
 * Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.
 * Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
 */

func FixReport(entries []int) (int, error) {
	components, err := sumsTo(2020, entries)
	if err != nil {
		return 0, err
	}

	return components[0] * components[1], nil
}

type IntSet struct {
	data map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{
		data: make(map[int]bool),
	}
}

func (s *IntSet) Has(val int) bool {
	_, hasValue := s.data[val]
	return hasValue
}

func (s *IntSet) Add(val int) {
	s.data[val] = true
}

func sumsTo(target int, search []int) ([2]int, error) {
	var result [2]int

	seen := NewIntSet()

	for _, n := range search {
		if seen.Has(target - n) {
			result[0] = n
			result[1] = target - n
			return result, nil
		}
		seen.Add(n)
	}

	return result, errors.New("unable to find two numbers that sum to target")
}
