package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{

		{
			name: "example",
			input: `029A
`,
			want: 1972,
		},
		{
			name: "example",
			input: `379A
`,
			want: 24256,
		},

		{
			name: "example",
			input: `029A
980A
179A
456A
379A
`,
			want: 126384,
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
			input: `029A
`,
			want: 1972,
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

func TestChanger_generateCombinations(t *testing.T) {

	tests := []struct {
		seq  string
		want []string
	}{
		{
			seq:  "v<A",
			want: []string{"<vA", "v<A"},
		},
		{
			seq:  "v<Av<A",
			want: []string{"v<Av<A", "<vAv<A", "v<A<vA", "<vA<vA"},
		},
		{
			seq:  "<<A",
			want: []string{"<<A"},
		},
		{
			seq:  "<<A>>A^AvAAA",
			want: []string{"<<A>>A^AvAAA"},
		},
		{
			seq:  ">>^^A",
			want: []string{"^^>>A", ">>^^A", ">^^>A"},
		},
		{
			seq:  "<v<A",
			want: []string{"v<<A", "<v<A", "<<vA"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Next(%s)", tt.seq), func(t *testing.T) {
			var gotSlice []string
			gotSlice = generateCombinations(tt.seq)
			slices.Sort(gotSlice)
			slices.Sort(tt.want)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("Next() got = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
