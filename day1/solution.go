package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	var left []int
	var right []int
	for _, line := range strings.Split(input, "\n") {
		_ = line
		fields := strings.Fields(line)
		lNum, _ := strconv.Atoi(fields[0])
		rNum, _ := strconv.Atoi(fields[1])
		left = append(left, lNum)
		right = append(right, rNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	diff := 0
	for i := 0; i < len(left); i++ {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}

	return diff
}

func part2(input string) any {
	var left []int
	//var right []int
	rightOcc := map[int]int{}
	for _, line := range strings.Split(input, "\n") {
		_ = line
		fields := strings.Fields(line)
		lNum, _ := strconv.Atoi(fields[0])
		rNum, _ := strconv.Atoi(fields[1])

		left = append(left, lNum)
		rightOcc[rNum]++
	}

	wSum := 0
	for i := 0; i < len(left); i++ {
		wSum += left[i] * rightOcc[left[i]]
	}

	return wSum
}
