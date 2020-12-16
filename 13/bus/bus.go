package bus

import (
	"strconv"
	"strings"
)

type Bus struct {
	id int
}

func NewBus(id int) Bus {
	return Bus{id}
}

func BusesFrom(input string) []Bus {
	var result []Bus

	for _, bid := range strings.Split(input, ",") {
		if bid == "x" {
			continue
		}

		parsed, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}

		result = append(result, Bus{id: parsed})
	}

	return result
}
