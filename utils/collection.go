package utils

import (
	"cmp"
	"golang.org/x/exp/constraints"
	"sort"
)

func Map[T, R any](array []T, fn func(T) R) []R {
	var result []R
	for _, v := range array {
		result = append(result, fn(v))
	}
	return result
}

func FlatMap[T, R any](array []T, fn func(T) []R) []R {
	var result []R
	for _, v := range array {
		result = append(result, fn(v)...)
	}
	return result
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func FindFirst[T any](ss []T, test func(T) bool) (ret *T) {
	for _, s := range ss {
		if test(s) {
			ret := s
			return &ret
		}
	}
	return
}

func Reduce[T, R any](array []T, accumulator R, fn func(T, R) R) R {

	var result R
	for _, v := range array {
		result = fn(v, accumulator)
	}
	return result
}

// Intersection returns the intersection between the two arrays passed as arguments, nil if the intersection is empty.
func Intersection[T comparable](array1, array2 []T) []T {
	var intersection []T

	lookupMap := make(map[T]int)
	for _, it := range array1 {
		lookupMap[it] += 1
	}

	for _, it := range array2 {
		if lookupMap[it] > 0 {
			intersection = append(intersection, it)
			lookupMap[it]--
		}
	}

	return intersection
}

func ToDictionary[K comparable, V any](array []V, keyFn func(x V) K) map[K]V {
	dict := make(map[K]V)
	for _, v := range array {
		key := keyFn(v)
		dict[key] = v
	}
	return dict
}

func GroupBy[K comparable, V any](array []V, keyFn func(x V) K) map[K][]V {
	dict := make(map[K][]V)
	for _, v := range array {
		key := keyFn(v)
		dict[key] = append(dict[key], v)
	}
	return dict
}

func SortedKeys[K cmp.Ordered, V any](x map[K]V) []K {
	keys := make([]K, 0, len(x))
	for k := range x {
		keys = append(keys, k)
	}
	sort.Slice(keys, cmp.Less)
	return keys
}

func SafeSum[T constraints.Integer | constraints.Float](numbers ...*T) T {
	var tot T
	for _, x := range numbers {
		if x != nil {
			tot += *x
		}
	}
	return tot
}

func Empty[T any](slice []T) bool {
	return slice == nil || len(slice) == 0
}

func NotEmpty[T any](slice []T) bool {
	return !Empty(slice)
}

func NotNil[T any](slice []*T) bool {
	if slice == nil {
		return false
	}
	for _, s := range slice {
		if s == nil {
			return false
		}
	}
	return true
}
