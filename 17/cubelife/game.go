package cubelife

import "strings"

type Game struct {
	active cubeSet
}

func GameFromSeed(seed string) Game {
	active := newCubeSet()

	lines := strings.Split(seed, "\n")
	reverseStrings(lines)

	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '.':
				break
			case '#':
				active.Add(newCube(x, y, 0))
			}
		}
	}

	return Game{
		active,
	}
}

func reverseStrings(xs []string) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}
