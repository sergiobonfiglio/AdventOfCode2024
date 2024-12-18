package utils

import (
	"fmt"
	"math"
)

type Cell struct {
	R int
	C int
}

func NewCell(r, c int) *Cell {
	return &Cell{
		R: r,
		C: c,
	}
}

func (c Cell) Up(d int) Cell {
	return Cell{
		R: c.R - d,
		C: c.C,
	}
}
func (c Cell) Down(d int) Cell {
	return Cell{
		R: c.R + d,
		C: c.C,
	}
}
func (c Cell) Left(d int) Cell {
	return Cell{
		R: c.R,
		C: c.C - d,
	}
}
func (c Cell) Right(d int) Cell {
	return Cell{
		R: c.R,
		C: c.C + d,
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("(%d,%d)", c.R, c.C)
}

func (c Cell) DistManhattan(x *Cell) int {
	return int(math.Abs(float64(c.C-x.C)) + math.Abs(float64(c.R-x.R)))
}
