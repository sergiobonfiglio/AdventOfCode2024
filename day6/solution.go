package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func part1(input string) any {

	rows := [][]int{}
	cols := [][]int{}

	maxCol := 0
	var start *utils.Cell
	for r, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		if len(cols) == 0 {
			cols = make([][]int, len(line))
			maxCol = len(line)
		}

		if start == nil {
			if stCol := strings.LastIndex(line, "^"); stCol != -1 {
				start = &utils.Cell{
					R: r,
					C: stCol,
				}
			}
		}

		regex := regexp.MustCompile(`#`)
		ix := regex.FindAllStringIndex(line, -1)

		rows = append(rows, []int{})
		for _, loc := range ix {
			rows[r] = append(rows[r], loc[0])
			cols[loc[0]] = append(cols[loc[0]], r)
		}
	}
	if start == nil {
		panic("no start")
	}

	visited := map[utils.Cell]bool{}

	var nextObstacle *utils.Cell
	dir := '^'
	nextObstacle = findNextObstacle(start, rows, cols, dir)
	for nextObstacle != nil {
		updateVisited(visited, start, nextObstacle, dir, -1)
		start, dir = nextStartDir(nextObstacle, dir)
		nextObstacle = findNextObstacle(start, rows, cols, dir)
	}

	//last leg
	updateVisited(visited, start, nil, dir, maxCol)

	return len(visited)
}

func findNextObstacle(start *utils.Cell, rows [][]int, cols [][]int, dir rune) *utils.Cell {
	if dir == '^' {
		if len(cols[start.C]) == 0 {
			return nil
		}
		for i := len(cols[start.C]) - 1; i >= 0; i-- {
			if cols[start.C][i] < start.R {
				return &utils.Cell{
					R: cols[start.C][i],
					C: start.C,
				}
			}
		}
		return nil

	} else if dir == '<' {
		if len(rows[start.R]) == 0 {
			return nil
		}
		for i := len(rows[start.R]) - 1; i >= 0; i-- {
			if rows[start.R][i] < start.C {
				return &utils.Cell{
					R: start.R,
					C: rows[start.R][i],
				}
			}
		}
		return nil

	} else if dir == '>' {

		if len(rows[start.R]) == 0 {
			return nil
		}
		for i := 0; i < len(rows[start.R]); i++ {
			if rows[start.R][i] > start.C {
				return &utils.Cell{
					R: start.R,
					C: rows[start.R][i],
				}
			}
		}
		return nil

	} else { // v
		if len(cols[start.C]) == 0 {
			return nil
		}
		for i := 0; i < len(cols[start.C]); i++ {
			if cols[start.C][i] > start.R {
				return &utils.Cell{
					R: cols[start.C][i],
					C: start.C,
				}
			}
		}
		return nil
	}
}

func nextStartDir(end *utils.Cell, dir rune) (*utils.Cell, rune) {
	if dir == '^' {
		return utils.Ptr(end.Down(1)), '>'
	} else if dir == '<' {
		return utils.Ptr(end.Right(1)), '^'
	} else if dir == '>' {
		return utils.Ptr(end.Left(1)), 'v'
	} else {
		return utils.Ptr(end.Up(1)), '<'
	}
}

func nextDir(dir rune) rune {
	if dir == '^' {
		return '>'
	} else if dir == '<' {
		return '^'
	} else if dir == '>' {
		return 'v'
	} else {
		return '<'
	}
}

func updateVisited(visited map[utils.Cell]bool, start, end *utils.Cell, dir rune, max int) {
	if end == nil {
		if dir == '<' || dir == '^' {
			end = &utils.Cell{
				R: -1,
				C: -1,
			}
		} else {
			end = &utils.Cell{
				R: max,
				C: max,
			}
		}
	}

	if dir == '^' {
		for i := start.R; i > end.R; i-- {
			visited[utils.Cell{R: i, C: start.C}] = true
		}
	} else if dir == '<' {
		for i := start.C; i > end.C; i-- {
			visited[utils.Cell{R: start.R, C: i}] = true
		}
	} else if dir == '>' {
		for i := start.C; i < end.C; i++ {
			visited[utils.Cell{R: start.R, C: i}] = true
		}
		dir = 'v'
	} else {
		for i := start.R; i < end.R; i++ {
			visited[utils.Cell{R: i, C: start.C}] = true
		}
	}
}

func part2(input string) any {
	rows := [][]int{}
	cols := [][]int{}

	maxCol := 0
	var start *utils.Cell
	for r, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		if len(cols) == 0 {
			cols = make([][]int, len(line))
			maxCol = len(line)
		}

		if start == nil {
			if stCol := strings.LastIndex(line, "^"); stCol != -1 {
				start = &utils.Cell{
					R: r,
					C: stCol,
				}
			}
		}

		regex := regexp.MustCompile(`#`)
		ix := regex.FindAllStringIndex(line, -1)

		rows = append(rows, []int{})
		for _, loc := range ix {
			rows[r] = append(rows[r], loc[0])
			cols[loc[0]] = append(cols[loc[0]], r)
		}
	}
	if start == nil {
		panic("no start")
	}

	origStart := &utils.Cell{R: start.R, C: start.C}

	visited := map[utils.Cell]bool{}

	var nextObstacle *utils.Cell
	dir := '^'
	nextObstacle = findNextObstacle(start, rows, cols, dir)
	for nextObstacle != nil {
		updateVisited(visited, start, nextObstacle, dir, -1)
		start, dir = nextStartDir(nextObstacle, dir)
		nextObstacle = findNextObstacle(start, rows, cols, dir)
	}

	//last leg
	updateVisited(visited, start, nil, dir, maxCol)

	//reset start
	start = origStart

	possibleLocs := map[utils.Cell]bool{}

	for cell, _ := range visited {
		if cell.R == start.R && cell.C == start.C {
			continue
		}
		// copy rows, cols
		simRows, simCols := addObstacle(&cell, rows, cols)

		// simulate
		if hasLoop(start, simRows, simCols) {
			possibleLocs[cell] = true
			continue
		}
	}

	return len(possibleLocs)
}

func hasLoop(oStart *utils.Cell, rows, cols [][]int) bool {

	start := &utils.Cell{R: oStart.R, C: oStart.C}

	hitMap := map[Hit]bool{}

	var nextObstacle *utils.Cell
	dir := '^'
	nextObstacle = findNextObstacle(start, rows, cols, dir)
	for nextObstacle != nil {

		hit := &Hit{Cell: *nextObstacle, Dir: dir}

		if _, present := hitMap[*hit]; present {
			return true
		}
		hitMap[*hit] = true

		start, dir = nextStartDir(nextObstacle, dir)
		nextObstacle = findNextObstacle(start, rows, cols, dir)
	}

	return false
}

func addObstacle(obst *utils.Cell, rows, cols [][]int) (newRows, newCols [][]int) {
	newRows = make([][]int, len(rows))
	for i, row := range rows {
		newRows[i] = make([]int, len(row))
		copy(newRows[i], row)
	}
	newCols = make([][]int, len(cols))
	for i, col := range cols {
		newCols[i] = make([]int, len(col))
		copy(newCols[i], col)
	}

	newRows[obst.R] = append(newRows[obst.R], obst.C)
	slices.Sort(newRows[obst.R])
	newCols[obst.C] = append(newCols[obst.C], obst.R)
	slices.Sort(newCols[obst.C])

	return newRows, newCols
}

type Hit struct {
	Cell utils.Cell
	Dir  rune
}

func (h Hit) String() string {
	return fmt.Sprintf("%v %v", h.Cell, string(h.Dir))
}
