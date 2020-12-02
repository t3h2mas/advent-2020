package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/02/passwords"
	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	passwordEntries, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day two, part one")
	solution := passwords.ValidPasswordsByPolicy(passwordEntries)

	fmt.Printf("solution: %d\n", solution)
}
