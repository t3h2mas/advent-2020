package solution

import (
	"errors"
	"strings"
)

// LineParts transforms a line like "it parses a line with two contains clauses"
// into a nested array of strings like [[light, red, bags], [1, bright white bag], ...]
func LineParts(line string) [][]string {
	var parts [][]string
	for idx, chunk := range strings.Split(line, ", ") {
		if idx == 0 {
			rootParts := strings.Split(chunk, " contain ")
			for _, rootPart := range rootParts {
				parts = append(parts, strings.Split(rootPart, " "))
			}
		} else {
			parts = append(parts, strings.Split(chunk, " "))
		}
	}

	return parts
}

type Bag struct {
	Color    string
	Contains []string
}

func BagClauses(parts [][]string) *Bag {
	// parse root
	if len(parts) == 0 {
		panic(errors.New("cannot parse empty parts"))
	}

	result := &Bag{}

	root := parts[0]
	result.Color = strings.Join([]string{root[0], root[1]}, " ")

	// handle additional clauses
	for i := 1; i < len(parts); i++ {
		result.Contains = append(
			result.Contains,
			strings.Join([]string{parts[i][1], parts[i][2]}, " "),
		)
	}

	return result
}
