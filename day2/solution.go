package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) any {
	safe := 0
	for _, line := range strings.Split(input, "\n") {
		levelsStr := strings.Fields(line)
		var levels []int
		for _, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if levelsAreSafe(levels) {
			safe += 1
		}
	}

	return safe
}

func levelsAreSafe(levels []int) bool {

	first := levels[0]
	second := levels[1]
	isAsc := first < second
	diff := math.Abs(float64(second - first))
	isSafe := diff > 0 && diff < 4

	i := 2
	prev := second
	for isSafe && i < len(levels) {
		curr := levels[i]
		diff = math.Abs(float64(prev - curr))
		isSafe = diff > 0 && diff < 4
		isSafe = isSafe && ((isAsc && curr >= prev) || (!isAsc && curr <= prev))
		prev = curr
		i++
	}
	return isSafe
}

func part2(input string) any {
	safe := 0
	for _, line := range strings.Split(input, "\n") {
		levelsStr := strings.Fields(line)
		var levels []int
		for _, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if levelsAreSafeP2(levels) {
			safe += 1
		}
	}

	return safe
}

func levelsAreSafeP2(levels []int) bool {
	isSafe := levelsAreSafe(levels)

	if !isSafe {

		if levelsAreSafe(levels[1:]) {
			return true
		}

		for i := 1; i < len(levels)-1; i++ {
			dLev := []int{}
			dLev = append(dLev, levels[0:i]...)
			dLev = append(dLev, levels[i+1:]...)
			if levelsAreSafe(dLev) {
				return true
			}
		}
		if levelsAreSafe(levels[0 : len(levels)-1]) {
			return true
		}

	}
	return isSafe
}
