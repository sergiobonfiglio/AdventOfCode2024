package main

import (
	"regexp"
	"strconv"
)

func part1(input string) any {
	sum := 0

	regex := regexp.MustCompile(`mul\(([0-9+]+),([0-9]+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		g1, _ := strconv.Atoi(match[1])
		g2, _ := strconv.Atoi(match[2])
		sum += g1 * g2
	}

	return sum
}

func part2(input string) any {
	regexDont := regexp.MustCompile(`don't\(\)`)
	regexDo := regexp.MustCompile(`do\(\)`)
	regexMul := regexp.MustCompile(`mul\(([0-9+]+),([0-9]+)\)`)
	sum := 0

	left := 0
	right := len(input)
	for left < len(input) {
		if dontLoc := regexDont.FindIndex([]byte(input[left:])); dontLoc != nil {
			right = dontLoc[1] + left
		}

		matches := regexMul.FindAllStringSubmatch(input[left:right], -1)

		for _, match := range matches {
			g1, _ := strconv.Atoi(match[1])
			g2, _ := strconv.Atoi(match[2])
			sum += g1 * g2
		}

		left = len(input)
		if doLoc := regexDo.FindIndex([]byte(input[right:])); doLoc != nil {
			left = doLoc[1] + right
			right = len(input)
		}
	}

	return sum
}
