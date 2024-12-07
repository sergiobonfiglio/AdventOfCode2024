package utils

type Iter[T any] interface {
	Next() (T, bool)
	Reset()
}

type StepIter[T Number] struct {
	current, end, step T
}

func (r *StepIter[T]) Reset() {
	r.current = 0
}

func NewStepIter[T Number](start, end, step T) Iter[T] {
	if step == 0 {
		panic("Step cannot be zero")
	}
	return &StepIter[T]{current: start, end: end, step: step}
}

func (r *StepIter[T]) Next() (T, bool) {
	if (r.step > 0 && r.current >= r.end) || (r.step < 0 && r.current <= r.end) {
		return *new(T), false
	}
	value := r.current
	r.current += r.step
	return value, true
}

type MatrixIter[T Number | rune | string] struct {
	matrix  [][]T
	currRow int
	currCol int
}

var _ Iter[int] = new(MatrixIter[int])

func NewMatrixIter[T Number | rune | string](matrix [][]T) *MatrixIter[T] {
	return &MatrixIter[T]{matrix: matrix, currRow: 0, currCol: -1}
}

func (m *MatrixIter[T]) Reset() {
	m.currCol = -1
	m.currRow = 0
}

func (m *MatrixIter[T]) Next() (T, bool) {

	nextC := (m.currCol + 1) % len(m.matrix[m.currRow])
	if nextC < m.currCol {
		if m.currRow+1 == len(m.matrix) {
			return *new(T), false
		}
		m.currRow = (m.currRow + 1) % len(m.matrix)
	}
	m.currCol = nextC
	return m.matrix[m.currRow][m.currCol], true
}
