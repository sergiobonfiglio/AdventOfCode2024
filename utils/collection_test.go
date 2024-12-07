package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args[T comparable] struct {
		array1 []T
		array2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty intersection",
			args: args[int]{
				array1: []int{1, 2},
				array2: []int{3, 4},
			},
			want: nil,
		},
		{
			name: "full intersection",
			args: args[int]{
				array1: []int{1, 2},
				array2: []int{1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "partial intersection",
			args: args[int]{
				array1: []int{1, 2},
				array2: []int{2, 3},
			},
			want: []int{2},
		},
		{
			name: "one empty set",
			args: args[int]{
				array1: []int{},
				array2: []int{1, 2},
			},
			want: nil,
		},
		{
			name: "both empty",
			args: args[int]{
				array1: []int{},
				array2: []int{},
			},
			want: nil,
		},
		{
			name: "duplicates in both",
			args: args[int]{
				array1: []int{1, 2, 2, 3, 4},
				array2: []int{2, 2, 5, 5},
			},
			want: []int{2, 2},
		},
		{
			name: "duplicates in one",
			args: args[int]{
				array1: []int{1, 2, 2, 3, 4},
				array2: []int{2, 5},
			},
			want: []int{2},
		},
		{
			name: "different cardinality of duplicates",
			args: args[int]{
				array1: []int{1, 2, 2, 2, 3, 4},
				array2: []int{2, 2, 5},
			},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})

		// inverting args should not change the result
		t.Run(tt.name+" (inverse)", func(t *testing.T) {
			if got := Intersection(tt.args.array2, tt.args.array1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSum(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		numbers []*T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}

	one := 1
	two := 2
	three := 3

	tests := []testCase[int]{
		{
			name: "all nil",
			args: args[int]{[]*int{nil, nil, nil}},
			want: 0,
		},
		{
			name: "all numbers",
			args: args[int]{[]*int{nil, &one, nil}},
			want: 1,
		},
		{
			name: "mixed",
			args: args[int]{[]*int{&one, &two, &three}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeSum(tt.args.numbers...); got != tt.want {
				t.Errorf("SafeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type args[K comparable, V any] struct {
		array []V
		keyFn func(x V) K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K][]V
	}
	tests := []testCase[string, int]{
		{
			name: "GroupBy",
			args: args[string, int]{
				array: []int{1, 2, 2, 1, 3, 1},
				keyFn: func(x int) string {
					return fmt.Sprintf("%d", x)
				},
			},
			want: map[string][]int{
				"1": {1, 1, 1},
				"2": {2, 2},
				"3": {3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.array, tt.args.keyFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
