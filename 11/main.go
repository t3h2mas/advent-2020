package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/11/seats"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	gridLines, err := input.ReadLines("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day eleven, part one")
	layout := seats.LayoutFrom(gridLines)
	occupied := seats.SimulateStableOccupiedSeatCount(layout)

	fmt.Printf("solution: %d\n", occupied)
}
