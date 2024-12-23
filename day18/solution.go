package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strings"
)

func part1(input string) any {
	return part1_params(input, 71, 71, 1024)
}

func part2(input string) any {
	return part2_params(input, 71, 71)

}

func part2_params(input string, width int, height int) string {
	var falling []*utils.Cell
	//falled := map[utils.Cell]bool{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		xy := utils.ToIntArray(line, ",")
		x, y := xy[0], xy[1]

		cell := utils.NewCell(y, x)
		falling = append(falling, cell)
		//falled[*cell] = true
	}

	bounds := &Bounds{
		MinR: 0,
		MinC: 0,
		MaxR: height - 1,
		MaxC: width - 1,
	}

	start := *utils.NewCell(0, 0)
	end := utils.NewCell(height-1, width-1)

	sol := utils.BinarySearch(0, len(falling), func(t int) int {
		falled := falledAt(t, falling)
		distances, _ := mazeShortestPath(bounds, falled, start)
		if _, ok := distances[*end]; !ok {
			return 1
		}
		return -1
	})
	blocking := falling[sol-1]
	return fmt.Sprintf("%d,%d", blocking.C, blocking.R)
}

func falledAt(t int, falling []*utils.Cell) map[utils.Cell]bool {
	falled := map[utils.Cell]bool{}
	for i := 0; i < t; i++ {
		falled[*falling[i]] = true
	}
	return falled
}

func part1_params(input string, width int, height int, after int) int {

	falled := map[utils.Cell]bool{}
	for i, line := range strings.Split(input, "\n") {

		if i >= after {
			break
		}

		xy := utils.ToIntArray(line, ",")
		x, y := xy[0], xy[1]

		cell := utils.NewCell(y, x)
		falled[*cell] = true
	}

	if len(falled) != after {
		panic("not enough falling cells")
	}

	bounds := &Bounds{
		MinR: 0,
		MinC: 0,
		MaxR: height - 1,
		MaxC: width - 1,
	}

	distances, _ := mazeShortestPath(bounds, falled, *utils.NewCell(0, 0))

	end := utils.NewCell(height-1, width-1)
	if dist, ok := distances[*end]; !ok {
		panic("end not reached")
	} else {
		return dist
	}

}

type Bounds struct {
	MinR, MinC, MaxR, MaxC int
}

func (b *Bounds) Contains(c *utils.Cell) bool {
	return c.R >= b.MinR && c.R <= b.MaxR &&
		c.C >= b.MinC && c.C <= b.MaxC
}

func isValid(c *utils.Cell, falled map[utils.Cell]bool, bounds *Bounds) bool {
	return bounds.Contains(c) && !falled[*c]
}

func mazeShortestPath(bounds *Bounds, falled map[utils.Cell]bool, start utils.Cell) (map[utils.Cell]int, map[utils.Cell]*utils.Cell) {

	distances := map[utils.Cell]int{}
	prevs := map[utils.Cell]*utils.Cell{}
	unvisited := utils.NewMinHeap[utils.Cell]()

	itemsByCell := map[utils.Cell]*utils.Item[utils.Cell]{}

	distances[start] = 0
	startItem := unvisited.HeapPush(start, 0)
	itemsByCell[start] = startItem

	for unvisited.Len() > 0 {
		curr, _ := unvisited.HeapPop()

		for _, n := range curr.NeighborsCross() {

			valid := isValid(&n, falled, bounds)
			if valid {
				cost := 1
				alt := distances[curr] + cost

				if bestDist, ok := distances[n]; !ok || alt < bestDist {
					distances[n] = alt
					prevs[n] = &curr

					if _, ok := itemsByCell[n]; !ok {
						// add to unvisited if not already there
						item := unvisited.HeapPush(n, alt)
						itemsByCell[n] = item
					} else {
						// just update the distance if already there
						unvisited.Update(itemsByCell[n], alt)
					}
				}
			}
		}
	}

	return distances, prevs
}
