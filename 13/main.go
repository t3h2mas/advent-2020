package main

import (
	"fmt"
	"strconv"

	"github.com/t3h2mas/advent-2020/13/bus"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	notes, err := input.ReadLines("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day thirteen, part one")
	earliestDeparture, err := strconv.Atoi(notes[0])
	if err != nil {
		panic(err)
	}

	buses := bus.BusesFrom(notes[1])

	departure := bus.ClosestDeparture(buses, earliestDeparture)

	minutesToWait := departure.Time() - earliestDeparture
	fmt.Printf("solution: %d\n", minutesToWait*departure.BusID())
}
