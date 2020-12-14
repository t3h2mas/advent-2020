package seats

func neighborsOf(seat Point) []Point {
	var results []Point
	for xAway := -1; xAway < 2; xAway++ {
		for yAway := -1; yAway < 2; yAway++ {
			if xAway == 0 && yAway == 0 {
				continue
			}
			results = append(
				results,
				NewPoint(
					seat.X()+xAway,
					seat.Y()+yAway,
				),
			)
		}
	}
	return results
}

func calculateOccupied(layout Layout) PointSet {
	result := make(PointSet)

	for y := 0; y < layout.height; y++ {
		for x := 0; x < layout.width; x++ {
			seat := NewPoint(x, y)

			if layout.floor.Has(seat) {
				continue
			}

			occupiedNeighbors := 0

			for _, neighbor := range neighborsOf(seat) {
				if layout.occupied.Has(neighbor) {
					if seat.X() < 0 || seat.X() >= layout.width {
						// out of range
						continue
					}

					if seat.Y() < 0 || seat.Y() >= layout.height {
						// out of range
						continue
					}
					occupiedNeighbors++
				}
			}

			if layout.occupied.Has(seat) {
				// if a seat is occupied, and four or more seats adjacent are also occupied, the seat becomes empty
				if occupiedNeighbors < 4 {
					result.Add(seat)
				}
			} else {
				// if a seat is empty and there are no occupied neighbors the seat becomes occupied
				if occupiedNeighbors == 0 {
					result.Add(seat)
				}
			}
			if !layout.occupied.Has(seat) && occupiedNeighbors == 0 {
				result.Add(seat)
			}
		}
	}

	return result
}
