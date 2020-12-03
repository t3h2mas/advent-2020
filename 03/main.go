package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/03/slopes"

	"github.com/t3h2mas/advent-2020/03/grid"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	var slope grid.ScrollingGrid
	slope, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day three, part one")

	direction := slopes.NewTobogganDirection(3, 1)

	encounters := slopes.PathEncounters(slope, direction)

	treeCount := encounters['#']
	fmt.Printf("solution: %d\n", treeCount)

	fmt.Println("Day three, part two")

}
