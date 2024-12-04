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
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			want: 18,
		},
		{
			name:  "right",
			input: `XMAS`,
			want:  1,
		},
		{
			name:  "left",
			input: `SAMX`,
			want:  1,
		},
		{
			name: "down",
			input: `X
M
A
S`,
			want: 1,
		},
		{
			name: "up",
			input: `S
A
M
X`,
			want: 1,
		},
		{
			name: "down-right",
			input: `X___
_M__
__A_
___S`,
			want: 1,
		},
		{
			name: "up-left",
			input: `S___
_A__
__M_
___X`,
			want: 1,
		},

		{
			name: "down-left",
			input: `___X
__M_
_A__
S___`,
			want: 1,
		},
		{
			name: "up-right",
			input: `S__S
__A_
_M__
X___`,
			want: 1,
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
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			want: 9,
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
