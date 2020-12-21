package bus

type Departure struct {
	busID int
	time  int
}

func (d Departure) BusID() int {
	return d.busID
}

func (d Departure) Time() int {
	return d.time
}

func ClosestDeparture(buses []Bus, targetDeparture int) Departure {
	var closestDeparture int
	var busID int

	for idx, bus := range buses {
		busDeparture := bus.ID()
		multiple := 1
		for busDeparture <= targetDeparture {
			busDeparture = bus.ID() * multiple
			multiple++
		}

		if idx == 0 {
			closestDeparture = busDeparture
			busID = bus.ID()
			continue
		}

		if busDeparture < closestDeparture {
			closestDeparture = busDeparture
			busID = bus.ID()
		}
	}
	return Departure{
		busID,
		closestDeparture,
	}
}

func EarliestOrderedDeparture(buses map[int]Bus) int {
	startingBus := buses[0]

	multiple := 1
	for {
		timestamp := startingBus.ID() * multiple

		match := true

		for k, v := range buses {
			// skip our startingBus
			if k == 0 {
				continue
			}

			targetTS := timestamp + k

			if targetTS%v.ID() != 0 {
				match = false
				break
			}

		}

		if match {
			return timestamp
		}

		multiple++
	}

	return 0
}
