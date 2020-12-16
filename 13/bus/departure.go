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
