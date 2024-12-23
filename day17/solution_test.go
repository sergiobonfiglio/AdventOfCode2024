package main

import (
	"fmt"
	"reflect"
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
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			want: "4,6,3,5,6,3,5,2,1,0",
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
			input: `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`,
			want: int64(117440),
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

func Test_intsToBits(t *testing.T) {
	tests := []struct {
		args int
		want []int
	}{
		{0, []int{0}},
		{1, []int{1}},
		{2, []int{2}},
		{7, []int{7}},
		{8, []int{1, 0}},
		{9, []int{1, 1}},
		{63, []int{7, 7}},
		{64, []int{1, 0, 0}},
		{273, []int{4, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("intsToBits(%d)=%v", tt.args, tt.want), func(t *testing.T) {
			if got := intsToBits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intsToBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitsToInt(t *testing.T) {

	tests := []struct {
		bits []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 2},
		{[]int{7}, 7},
		{[]int{1, 0}, 8},
		{[]int{1, 1}, 9},
		{[]int{7, 7}, 63},
		{[]int{1, 0, 0}, 64},
		{[]int{4, 2, 1}, 273},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("bitsToInt(%v)=%d", tt.bits, tt.want), func(t *testing.T) {
			if got := bitsToInt(tt.bits); got != tt.want {
				t.Errorf("bitsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
