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
			input: `AAAA
BBCD
BBCC
EEEC`,
			want: 140,
		},
		{
			name: "example",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 772,
		},
		{
			name: "example",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1930,
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
			input: `AAAA
BBCD
BBCC
EEEC`,
			want: 80,
		},
		{
			name: "example",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 436,
		}, {
			name: "example",
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236,
		}, {
			name: "example",
			input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			want: 368,
		},
		{
			name: "example",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1206,
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
