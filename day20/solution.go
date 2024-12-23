package main

import (
	"AdventOfCode2024/utils"
)

func part1(input string) any {
	const maxPicoCheat = 2
	const minPico = 100

	maze, start, end := parseInput(input)
	cheats := findCheats(maze, start, end, maxPicoCheat)

	byPico := utils.GroupBy(cheats, func(cheat Cheat) int { return cheat.savedPico })

	tot := 0
	for pico, picoCheats := range byPico {
		if pico >= minPico {
			tot += len(picoCheats)
		}
	}

	return tot
}

func part2(input string) any {
	const maxPicoCheat = 20
	const minPico = 100

	maze, start, end := parseInput(input)
	cheats := findCheats(maze, start, end, maxPicoCheat)

	byPico := utils.GroupBy(cheats, func(cheat Cheat) int { return cheat.savedPico })

	tot := 0
	for pico, picoCheats := range byPico {
		if pico >= minPico {
			tot += len(picoCheats)
		}
	}

	return tot
}

func findCheats(maze *utils.Matrix[string], start, end *utils.Cell, max int) []Cheat {

	distStart, _ := mazeShortestPath(maze, start)
	distEnd, _ := mazeShortestPath(maze, end)

	shortestPath := distStart[*end]

	var cheats []Cheat

	for cell, _ := range distStart {

		neighbors := getCellsWithin(&cell, max)

		for _, neighbor := range neighbors {
			if _, ok := distEnd[*neighbor]; !ok {
				continue
			}
			mdist := cell.DistManhattan(neighbor)
			alt := distEnd[*neighbor] + mdist + distStart[cell]
			if alt < shortestPath {
				cheat := Cheat{
					start:     cell,
					end:       *neighbor,
					savedPico: shortestPath - alt,
				}
				cheats = append(cheats, cheat)
			}
		}
	}

	return cheats
}

func getCellsWithin(center *utils.Cell, maxDist int) []*utils.Cell {
	cells := make([]*utils.Cell, (2*maxDist)*(2*maxDist))
	i := 0
	for r := center.R - maxDist; r <= center.R+maxDist; r++ {
		for c := center.C - maxDist; c <= center.C+maxDist; c++ {
			if center.DistManhattan(utils.NewCell(r, c)) <= maxDist {
				cells[i] = utils.NewCell(r, c)
				i++
			}
		}
	}
	return cells[:i]
}

func parseInput(input string) (*utils.Matrix[string], *utils.Cell, *utils.Cell) {
	maze := utils.NewMatrixFromLinesStr(input)

	var start, end *utils.Cell
	for c, ok := maze.NextCell(); ok; c, ok = maze.NextCell() {
		val := *maze.GetAtCell(c)
		if val == "S" {
			start = c
		} else if val == "E" {
			end = c
		}
	}
	return maze, start, end
}

type Cheat struct {
	start, end utils.Cell
	savedPico  int
}

func mazeShortestPath(maze *utils.Matrix[string], start *utils.Cell) (map[utils.Cell]int, map[utils.Cell]*utils.Cell) {

	distances := map[utils.Cell]int{}
	prevs := map[utils.Cell]*utils.Cell{}
	unvisited := utils.NewMinHeap[utils.Cell]()

	itemsByCell := map[utils.Cell]*utils.Item[utils.Cell]{}

	distances[*start] = 0
	startItem := unvisited.HeapPush(*start, 0)
	itemsByCell[*start] = startItem

	for unvisited.Len() > 0 {
		curr, _ := unvisited.HeapPop()

		for _, n := range curr.NeighborsCross() {

			nVal := maze.GetAtCell(&n)
			if nVal != nil && (*nVal == "." || *nVal == "E") {
				cost := 1
				alt := distances[curr] + cost

				if bestDist, ok := distances[n]; !ok || alt < bestDist {
					distances[n] = alt
					prevs[n] = &curr

					if _, ok := itemsByCell[n]; !ok {
						item := unvisited.HeapPush(n, alt)
						itemsByCell[n] = item
					} else {
						unvisited.Update(itemsByCell[n], alt)
					}
				}
			}
		}
	}

	return distances, prevs
}
