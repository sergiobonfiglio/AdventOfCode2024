package utils

import "math"

type Line struct {
	A, B *Cell
	M, Q *float64 // y = mx+q
	K    *int     // x = k; in case of vertical line
}

func NewLine(a, b *Cell) *Line {

	// x = k; in case of vertical line
	if a.C == b.C {
		tmp := b.C
		return &Line{
			A: a, B: b,
			K: &tmp,
		}
	}

	// y = mx+q
	m := float64(b.R-a.R) / float64(b.C-a.C)
	// ya = mxa+q => q = ya-mxa
	q := float64(a.R) - m*float64(a.C)
	return &Line{
		A: a, B: b,
		M: &m,
		Q: &q,
	}
}
func (l *Line) IsVertical() bool {
	return l.K != nil
}

func (l *Line) CellAtCol(col int) *Cell {
	row := l.RowAtCol(col)
	if row == nil {
		panic("ops")
	}

	return NewCell(*row, col)
}

func (l *Line) RowAtCol(col int) *int {
	if l.K != nil { // vertical
		return nil
	}

	// row = m*col+q
	f := *l.M*(float64(col)) + *l.Q
	row := int(math.Round(f)) // rounding necessary for discrete grids
	return &row
}

func (l *Line) IsOnLine(c *Cell) bool {

	if l.K != nil { // vertical
		return c.C == *l.K
	}

	//y=mx+q
	return float64(c.R) == *l.M*float64(c.C)+*l.Q

	////same col
	//if l.b.C == l.a.C {
	//	return l.b.C == c.C
	//}
	////same row
	//if l.b.R == l.a.R {
	//	return l.b.R == c.R
	//}
	//
	//// (c - a_c) / (b_c - a_c) = (r - a_r) / (b_r - a_r)
	//return float64(c.C-l.a.C)/float64(l.b.C-l.a.C) == float64(c.R-l.a.R)/float64(l.b.R-l.a.R)
}
