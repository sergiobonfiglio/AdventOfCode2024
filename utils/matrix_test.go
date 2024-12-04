package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatrix_GetUpBy(t *testing.T) {

	type testCase[T any] struct {
		m       *Matrix[T]
		currRow int
		currCol int
		x       int
		want    []T
	}
	tests := []testCase[int]{
		{
			m: NewMatrix([][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}}),
			x:       2,
			currRow: 2,
			currCol: 0,
			want:    []int{1, 4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%d,%d]GetUp(%d)", tt.currRow, tt.currCol, tt.x), func(t *testing.T) {
			tt.m.Set(tt.currRow, tt.currCol)
			if got := tt.m.GetUpBy(tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
