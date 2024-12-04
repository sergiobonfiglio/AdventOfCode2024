package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStepIter(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{start: 1, end: 10, step: 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			args: args{start: 1, end: 10, step: 2},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			args: args{start: 1, end: 10, step: 3},
			want: []int{1, 4, 7},
		},
		{
			args: args{start: 1, end: 1, step: 1},
			want: []int{},
		},
		{
			args: args{start: 1, end: 2, step: 1},
			want: []int{1},
		},
		{
			args: args{start: 5, end: 1, step: -1},
			want: []int{5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("NewStepIter(start %d, end %d, step %d)",
			tt.args.start, tt.args.end, tt.args.step), func(t *testing.T) {

			i := 0
			got := []int{}
			r := NewStepIter(tt.args.start, tt.args.end, tt.args.step)
			for num, ok := r.Next(); ok; num, ok = r.Next() {
				got = append(got, num)
				i++
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStepIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrixIter(t *testing.T) {

	type testCase[T interface{ Number | rune | string }] struct {
		matrix [][]T
		want   []T
	}
	tests := []testCase[int]{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v", tt.matrix), func(t *testing.T) {
			iter := NewMatrixIter(tt.matrix)
			got := []int{}
			for num, ok := iter.Next(); ok; num, ok = iter.Next() {
				got = append(got, num)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrixIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
