package main

import (
	"AdventOfCode2024/utils"
	"strings"
)

func part1(input string) any {
	rules := map[int]map[int]bool{}
	i := 0
	lines := strings.Split(input, "\n")
	var line string
	for i, line = range lines {
		if line == "" {
			break
		}
		rule := utils.ToIntArray(line, "|")

		prev := rule[0]
		next := rule[1]

		if rules[prev] == nil {
			rules[prev] = map[int]bool{}
		}

		rules[prev][next] = true
	}

	sum := 0
	for i := i + 1; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		update := utils.ToIntArray(lines[i], ",")

		isOk := isUpdateOk(update, rules)
		if isOk {
			mid := update[len(update)/2]
			sum += mid
		}
	}

	return sum
}

func isUpdateOk(update []int, rules map[int]map[int]bool) bool {
	for ui := len(update) - 1; ui >= 0; ui-- {
		for j := 0; j < ui; j++ {
			if j != ui {
				if val, ok := rules[update[ui]][update[j]]; val && ok {
					return false
				}
			}
		}
	}
	return true
}

func part2(input string) any {
	rules := map[int]map[int]bool{}
	i := 0
	lines := strings.Split(input, "\n")
	var line string
	for i, line = range lines {
		if line == "" {
			break
		}
		rule := utils.ToIntArray(line, "|")

		prev := rule[0]
		next := rule[1]

		if rules[prev] == nil {
			rules[prev] = map[int]bool{}
		}

		rules[prev][next] = true
	}

	sum := 0
	for i := i + 1; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		update := utils.ToIntArray(lines[i], ",")

		isOk := isUpdateOk(update, rules)
		if !isOk {
			sorted := sort(update, rules)

			mid := update[len(sorted)/2]
			sum += mid
		}
	}

	return sum

}

func sort(update []int, rules map[int]map[int]bool) []int {
	for i := 0; i < len(update); i++ {
		free := findFreeNum(update[i:], rules)
		update[i], update[free+i] = update[free+i], update[i]
	}
	return update
}

func findFreeNum(update []int, rules map[int]map[int]bool) int {
	for i := 0; i < len(update); i++ {
		num := update[i]

		isFree := isFreeNum(num, update, rules)

		if isFree {
			return i
		}
	}
	panic("no free nums!")
}

func isFreeNum(num int, update []int, rules map[int]map[int]bool) bool {
	for i := 0; i < len(update); i++ {
		if rules[update[i]][num] {
			return false
		}
	}
	return true
}
