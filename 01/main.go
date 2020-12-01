package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/01/expenses"
	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	expenseEntries, err := input.ReadIntLines("./input")
	if err != nil {
		panic(err)
	}

	solution, err := expenses.FixReport(expenseEntries)

	if err != nil {
		panic(err)
	}

	fmt.Printf("solution: %d\n", solution)
}
