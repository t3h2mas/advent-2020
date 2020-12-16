package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	notes, err := input.ReadLines("./input")

	if err != nil {
		panic(err)
	}

	for idx, note := range notes {
		fmt.Printf("%d: '%s'\n", idx, note)
	}
}
