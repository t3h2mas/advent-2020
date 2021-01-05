package cubelife

type cubeSet map[cube]bool

func newCubeSet() cubeSet {
	return make(map[cube]bool)
}

func (s cubeSet) Add(c cube) {
	s[c] = true
}

func (s cubeSet) Has(c cube) bool {
	_, exists := s[c]
	return exists
}

func (s cubeSet) Remove(c cube) {
	delete(s, c)
}
