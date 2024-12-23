package main

import (
	"AdventOfCode2024/utils"
	"testing"
)

type numCheat struct {
	num  int
	time int
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []numCheat
	}{
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`,
			want: []numCheat{
				{num: 14, time: 2},
				{num: 14, time: 4},
				{num: 2, time: 6},
				{num: 4, time: 8},
				{num: 2, time: 10},
				{num: 3, time: 12},
				{num: 1, time: 20},
				{num: 1, time: 36},
				{num: 1, time: 38},
				{num: 1, time: 40},
				{num: 1, time: 64},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze, start, end := parseInput(tt.input)
			cheats := findCheats(maze, start, end, 2)
			byPico := utils.GroupBy(cheats, func(cheat Cheat) int { return cheat.savedPico })

			if len(byPico) != len(tt.want) {
				t.Errorf("part1() = %v, want %v", len(byPico), len(tt.want))
			}
			for _, c := range tt.want {
				if got, ok := byPico[c.time]; !ok || len(got) != c.num {
					t.Errorf("part1() = %v, want %v", got, c.num)
				}
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []numCheat
	}{
		{
			name: "example",
			input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`,
			want: []numCheat{
				{num: 32, time: 50},
				{num: 31, time: 52},
				{num: 29, time: 54},
				{num: 39, time: 56},
				{num: 25, time: 58},
				{num: 23, time: 60},
				{num: 20, time: 62},
				{num: 19, time: 64},
				{num: 12, time: 66},
				{num: 14, time: 68},
				{num: 12, time: 70},
				{num: 22, time: 72},
				{num: 4, time: 74},
				{num: 3, time: 76},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze, start, end := parseInput(tt.input)
			cheats := findCheats(maze, start, end, 20)
			byPico := utils.GroupBy(cheats, func(cheat Cheat) int { return cheat.savedPico })

			for _, c := range tt.want {
				if got, ok := byPico[c.time]; !ok || len(got) != c.num {
					t.Errorf("part2() = %v, want %v", got, c.num)
				}
			}
		})
	}
}
