package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(input string) any {
	return solve(input, 2)
}

func part2(input string) any {
	return solve(input, 25)
}

func solve(input string, inBetweenRobots int) int {
	sum := 0
	for _, code := range strings.Split(input, "\n") {
		if code == "" {
			continue
		}
		shortestSeq := shortestSequence(code, inBetweenRobots)
		comp := complexity(code, shortestSeq)
		sum += comp
	}

	return sum
}

func perm(str string) map[string]bool {
	if len(str) == 1 {
		return map[string]bool{str: true}
	}
	if len(str) == 2 {
		return map[string]bool{
			str:                             true,
			string(str[1]) + string(str[0]): true}
	}

	p1 := perm(str[1:])
	tot := map[string]bool{}
	for p, _ := range p1 {
		tot[str[0:1]+p] = true
		tot[p+str[0:1]] = true
	}

	return tot
}

func generateCombinations(seq string) []string {

	firstA := strings.Index(seq, "A")

	if firstA == -1 || firstA == 0 {
		return []string{seq}
	}

	comb := []string{seq[:firstA+1]}
	if !isAllOneChar(seq[:firstA]) {
		perms := perm(seq[:firstA])
		comb = []string{}
		for p, _ := range perms {
			comb = append(comb, p+"A")
		}
	}

	tail := generateCombinations(seq[firstA+1:])
	var concatenated []string
	if tail != nil {
		for _, t := range tail {
			for _, c := range comb {
				concatenated = append(concatenated, c+t)
			}
		}
		return concatenated
	} else {
		for _, c := range comb {
			concatenated = append(concatenated, c+seq[firstA+1:])
		}
		return concatenated
	}

}

func isAllOneChar(chunk string) bool {
	r := rune(chunk[0])
	for i := 1; i < len(chunk); i++ {
		if rune(chunk[i]) != r {
			return false
		}
	}
	return true
}

func shortestSequence(code string, inBetweenRobots int) int {

	startButton := 'A'
	memo := map[string]int{}

	tot := 0
	for _, char := range code {
		numKeypadSeq := shortestNumericKeypad(string(char), startButton)
		shortest := math.MaxInt

		for _, currSequence := range generateCombinations(numKeypadSeq) {
			if !isValidNumKeySeq(currSequence, startButton) {
				continue
			}

			recTot := 0
			prev := 'A'
			for _, c := range currSequence {
				val := recursiveVal(prev, c, inBetweenRobots, memo)
				prev = c
				recTot += val
			}
			if recTot < shortest {
				shortest = recTot
			}
		}

		tot += shortest
		startButton = char
	}

	return tot
}

func recursiveVal(start, end rune, level int, memo map[string]int) int {
	if level == 0 {
		return 1
	}

	key := string([]rune{start, end}) + strconv.Itoa(level)
	if v, ok := memo[key]; ok {
		return v
	}

	var minVal *int
	for _, altSeq := range generateCombinations(dirKeySeq(start, end)) {
		if !isValidDirKeySeq(altSeq, start) {
			continue
		}

		val := 0
		prev := 'A'
		for _, c := range altSeq {
			val += recursiveVal(prev, c, level-1, memo)
			prev = c
		}

		if minVal == nil || val < *minVal {
			minVal = &val
		}
	}

	if minVal == nil {
		panic("no valid sequence")
	}

	memo[key] = *minVal

	return *minVal
}

func nextRC(r, c int, x rune) (int, int) {
	switch x {
	case '^':
		return r - 1, c
	case 'v':
		return r + 1, c
	case '<':
		return r, c - 1
	case '>':
		return r, c + 1
	default:
		panic(fmt.Sprintf("Unknown direction %c", x))
	}
}

func isValidNumKeySeq(sequence string, start rune) bool {
	currR, currC := numKeyCoords(start)

	for _, dir := range sequence {
		if dir == 'A' {
			continue
		}
		nextR, nextC := nextRC(currR, currC, dir)
		if nextR < 0 || nextR > 3 ||
			nextC < 0 || nextC > 2 ||
			(nextR == 3 && nextC == 0) {
			return false
		}
		currR, currC = nextR, nextC
	}
	return true
}
func isValidDirKeySeq(sequence string, start rune) bool {
	currR, currC := dirKeyCoords(start)

	for _, dir := range sequence {
		if dir == 'A' {
			continue
		}
		nextR, nextC := nextRC(currR, currC, dir)
		if nextR < 0 || nextR > 1 ||
			nextC < 0 || nextC > 2 ||
			(nextR == 0 && nextC == 0) {
			return false
		}
		currR, currC = nextR, nextC
	}
	return true
}

func shortestNumericKeypad(code string, current rune) string {
	seq := ""
	for _, c := range code {
		seq += numKeySeq(current, c)
		current = c
	}
	return seq
}

func dirKeySeq(start rune, end rune) string {

	r1, c1 := dirKeyCoords(start)
	r2, c2 := dirKeyCoords(end)

	seq := ""

	if r1 == 0 && c2 == 0 {
		// let's go down to avoid the gap
		seq += "v"
		r1 = 1
	}

	if c1 < c2 {
		seq += strings.Repeat(">", c2-c1)
	} else if c1 > c2 {
		seq += strings.Repeat("<", c1-c2)
	}

	if r1 < r2 {
		seq += strings.Repeat("v", r2-r1)
	} else if r1 > r2 {
		seq += strings.Repeat("^", r1-r2)
	}

	return seq + "A"
}
func numKeySeq(start rune, end rune) string {
	r1, c1 := numKeyCoords(start)
	r2, c2 := numKeyCoords(end)

	var seq string
	if r1 == 3 && c2 == 0 {
		// let's go up to avoid the gap
		seq += "^"
		r1 = 2
	}
	if r1 < r2 {
		seq += strings.Repeat("v", r2-r1)
	} else if r1 > r2 {
		seq += strings.Repeat("^", r1-r2)
	}
	if c1 < c2 {
		seq += strings.Repeat(">", c2-c1)
	} else if c1 > c2 {
		seq += strings.Repeat("<", c1-c2)
	}

	return seq + "A"
}

func numKeyCoords(button rune) (r int, c int) {
	switch button {

	case '7':
		return 0, 0
	case '8':
		return 0, 1
	case '9':
		return 0, 2
	case '4':
		return 1, 0
	case '5':
		return 1, 1
	case '6':
		return 1, 2
	case '1':
		return 2, 0
	case '2':
		return 2, 1
	case '3':
		return 2, 2
	case '0':
		return 3, 1
	case 'A':
		return 3, 2
	}
	panic("invalid button")
}

func dirKeyCoords(button rune) (r int, c int) {
	switch button {
	case '^':
		return 0, 1
	case 'A':
		return 0, 2
	case '<':
		return 1, 0
	case 'v':
		return 1, 1
	case '>':
		return 1, 2
	}
	panic("invalid button")
}

func complexity(code string, minSequence int) int {
	num, _ := strconv.Atoi(code[:len(code)-1])
	return num * minSequence
}
