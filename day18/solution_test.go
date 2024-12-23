package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		width  int
		height int
		after  int
		want   any
	}{
		{
			name: "example",
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`,
			width:  7,
			height: 7,
			after:  12,
			want:   22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1_params(tt.input, tt.width, tt.height, tt.after); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		width  int
		height int
		want   any
	}{
		{
			name: "example",
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`,
			width:  7,
			height: 7,
			want:   "6,1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2_params(tt.input, tt.width, tt.height); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
