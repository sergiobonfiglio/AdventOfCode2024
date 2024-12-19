package main

import (
	"AdventOfCode2024/utils"
	"strings"
)

func parse1(input string) (*utils.Matrix[string], *utils.Cell, []string) {
	parts := strings.Split(input, "\n\n")

	grid, startPos := parseGrid(parts[0])

	moves := parseMoves(parts[1])

	return grid, startPos, moves
}

func parse2(input string) (*utils.Matrix[string], *utils.Cell, []string) {
	parts := strings.Split(input, "\n\n")

	gridStr := parts[0]

	gridStr = strings.NewReplacer(
		"#", "##",
		"O", "[]",
		".", "..",
		"@", "@.",
	).
		Replace(gridStr)

	grid, startPos := parseGrid(gridStr)

	moves := parseMoves(parts[1])

	return grid, startPos, moves
}

func parseMoves(movesStr string) []string {
	var moves []string
	for _, r := range movesStr {
		if r != '\n' {
			moves = append(moves, string(r))
		}
	}
	return moves
}

func parseGrid(gridStr string) (*utils.Matrix[string], *utils.Cell) {
	grid := utils.NewMatrixFromLinesStr(gridStr)
	var startPos *utils.Cell
	for cell, ok := grid.NextCell(); ok; cell, ok = grid.NextCell() {
		if *grid.GetAtCell(cell) == "@" {
			startPos = cell
			break
		}
	}

	if startPos == nil {
		panic("No start position found")
	}

	grid.Set(startPos.R, startPos.C)
	return grid, startPos
}

func part1(input string) any {

	grid, _, moves := parse1(input)

	for _, move := range moves {
		moveRobot(grid, move)
	}

	grid.Reset()
	totGps := 0
	for cell, ok := grid.NextCell(); ok; cell, ok = grid.NextCell() {
		if *grid.GetAtCell(cell) == "O" {
			gps := 100*cell.R + cell.C
			totGps += gps
		}
	}

	return totGps
}

func part2(input string) any {
	grid, _, moves := parse2(input)

	for _, move := range moves {
		moveRobot2(grid, move)
	}

	grid.Reset()
	totGps := 0
	for cell, ok := grid.NextCell(); ok; cell, ok = grid.NextCell() {
		if *grid.GetAtCell(cell) == "[" {
			gps := 100*cell.R + cell.C
			totGps += gps
		}
	}

	return totGps

}

func tryMove(grid *utils.Matrix[string], start, prev, curr *utils.Cell, dir string) bool {

	if *start == *curr {
		return true
	}

	val := *grid.GetAtCell(curr)
	if val == "#" {
		return false
	}

	if val == "." {
		grid.Swap(curr, prev)
		if *prev == *start {
			grid.Set(curr.R, curr.C)
		}

		return true
	}

	if val == "O" {
		next := curr.Dir(dir)
		move := tryMove(grid, start, curr, &next, dir)
		if move {
			grid.Swap(prev, curr)
			grid.Set(curr.R, curr.C)
		}
		return move
	}

	panic("Unknown cell " + string(val))
}

func moveRobot(grid *utils.Matrix[string], move string) {

	pos := grid.CurrCell()

	next := pos.Dir(move)
	tryMove(grid, pos, pos, &next, move)
}

func moveRobot2(grid *utils.Matrix[string], move string) {

	pos := grid.CurrCell()

	next := pos.Dir(move)

	tryMove2(grid, pos, pos, &next, move, false, false)

	if pos.DistManhattan(grid.CurrCell()) > 1 {
		panic("Moved more than 1 cell")
	}

}

func tryMove2(grid *utils.Matrix[string], start, prev, curr *utils.Cell, dir string, dryrun bool, backtrack bool) bool {

	if *start == *curr {
		return true
	}

	val := *grid.GetAtCell(curr)
	if val == "#" {
		return false
	}

	if val == "." {
		if !dryrun || backtrack {
			grid.Swap(curr, prev)
			if *prev == *start {
				grid.Set(curr.R, curr.C)
			}
		}
		return true
	}

	if dir == "<" {
		if val == "[" {
			next := curr.Dir(dir)
			move := tryMove2(grid, start, curr, &next, dir, false, backtrack)
			if move {
				grid.Swap(prev, curr)
			}
			return move
		} else if val == "]" {
			next := curr.Dir(dir)
			move := tryMove2(grid, start, curr, &next, dir, false, backtrack)
			if move {
				grid.Swap(prev, curr)
				grid.Set(curr.R, curr.C)
			}
			return move
		}
	}

	if dir == ">" {
		if val == "[" {
			next := curr.Dir(dir)
			move := tryMove2(grid, start, curr, &next, dir, false, backtrack)
			if move {
				grid.Swap(prev, curr)
				grid.Set(curr.R, curr.C)
			}
			return move

		} else if val == "]" {
			next := curr.Dir(dir)
			move := tryMove2(grid, start, curr, &next, dir, false, backtrack)
			if move {
				grid.Swap(prev, curr)
			}
			return move
		}
	}

	if (val == "[" || val == "]") && (dir == "^" || dir == "v") {
		next := curr.Dir(dir)
		moveA := tryMove2(grid, start, curr, &next, dir, true, backtrack)

		moveB := false

		if moveA {
			otherDir := ">"
			if val == "]" {
				otherDir = "<"
			}
			currB := curr.Dir(otherDir)
			nextB := currB.Dir(dir)
			moveB = tryMove2(grid, start, &currB, &nextB, dir, true, backtrack)
			if moveB {
				if backtrack {
					grid.Swap(prev, curr)
				}
				if *prev == *start {
					if !backtrack {
						tryMove2(grid, start, curr, &next, dir, false, true)
						tryMove2(grid, start, &currB, &nextB, dir, false, true)
						grid.Swap(prev, curr)
						grid.Set(curr.R, curr.C)
					}
				}
			}
		}

		return moveA && moveB
	}

	panic("Unknown cell " + string(val))
}
