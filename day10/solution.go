package main

import (
	"AdventOfCode2024/utils"
	"maps"
)

func part1(input string) any {

	matrix := utils.NewIntMatrixFromLines(input)

	var heads []*utils.Cell
	for val, ok := matrix.Next(); ok; val, ok = matrix.Next() {
		if val == 0 {
			heads = append(heads, utils.NewCell(matrix.CurrRow, matrix.CurrCol))
		}
	}

	scoreSum := 0
	for _, head := range heads {
		scoreSum += score(matrix, head, map[utils.Cell]bool{})
	}

	return scoreSum
}

func score(matrix *utils.Matrix[int], head *utils.Cell, visited map[utils.Cell]bool) int {

	neighbors := []utils.Cell{head.Up(1), head.Right(1), head.Down(1), head.Left(1)}

	if visited[*head] {
		return 0
	}

	visited[*head] = true
	totScore := 0
	headVal := matrix.GetAtCell(head)
	if *headVal == 9 {
		return 1
	}
	for _, n := range neighbors {
		if visited[n] {
			continue
		}

		val := matrix.GetAtCell(&n)

		if val != nil && *val-*headVal == 1 {
			//we can go there
			totScore += score(matrix, &n, visited)
		}
	}

	return totScore
}

func score2(matrix *utils.Matrix[int], head *utils.Cell, visited map[utils.Cell]bool) int {

	neighbors := []utils.Cell{head.Up(1), head.Right(1), head.Down(1), head.Left(1)}

	if visited[*head] {
		return 0
	}

	visited[*head] = true
	totScore := 0
	headVal := matrix.GetAtCell(head)
	if *headVal == 9 {
		return 1
	}
	for _, n := range neighbors {
		if visited[n] {
			continue
		}

		val := matrix.GetAtCell(&n)

		if val != nil && *val-*headVal == 1 {
			//we can go there
			visitedChild := map[utils.Cell]bool{}
			maps.Copy(visitedChild, visited)
			totScore += score2(matrix, &n, visitedChild)
		}
	}

	return totScore
}

func part2(input string) any {
	matrix := utils.NewIntMatrixFromLines(input)

	var heads []*utils.Cell
	for val, ok := matrix.Next(); ok; val, ok = matrix.Next() {
		if val == 0 {
			heads = append(heads, utils.NewCell(matrix.CurrRow, matrix.CurrCol))
		}
	}

	scoreSum := 0
	for _, head := range heads {
		scoreSum += score2(matrix, head, map[utils.Cell]bool{})
	}

	return scoreSum

}
