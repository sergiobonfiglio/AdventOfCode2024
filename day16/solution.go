package main

import (
	"AdventOfCode2024/utils"
	"maps"
	"math"
	"slices"
	"sync"
)

func part1(input string) any {

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
	if start == nil || end == nil {
		panic("start or end not found")
	}
	maze.Reset()

	//visited := &sync.Map{}
	//visited.Store(*start, 0)
	//
	//bestScore := solveMaze(maze, nil, start, 0, visited)
	//return bestScore

	dist, _ := mazeShortestPath(maze, start)
	return dist[*end]
}

func part2(input string) any {
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
	if start == nil || end == nil {
		panic("start or end not found")
	}
	maze.Reset()

	_, prevs := mazeShortestPaths(maze, start)

	var ends []DirCell
	for _, r := range []rune{'^', '>', 'v', '<'} {
		dirEnd := DirCell{
			dir:  r,
			Cell: *end,
		}
		ends = append(ends, dirEnd)
	}

	bpcells := map[utils.Cell]bool{}
	countCellsInShortestPaths(prevs, ends, bpcells)

	//for cell := range bpcells {
	//	maze.SetValAtCell(&cell, "O")
	//}
	//maze.Print()

	return len(bpcells)

	//visited := &sync.Map{}
	//visited.Store(DirCell{
	//	dir:  '>',
	//	Cell: *start,
	//}, 0)
	//_, _, bestPathCells := solveMaze2(maze, nil, start, 0, bestScore, map[utils.Cell]bool{*start: true}, visited)
	//return len(bestPathCells)
}

func countCellsInShortestPaths(prevs map[DirCell]map[DirCell]bool, targets []DirCell, visited map[utils.Cell]bool) {

	for _, target := range targets {
		visited[target.Cell] = true
		var sources []DirCell
		for source, _ := range prevs[target] {
			visited[source.Cell] = true
			sources = append(sources, source)
		}
		countCellsInShortestPaths(prevs, sources, visited)
	}

}

type DirCell struct {
	dir rune
	utils.Cell
}

func (c DirCell) rotate(n int) DirCell {
	poss := []rune{'^', '>', 'v', '<'}

	index := slices.Index(poss, c.dir)
	next := (index + n) % len(poss)
	if next < 0 {
		next = (len(poss) + next) % len(poss)
	}

	return DirCell{
		dir:  poss[next],
		Cell: c.Cell,
	}
}

func solveMaze2(
	maze *utils.Matrix[string],
	prev, current *utils.Cell,
	currCost int,
	maxCost int,
	bestPathCells map[utils.Cell]bool,
	visited *sync.Map,
) (int, int, map[utils.Cell]bool) {

	if *maze.GetAtCell(current) == "E" {
		bestPathCells[*current] = true
		return currCost, min(maxCost, currCost), bestPathCells
	}
	if currCost >= maxCost {
		return math.MaxInt32, maxCost, nil
	}

	currDir := currentDir(prev, current)

	next := current.NeighborsCross()

	results := make(chan *VisitResult, len(next))
	var wg sync.WaitGroup
	for _, n := range next {
		wg.Add(1)
		go func(n utils.Cell) {
			defer wg.Done()

			result := VisitCell(maze, current, &n, currDir, currCost, maxCost, maps.Clone(bestPathCells), visited)
			results <- result
		}(n)
	}
	wg.Wait()
	close(results)

	minCost := math.MaxInt32
	costByCell := map[utils.Cell]int{}
	bpsByCell := map[utils.Cell]map[utils.Cell]bool{}
	for r := range results {
		if r == nil {
			continue
		}
		costByCell[*r.Cell] = r.Cost
		bpsByCell[*r.Cell] = r.BestPathCells
		if r.Cost <= minCost {
			minCost = r.Cost
		}
	}

	bestPathCells[*current] = true
	for cell, i := range costByCell {
		if i <= minCost {
			for c := range bpsByCell[cell] {
				bestPathCells[c] = true
			}
		}
	}

	return minCost, maxCost, bestPathCells
}

type VisitResult struct {
	Cell          *utils.Cell
	Cost          int
	BestPathCells map[utils.Cell]bool
}

func VisitCell(
	maze *utils.Matrix[string],
	current *utils.Cell,
	next *utils.Cell,
	currDir rune,
	currCost int,
	maxCost int,
	bestPathCells map[utils.Cell]bool,
	visited *sync.Map,
) *VisitResult {
	nVal := maze.GetAtCell(next)
	if nVal != nil && (*nVal == "." || *nVal == "E") {

		nextDir := currentDir(current, next)
		isOpposite := isOppositeDir(currDir, nextDir)
		cost := 1
		if isOpposite {
			cost += 2000
		} else if currDir != nextDir {
			cost += 1000
		}

		dirCell := DirCell{dir: nextDir, Cell: *next}
		if storedCost, ok := visited.Load(dirCell); ok && storedCost.(int) < currCost+cost {
			return nil
		}
		visited.Store(dirCell, currCost+cost)

		for _, i := range []int{1, -1, 2} {
			rotated := dirCell.rotate(i)
			additionalCost := 1000 * int(math.Abs(float64(i)))
			if storedCost, ok := visited.Load(rotated); !ok || storedCost.(int) > currCost+cost+additionalCost {
				visited.Store(rotated, currCost+cost+additionalCost)
			}
		}

		totCost, _, bps := solveMaze2(maze, current, next, currCost+cost, maxCost, maps.Clone(bestPathCells), visited)

		return &VisitResult{
			Cell:          next,
			Cost:          totCost,
			BestPathCells: bps,
		}
	}
	return nil
}

func mazeShortestPaths(maze *utils.Matrix[string], start *utils.Cell) (map[DirCell]int, map[DirCell]map[DirCell]bool) {

	distances := map[DirCell]int{}
	prevs := map[DirCell]map[DirCell]bool{}
	unvisited := utils.NewMinHeap[DirCell]()

	itemsByCell := map[DirCell]*utils.Item[DirCell]{}

	startDir := DirCell{
		dir:  '>',
		Cell: *start,
	}
	distances[startDir] = 0
	startItem := unvisited.HeapPush(startDir, 0)
	itemsByCell[startDir] = startItem

	maxCost := math.MaxInt32

	for unvisited.Len() > 0 {
		curr, _ := unvisited.HeapPop()

		if *maze.GetAtCell(&curr.Cell) == "E" {
			maxCost = distances[curr]
		}
		if distances[curr] > maxCost {
			// no point in going further
			continue
		}

		for _, n := range curr.NeighborsCross() {
			nVal := maze.GetAtCell(&n)
			if nVal != nil && (*nVal == "." || *nVal == "E") {

				nextDir := currentDir(&curr.Cell, &n)
				if isOppositeDir(curr.dir, nextDir) {
					// no point in going back
					continue
				}
				nDir := DirCell{
					dir:  nextDir,
					Cell: n,
				}

				cost := calcCostDir(curr.dir, curr.Cell, n)
				alt := distances[curr] + cost

				if bestDist, ok := distances[nDir]; !ok || alt < bestDist {
					distances[nDir] = alt
					prevs[nDir] = map[DirCell]bool{curr: true}

					if _, hOk := itemsByCell[nDir]; !hOk {
						item := unvisited.HeapPush(nDir, alt)
						itemsByCell[nDir] = item
					} else {
						unvisited.Update(itemsByCell[nDir], alt)
					}
				} else if alt == bestDist {
					prevs[nDir][curr] = true
				}
			}
		}
	}

	return distances, prevs
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
				cost := calcCost(prevs[curr], curr, n)
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

func calcCostDir(currDir rune, curr utils.Cell, next utils.Cell) int {
	nextDir := currentDir(&curr, &next)
	isOpposite := isOppositeDir(currDir, nextDir)
	cost := 1
	if isOpposite {
		cost += 2000
	} else if currDir != nextDir {
		cost += 1000
	}
	return cost
}

func calcCost(prev *utils.Cell, curr utils.Cell, next utils.Cell) int {
	currDir := '>'
	if prev != nil {
		currDir = currentDir(prev, &curr)
	}

	return calcCostDir(currDir, curr, next)
}

func solveMaze(maze *utils.Matrix[string], prev, current *utils.Cell, currCost int, visited *sync.Map) int {

	if *maze.GetAtCell(current) == "E" {
		return currCost
	}

	currDir := currentDir(prev, current)

	next := []utils.Cell{
		current.Up(1),
		current.Right(1),
		current.Down(1),
		current.Left(1),
	}

	results := make(chan int, len(next))
	var wg sync.WaitGroup
	for _, n := range next {
		wg.Add(1)
		go func(n utils.Cell) {
			defer wg.Done()

			nVal := maze.GetAtCell(&n)
			if nVal != nil && (*nVal == "." || *nVal == "E") {

				nextDir := currentDir(current, &n)
				isOpposite := isOppositeDir(currDir, nextDir)
				cost := 1
				if isOpposite {
					cost += 2000
				} else if currDir != nextDir {
					cost += 1000
				}

				if storedCost, ok := visited.Load(n); ok && storedCost.(int) <= currCost+cost {
					return
				}
				visited.Store(n, currCost+cost)

				totCost := solveMaze(maze, current, &n, currCost+cost, visited)

				results <- totCost
			}

		}(n)
	}

	wg.Wait()
	close(results)

	minCost := math.MaxInt32
	for totCost := range results {
		if totCost < minCost {
			minCost = totCost
		}
	}

	return minCost
}

func isOppositeDir(dir rune, dir2 rune) bool {
	return (dir == '^' && dir2 == 'v') || (dir == 'v' && dir2 == '^') || (dir == '<' && dir2 == '>') || (dir == '>' && dir2 == '<')
}

func currentDir(curr *utils.Cell, next *utils.Cell) rune {
	if curr == nil {
		return '>'
	}

	if curr.R == next.R {
		if curr.C < next.C {
			return '>'
		} else if curr.C > next.C {
			return '<'
		}
		panic("invalid direction")
	} else {
		if curr.R < next.R {
			return 'v'
		} else {
			return '^'
		}
	}
}
