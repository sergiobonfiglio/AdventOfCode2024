package utils

type Number interface {
	~int | ~float32 | ~float64
}

func GCD[T int | int64](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T int | int64](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM[T](result, integers[i])
	}

	return result
}

func Between[T Number](x, start, end T) bool {
	return x >= start && x < end
}
