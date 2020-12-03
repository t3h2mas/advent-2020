package main

import (
	"errors"
	"fmt"

	"github.com/t3h2mas/advent-2020/input"
)

type Grid []string

var Tree = '#'
var Open rune = '.'

var OutOfBoundsError = errors.New("OutOfBounds")

func (g Grid) CellAt(x, y int) (rune, error) {
	var cell rune
	if y < 0 || y > len(g)-1 {
		return cell, OutOfBoundsError
	}

	cell = rune(g[y][x%len(g[y])])

	return cell, nil
}

func main() {
	var slope Grid
	slope, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	rightPace := 3
	downPace := 1

	x := 0
	y := 0

	treeCount := 0

	for {
		cell, err := slope.CellAt(x, y)

		if err != nil {
			break
		}

		if cell == Tree {
			treeCount++
		}

		x += rightPace
		y += downPace
	}

	fmt.Printf("Day three, part one")
	fmt.Printf("solution: %d\n", treeCount)
}
