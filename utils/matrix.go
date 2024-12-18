package utils

import (
	"strconv"
	"strings"
)

type Matrix[T any] struct {
	Matrix          [][]T
	CurrRow         int
	CurrCol         int
	hasBeenIterated bool
}

var _ Iter[int] = new(Matrix[int])

func (m *Matrix[T]) Reset() {
	m.hasBeenIterated = false
	m.CurrRow = 0
	m.CurrCol = 0
}

func (m *Matrix[T]) Next() (T, bool) {
	nextC := (m.CurrCol + 1) % len(m.Matrix[m.CurrRow])

	//skip first increment if this is the first Next()
	if !m.hasBeenIterated {
		m.hasBeenIterated = true
		nextC = m.CurrCol
	}

	if nextC < m.CurrCol {
		if m.CurrRow+1 == len(m.Matrix) {
			return *new(T), false
		}
		m.CurrRow = (m.CurrRow + 1) % len(m.Matrix)
	}
	m.CurrCol = nextC
	return m.Matrix[m.CurrRow][m.CurrCol], true
}

func NewMatrix[T any](matrix [][]T) *Matrix[T] {
	return &Matrix[T]{
		Matrix: matrix,
	}
}

func MapToInt(c rune) int {
	n, err := strconv.Atoi(string(c))
	if err != nil {
		panic("conversion error")
	}
	return n
}

func NewIntMatrixFromLines(input string) *Matrix[int] {
	matrix := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		runes := []rune(line)
		mapped := []int{}
		for i := 0; i < len(runes); i++ {
			mapped = append(mapped, MapToInt(runes[i]))
		}
		matrix = append(matrix, mapped)
	}

	return NewMatrix[int](matrix)
}
func NewMatrixFromLines(input string) *Matrix[rune] {
	matrix := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		runes := []rune(line)
		matrix = append(matrix, runes)
	}

	return NewMatrix[rune](matrix)
}

func (m *Matrix[T]) IsIn(row, col int) bool {
	if row < 0 || row >= len(m.Matrix) ||
		col < 0 || col >= len(m.Matrix[row]) {
		return false
	}
	return true
}
func (m *Matrix[T]) Set(row, col int) bool {
	if !m.IsIn(row, col) {
		return false
	}

	m.CurrRow = row
	m.CurrCol = col
	return true
}

func (m *Matrix[T]) GetAt(row int, col int) *T {
	if !m.IsIn(row, col) {
		return nil
	}
	return &m.Matrix[row][col]
}
func (m *Matrix[T]) GetAtCell(cell *Cell) *T {
	return m.GetAt(cell.R, cell.C)
}

func (m *Matrix[T]) Curr() *T {
	if !m.IsIn(m.CurrRow, m.CurrCol) {
		return nil
	}
	return &m.Matrix[m.CurrRow][m.CurrCol]
}

func (m *Matrix[T]) CurrCell() *Cell {
	if !m.IsIn(m.CurrRow, m.CurrCol) {
		return nil
	}
	return NewCell(m.CurrRow, m.CurrCol)
}

func (m *Matrix[T]) Left() *T {
	return m.LeftBy(1)
}

func (m *Matrix[T]) LeftBy(x int) *T {
	if !m.IsIn(m.CurrRow, m.CurrCol-x) {
		return nil
	}
	return &m.Matrix[m.CurrRow][m.CurrCol-x]
}

func (m *Matrix[T]) GetLeft(x int) []T {
	if !m.IsIn(m.CurrRow, m.CurrCol-x) || !m.IsIn(m.CurrRow, m.CurrCol) {
		return nil
	}
	return m.Matrix[m.CurrRow][m.CurrCol-x : m.CurrCol]
}

func (m *Matrix[T]) Right() *T {
	return m.RightBy(1)
}

func (m *Matrix[T]) RightBy(x int) *T {
	if !m.IsIn(m.CurrRow, m.CurrCol+x) {
		return nil
	}
	return &m.Matrix[m.CurrRow][m.CurrCol+x]
}

func (m *Matrix[T]) GetRight(x int) []T {
	if !m.IsIn(m.CurrRow, m.CurrCol+x) || !m.IsIn(m.CurrRow, m.CurrCol) {
		return nil
	}
	return m.Matrix[m.CurrRow][m.CurrCol : m.CurrCol+x]
}

func (m *Matrix[T]) Up() *T {
	return m.UpBy(1)
}

func (m *Matrix[T]) UpBy(x int) *T {
	if !m.IsIn(m.CurrRow-x, m.CurrCol) {
		return nil
	}
	return &m.Matrix[m.CurrRow-x][m.CurrCol]
}

func (m *Matrix[T]) GetUpBy(x int) []T {
	if !m.IsIn(m.CurrRow-x, m.CurrCol) || !m.IsIn(m.CurrRow, m.CurrCol) {
		return nil
	}
	var out []T
	for r := m.CurrRow - x; r <= m.CurrRow; r++ {
		out = append(out, m.Matrix[r][m.CurrCol])
	}
	return out
}

func (m *Matrix[T]) Down() *T {
	return m.DownBy(1)
}

func (m *Matrix[T]) DownBy(x int) *T {
	if !m.IsIn(m.CurrRow+x, m.CurrCol) {
		return nil
	}
	return &m.Matrix[m.CurrRow+x][m.CurrCol]
}

func (m *Matrix[T]) UpLeft() *T {
	return m.UpLeftBy(1)
}

func (m *Matrix[T]) UpLeftBy(x int) *T {
	if !m.IsIn(m.CurrRow-x, m.CurrCol-x) {
		return nil
	}
	return &m.Matrix[m.CurrRow-x][m.CurrCol-x]
}

func (m *Matrix[T]) UpRight() *T {
	return m.UpRightBy(1)
}

func (m *Matrix[T]) UpRightBy(x int) *T {
	if !m.IsIn(m.CurrRow-x, m.CurrCol+x) {
		return nil
	}
	return &m.Matrix[m.CurrRow-x][m.CurrCol+x]
}

func (m *Matrix[T]) DownLeft() *T {
	return m.DownLeftBy(1)
}

func (m *Matrix[T]) DownLeftBy(x int) *T {
	if !m.IsIn(m.CurrRow+x, m.CurrCol-x) {
		return nil
	}
	return &m.Matrix[m.CurrRow+x][m.CurrCol-x]
}

func (m *Matrix[T]) DownRight() *T {
	return m.DownRightBy(1)
}

func (m *Matrix[T]) DownRightBy(x int) *T {
	if !m.IsIn(m.CurrRow+x, m.CurrCol+x) {
		return nil
	}
	return &m.Matrix[m.CurrRow+x][m.CurrCol+x]
}

func (m *Matrix[T]) Print() {
	for r := 0; r < len(m.Matrix); r++ {
		for c := 0; c < len(m.Matrix[r]); c++ {
			print(m.Matrix[r][c])
		}
		println()
	}
}
