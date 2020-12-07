package boarding

import (
	"strconv"
	"strings"
)

type Pass struct {
	row int
	col int
}

func (p *Pass) SeatID() int {
	return p.row*8 + p.col
}

func NewPass(seat string) (*Pass, error) {
	// convert seat to binary string F:0, B: 1, R: 1, L: 0

	var b strings.Builder
	for _, ch := range seat {
		if ch == 'F' || ch == 'L' {
			b.WriteRune('0')
		}

		if ch == 'B' || ch == 'R' {
			b.WriteRune('1')
		}
	}

	bin := b.String()

	row, err := calculateRow(bin)

	if err != nil {
		return nil, err
	}

	col, err := calculateCol(bin)
	if err != nil {
		return nil, err
	}

	return &Pass{
		row,
		col,
	}, nil
}

func calculateRow(seat string) (int, error) {
	res, err := strconv.ParseInt(seat[0:7], 2, 64)

	return int(res), err
}

func calculateCol(seat string) (int, error) {
	res, err := strconv.ParseInt(seat[7:], 2, 64)

	return int(res), err
}
