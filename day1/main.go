package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	var part int
	flag.IntVar(&part, "part", -1, "part number")
	flag.Parse()

	fmt.Printf("== Day %d ==\n", 1)
	if part == -1 {
		fmt.Printf("Part 1: %v\n", part1(input))
		fmt.Printf("Part 2: %v\n", part2(input))
	} else if part == 1 {
		fmt.Printf("Part 1: %v\n", part1(input))
	} else if part == 2 {
		fmt.Printf("Part 1: %v\n", part2(input))
	}
}

