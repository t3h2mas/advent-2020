package bus

import "math/bits"

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

type Int64 struct {
	absd   uint64
	hi, lo uint64
	neg    bool
}

// NewInt64 initializes a new pre-computed inverse.
func NewInt64(d int64) Int64 {
	var neg bool
	if d < 0 {
		neg = true
		d = -d
	}
	absd := uint64(d)

	hi, r := ^uint64(0)/absd, ^uint64(0)%absd
	lo, _ := bits.Div64(r, ^uint64(0), absd)

	var c uint64 = 1
	if absd&(absd-1) == 0 {
		c++
	}
	lo, c = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, c)
	return Int64{
		absd: absd,
		hi:   hi,
		lo:   lo,
		neg:  neg,
	}
}

func (d Int64) Mod(n int64) int64 {
	var neg bool
	if n < 0 {
		n = -n
		neg = true
	}
	hi, lo := bits.Mul64(d.lo, uint64(n))
	hi += d.hi * uint64(n)
	modlo1, _ := bits.Mul64(lo, d.absd)
	mod, modlo2 := bits.Mul64(hi, d.absd)
	var c uint64
	_, c = bits.Add64(modlo1, modlo2, 0)
	mod, _ = bits.Add64(mod, 0, c)
	if neg {
		return -int64(mod)
	}
	return int64(mod)
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
