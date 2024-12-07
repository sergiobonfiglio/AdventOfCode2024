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
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			want: 41,
		},
		{
			name: "example1",
			input: `.....
..#..
..^..
.#...
....#`,
			want: 3,
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

//0123456789
//....#.....0
//.........#1
//..........2
//..#.......3
//.......#..4
//..........5
//.#..^.....6
//........#.7
//#.........8
//......#...9
//0123456789

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name: "example",
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			want: 6,
		},

		{
			name: "example1",
			input: `.....
..#..
..^..
.#...
...##`,
			want: 1,
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
