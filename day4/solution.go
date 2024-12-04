package main

import (
	"AdventOfCode2024/utils"
	"strings"
)

// 2649 too high
func part1(input string) any {
	matrix := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, []rune(line))
	}

	rest := []rune{'M', 'A', 'S'}
	numMatches := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			//right
			if matrix[r][c] == 'X' && c+3 < len(matrix[0]) {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r][c+i+1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//left
			if matrix[r][c] == 'X' && c-3 >= 0 {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r][c-i-1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//down
			if matrix[r][c] == 'X' && r+3 < len(matrix) {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r+i+1][c] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//up
			if matrix[r][c] == 'X' && r-3 >= 0 {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r-i-1][c] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//diag \ down-right
			if matrix[r][c] == 'X' && r+3 < len(matrix) && c+3 < len(matrix[0]) {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r+i+1][c+i+1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//diag \ up-left
			if matrix[r][c] == 'X' && r-3 >= 0 && c-3 >= 0 {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r-i-1][c-i-1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//diag / down-left
			if matrix[r][c] == 'X' && r+3 < len(matrix) && c-3 >= 0 {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r+i+1][c-i-1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

			//diag / up-right
			if matrix[r][c] == 'X' && r-3 >= 0 && c+3 < len(matrix[0]) {
				match := true
				for i := 0; match && i < len(rest); i++ {
					match = matrix[r-i-1][c+i+1] == rest[i]
				}
				if match {
					numMatches++
				}
			}

		}
	}

	return numMatches
}

func part2(input string) any {
	matrix := utils.NewMatrixFromLines(input)

	numMatches := 0

	for val, ok := matrix.Next(); ok; val, ok = matrix.Next() {
		if val == 'A' {
			ul, ur, dl, dr := matrix.UpLeft(), matrix.UpRight(), matrix.DownLeft(), matrix.DownRight()

			if !utils.NotNil([]*rune{ul, ur, dl, dr}) {
				continue
			}

			w1 := string([]rune{*ul, *dr})
			w2 := string([]rune{*ur, *dl})

			w1Match := w1 == "MS" || w1 == "SM"
			w2Match := w2 == "MS" || w2 == "SM"

			if w1Match && w2Match {
				numMatches++
			}

		}
	}

	return numMatches
}

func part2_old(input string) any {
	matrix := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, []rune(line))
	}

	numMatches := 0

	for r := 1; r < len(matrix)-1; r++ {
		for c := 1; c < len(matrix[0])-1; c++ {
			if matrix[r][c] == 'A' {
				ul := matrix[r-1][c-1]
				ur := matrix[r-1][c+1]
				dl := matrix[r+1][c-1]
				dr := matrix[r+1][c+1]

				w1 := string([]rune{ul, dr})
				w2 := string([]rune{ur, dl})

				w1Match := w1 == "MS" || w1 == "SM"
				w2Match := w2 == "MS" || w2 == "SM"

				if w1Match && w2Match {
					numMatches++
				}

			}
		}
	}

	return numMatches
}
