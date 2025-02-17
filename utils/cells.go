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

func (c Cell) Dir(x string) Cell {
	switch x {
	case "^":
		return c.Up(1)
	case "v":
		return c.Down(1)
	case "<":
		return c.Left(1)
	case ">":
		return c.Right(1)
	default:
		panic(fmt.Sprintf("Unknown direction %s", x))
	}
}

func (c Cell) GetDir(next Cell) string {
	if c.R == next.R {
		if c.C < next.C {
			return ">"
		} else if c.C > next.C {
			return "<"
		}
		panic("invalid direction")
	} else {
		if c.R < next.R {
			return "v"
		} else {
			return "^"
		}
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

func (c Cell) NeighborsCross() []Cell {
	return []Cell{
		c.Up(1),
		c.Right(1),
		c.Down(1),
		c.Left(1),
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("(%d,%d)", c.R, c.C)
}

func (c Cell) DistManhattan(x *Cell) int {
	return int(math.Abs(float64(c.C-x.C)) + math.Abs(float64(c.R-x.R)))
}
