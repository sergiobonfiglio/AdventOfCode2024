package main

import (
	"fmt"
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
			input: `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`,
			want: 7,
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
			input: `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`,
			want: "co,de,ka,ta",
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

func Test_subsets(t *testing.T) {

	tests := []struct {
		nodes   []string
		minSize int
		want    [][]string
	}{
		{
			nodes:   []string{"a", "b"},
			minSize: 2,
			want: [][]string{
				{"a", "b"},
			},
		},
		{
			nodes:   []string{"a", "b", "c"},
			minSize: 2,
			want: [][]string{
				{"a", "b"},
				{"b", "c"},
				{"a", "c"},
			},
		},
		{
			nodes:   []string{"a", "b", "c", "d"},
			minSize: 2,
			want: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"a", "d"},
				{"b", "c"},
				{"b", "d"},
				{"c", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("subsets(%v, %d)", tt.nodes, tt.minSize), func(t *testing.T) {
			got := subsets(tt.nodes, tt.minSize)
			if len(got) != len(tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}

			for _, want := range tt.want {
				slices.Sort(want)
				contain := false
				for _, got := range got {
					slices.Sort(got)
					if slices.Equal(got, want) {
						contain = true
						break
					}
				}
				if !contain {
					t.Errorf("subsets() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
