package utils

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {

	cmpFromTarget := func(target int) func(i int) int {
		return func(i int) int {
			return i - target
		}
	}

	type args struct {
		start  int
		end    int
		target int
		//cmp   func(int) int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				start:  0,
				end:    11,
				target: 5,
			},
			want: 5,
		},
		{
			args: args{
				start:  0,
				end:    10,
				target: 5,
			},
			want: 5,
		},
		{
			args: args{
				start:  0,
				end:    10,
				target: -1,
			},
			want: 0,
		},
		{
			args: args{
				start:  0,
				end:    10,
				target: 0,
			},
			want: 0,
		},
		{
			args: args{
				start:  0,
				end:    10,
				target: 11,
			},
			want: 10,
		},
	}
	for _, tt := range tests {

		t.Run(fmt.Sprintf("[%d...%d]", tt.args.start, tt.args.end), func(t *testing.T) {
			if got := BinarySearch(tt.args.start, tt.args.end, cmpFromTarget(tt.args.target)); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
