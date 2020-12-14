package seats

type Point struct {
	x int
	y int
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type PointSet map[Point]bool

func NewPointSet() PointSet {
	return make(map[Point]bool)
}

func (s PointSet) Has(val Point) bool {
	_, exists := s[val]
	return exists
}

func (s PointSet) Add(val Point) {
	s[val] = true
}

type Layout struct {
	occupied []Point
	floor    PointSet
}

func NewLayout() *Layout {
	return &Layout{
		occupied: []Point{},
		floor:    NewPointSet(),
	}
}

func LayoutFrom(grid []string) *Layout {
	result := NewLayout()
	for y := range grid {
		for x := range grid[y] {
			p := NewPoint(x, y)
			switch grid[y][x] {
			case 'L':
				result.occupied = append(result.occupied, p)
			case '.':
				result.floor.Add(p)
			}
		}
	}

	return result
}
