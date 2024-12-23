package utils

import (
	"strconv"
	"strings"
)

func ToIntegerArray[T int | int64](str string, sep string) []T {
	parts := strings.Split(strings.TrimSpace(str), sep)

	var res []T
	for _, part := range parts {
		if part != "" {
			n, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic("error")
			}
			res = append(res, T(n))
		}
	}
	return res
}

func ToIntArray(str string, sep string) []int {
	return ToIntegerArray[int](str, sep)
}

func ToInt64Array(str string, sep string) []int64 {
	return ToIntegerArray[int64](str, sep)
}

func FilterNil[T any](x []*T) []*T {
	var next []*T
	for _, el := range x {
		if el != nil {
			next = append(next, el)
		}
	}
	return next
}

// Coalesce returns the value of `v1` if it's not nil; otherwise, it returns `fallback`
func Coalesce[T any](v1 *T, fallback T) T {
	if v1 == nil {
		return fallback
	}
	return *v1
}

func Ptr[T any](x T) *T {
	return &x
}

func BinarySearch(start int, end int, compare func(int) int) int {
	for start < end {
		mid := start + (end-start)/2
		cmp := compare(mid)
		if cmp == 0 {
			return mid
		} else if cmp > 0 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
