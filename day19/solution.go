package main

import "strings"

func part1(input string) any {

	sections := strings.Split(input, "\n\n")

	designs := map[string]bool{}
	for _, des := range strings.Split(sections[0], ",") {
		designs[strings.TrimSpace(des)] = true
	}

	possible := 0
	for _, target := range strings.Split(sections[1], "\n") {
		isPossible := checkPossible(designs, target)
		if isPossible {
			possible++
		}
	}

	return possible
}

func checkPossible(designs map[string]bool, target string) bool {

	if designs[target] {
		return true
	}

	for i := 1; i < len(target); i++ {
		if designs[target[:i]] && checkPossible(designs, target[i:]) {
			return true
		}
	}

	return false
}

func part2(input string) any {
	sections := strings.Split(input, "\n\n")

	designs := map[string]bool{}
	for _, des := range strings.Split(sections[0], ",") {
		designs[strings.TrimSpace(des)] = true
	}

	possible := 0
	memo := map[string]int{}
	for _, target := range strings.Split(sections[1], "\n") {
		possibilities := countPossible(designs, target, memo)
		possible += possibilities
	}

	return possible
}

func countPossible(designs map[string]bool, target string, memo map[string]int) int {

	if val, ok := memo[target]; ok {
		return val
	}

	possible := 0
	if designs[target] {
		possible += 1
	}

	for i := 1; i < len(target); i++ {
		if designs[target[:i]] {
			subPoss := countPossible(designs, target[i:], memo)
			possible += subPoss
		}
	}

	memo[target] = possible
	return possible
}
