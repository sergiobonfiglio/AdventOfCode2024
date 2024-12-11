package main

import (
	"AdventOfCode2024/utils"
	"strconv"
)

func part1(input string) any {

	stones := utils.ToIntArray(input, " ")
	return applyRules(stones, 25, map[Mem]int{})

}

type Mem struct {
	Val  int
	Iter int
}

func applyRules(stones []int, iter int, memo map[Mem]int) int {
	sum := 0
	if iter == 0 {
		return len(stones)
	}

	for i := 0; i < len(stones); i++ {
		key := Mem{Val: stones[i], Iter: iter}

		if _, ok := memo[key]; !ok {
			next := applyRuleSingle(stones[i])
			subVal := applyRules(next, iter-1, memo)
			memo[key] = subVal
		}
		sum += memo[key]
	}

	return sum

}

func applyRuleSingle(x int) []int {

	if x == 0 {
		return []int{1}
	} else if str := strconv.Itoa(x); len(str)%2 == 0 {
		//If the stone is engraved with a number that has an even number of digits,
		//	it is replaced by two stones. The left half of the digits are engraved on the new left stone,
		//	and the right half of the digits are engraved on the new right stone. (The new numbers don't
		//	keep extra leading zeroes: 1000 would become stones 10 and 0.)
		first, _ := strconv.Atoi(str[:len(str)/2])
		second, _ := strconv.Atoi(str[len(str)/2:])
		return []int{first, second}
	} else {
		//If none of the other rules apply, the stone is replaced by a new stone;
		//	the old stone's number multiplied by 2024 is engraved on the new stone.
		return []int{x * 2024}
	}
}

func part2(input string) any {
	stones := utils.ToIntArray(input, " ")

	return applyRules(stones, 75, map[Mem]int{})
}
