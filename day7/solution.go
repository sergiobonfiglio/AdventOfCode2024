package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(input string) any {

	validSum := int64(0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		eq := strings.Split(line, ":")
		val, _ := strconv.ParseInt(eq[0], 10, 64)
		nums := utils.ToInt64Array(eq[1], " ")

		if canBeValid(val, nums) {
			validSum += val
		}
	}

	return validSum
}

func canBeValid(val int64, nums []int64) bool {

	if len(nums) == 1 {
		return nums[0] == val
	}

	if nums[0] > val {
		return false
	}

	wPlusValid := canBeValid(val, append([]int64{nums[0] + nums[1]}, nums[2:]...))
	if wPlusValid {
		return true
	}

	wMultValid := canBeValid(val, append([]int64{nums[0] * nums[1]}, nums[2:]...))
	if wMultValid {
		return true
	}

	return false
}

func part2(input string) any {
	validSum := int64(0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		eq := strings.Split(line, ":")
		val, _ := strconv.ParseInt(eq[0], 10, 64)
		nums := utils.ToInt64Array(eq[1], " ")

		if canBeValid2(val, nums) {
			validSum += val
		}
	}

	return validSum
}

func canBeValid2(val int64, nums []int64) bool {
	if len(nums) == 1 {
		return nums[0] == val
	}

	if nums[0] > val {
		return false
	}

	wPlusValid := canBeValid2(val, append([]int64{nums[0] + nums[1]}, nums[2:]...))
	if wPlusValid {
		return true
	}

	wMultValid := canBeValid2(val, append([]int64{nums[0] * nums[1]}, nums[2:]...))
	if wMultValid {
		return true
	}

	concNum, err := strconv.ParseInt(fmt.Sprintf("%d%d", nums[0], nums[1]), 10, 64)
	if err != nil {
		panic("help")
	}
	wConcatValid := canBeValid2(val, append([]int64{concNum}, nums[2:]...))
	if wConcatValid {
		return true
	}

	return false
}
