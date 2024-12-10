package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name: "example",
			input: `1190119
1111198
1112117
6543456
7651987
8761111
9871111`,
			want: 4,
		},
		{
			name: "example",
			input: `1011911
2111811
3111711
4567654
1118113
1119112
1111101`,
			want: 3,
		},
		{
			name:  "example",
			input: `0123456789`,
			want:  1,
		},
		{
			name: "example",
			input: `8888888888888888888
0123456789876543210
8888888888888888881
8808888888898765432`,
			want: 3,
		},
		{
			name: "example",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`,
			want: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name: "example",
			input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 81,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
