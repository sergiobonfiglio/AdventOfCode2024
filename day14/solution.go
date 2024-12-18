package main

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
)

func part1(input string) any {
	return part1_wh(input, 101, 103, 100)
}

type Vector struct {
	P   *utils.Cell
	Vel *utils.Cell
}

func part1_wh(input string, width, height int, seconds int) any {
	var lines []*utils.Line
	for _, line := range strings.Split(input, "\n") {
		vec := parseLine(line)
		lines = append(lines, utils.NewLine(vec.P, utils.NewCell(vec.P.R+vec.Vel.R, vec.P.C+vec.Vel.C)))
	}

	var finalCells = afterSeconds(lines, width, height, seconds)

	q1, q2, q3, q4 := countQuadrants(finalCells, height, width)

	return q1 * q2 * q3 * q4
}

func countQuadrants(finalCells []*utils.Cell, height int, width int) (q1, q2, q3, q4 int) {
	for _, cell := range finalCells {
		if cell.R < height/2 && cell.C < width/2 {
			q1++
		} else if cell.R < height/2 && cell.C > width/2 {
			q2++
		} else if cell.R > height/2 && cell.C < width/2 {
			q3++
		} else if cell.R > height/2 && cell.C > width/2 {
			q4++
		}
	}
	return q1, q2, q3, q4
}

func groupByQuadrant(finalCells []*utils.Cell, height int, width int) (q1, q2, q3, q4, out []*utils.Cell) {
	for _, cell := range finalCells {
		if cell.R < height/2 && cell.C < width/2 {
			q1 = append(q1, cell)
		} else if cell.R < height/2 && cell.C > width/2 {
			q2 = append(q2, cell)
		} else if cell.R > height/2 && cell.C < width/2 {
			q3 = append(q3, cell)
		} else if cell.R > height/2 && cell.C > width/2 {
			q4 = append(q4, cell)
		} else {
			out = append(out, cell)
		}
	}
	return q1, q2, q3, q4, out
}

func afterSeconds(lines []*utils.Line, width, height, seconds int) []*utils.Cell {
	var finalCells []*utils.Cell
	for _, line := range lines {
		if line.IsVertical() {
			panic("vertical")
		}

		colVel := line.B.C - line.A.C // velocity in columns/sec

		finalCol := line.A.C + colVel*seconds
		finalRow := line.RowAtCol(finalCol)

		if finalRow == nil {
			panic("ops")
		}
		finalCell := utils.NewCell(*finalRow%height, finalCol%width)
		if finalCell.R < 0 {
			finalCell.R = height + finalCell.R
		}
		if finalCell.C < 0 {
			finalCell.C = width + finalCell.C
		}

		finalCells = append(finalCells, finalCell)
	}
	return finalCells
}

func printGrid(cells []*utils.Cell, width, height int) {
	grid := [][]string{}
	for i := 0; i < height; i++ {
		row := make([]string, width)
		for j := 0; j < width; j++ {
			row[j] = "."
		}
		grid = append(grid, row)
	}

	for _, cell := range cells {
		grid[cell.R][cell.C] = "#"
	}

	matrix := utils.NewMatrix[string](grid)

	matrix.Print()
}

func parseLine(line string) *Vector {
	split := strings.Split(line, " ")

	p := parseCell(split[0])
	v := parseCell(split[1])

	return &Vector{
		P:   p,
		Vel: v,
	}
}

func parseCell(str string) *utils.Cell {
	split := strings.Split(str, "=")
	numbers := strings.Split(split[1], ",")

	x, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	return &utils.Cell{
		R: y,
		C: x,
	}

}

func part2(input string) any {
	var lines []*utils.Line
	for _, line := range strings.Split(input, "\n") {
		vec := parseLine(line)
		lines = append(lines, utils.NewLine(vec.P, utils.NewCell(vec.P.R+vec.Vel.R, vec.P.C+vec.Vel.C)))
	}

	width := 101
	height := 103

	for seconds := 1; seconds < 10000; seconds++ {
		finalCells := afterSeconds(lines, width, height, seconds)

		if hasTree(finalCells) {
			//printGrid(finalCells, width, height)
			return seconds
		}

	}

	return nil
}

func hasTree(cells []*utils.Cell) bool {
	byRow := utils.GroupBy(cells, func(cell *utils.Cell) int { return cell.R })
	gt31 := 0

	for _, v := range byRow {
		if len(v) >= 31 {
			gt31++
		}
	}

	byCol := utils.GroupBy(cells, func(cell *utils.Cell) int { return cell.C })
	colGt31 := 0
	for _, v := range byCol {
		if len(v) >= 31 {
			colGt31++
		}
	}

	return gt31 >= 2 && colGt31 >= 2
}
