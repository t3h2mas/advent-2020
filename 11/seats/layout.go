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
	height   int
	width    int
	occupied PointSet
	floor    PointSet
}

func newLayout() *Layout {
	return &Layout{
		occupied: NewPointSet(),
		floor:    NewPointSet(),
	}
}

func (l *Layout) setHeight(val int) {
	l.height = val
}

func (l *Layout) setWidth(val int) {
	l.width = val
}

func LayoutFrom(grid []string) *Layout {
	result := newLayout()

	result.setHeight(len(grid))
	result.setWidth(len(grid[0]))

	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case '.':
				p := NewPoint(x, y)
				result.floor.Add(p)
			}
		}
	}

	return result
}
