package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
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
	var prof bool
	flag.BoolVar(&prof, "prof", false, "enable profiling")
	flag.Parse()

	if prof {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Printf("== Day %d ==\n", 14)
	if part == 1 || part == -1 {
		sol, elapsed := getTime(part1)
		fmt.Printf("Part 1: %v [%s]\n", sol, elapsed)
	}
	if part == 2 || part == -1 {
		sol, elapsed := getTime(part2)
		fmt.Printf("Part 2: %v [%s]\n", sol, elapsed)
	}
}

func getTime(fn func(input string) any) (any, time.Duration) {
	start := time.Now()
	sol1 := fn(input)
	elapsed := time.Since(start)
	return sol1, elapsed
}
