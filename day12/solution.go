package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"maps"
)

type Region struct {
	visited   map[utils.Cell]bool
	perimeter int
	seed      rune
}

func (r *Region) String() string {
	cells := ""
	for cell, _ := range r.visited {
		cells += cell.String()
	}

	return fmt.Sprintf("[seed: %s, area: %d, perimeter: %d, price: %d]\n%s\n", string(r.seed), len(r.visited), r.perimeter, r.price(), cells)
}

func (r *Region) price() int {
	return len(r.visited) * r.perimeter
}
func (r *Region) merge(o *Region) {
	r.perimeter += o.perimeter
	for cell, _ := range o.visited {
		r.visited[cell] = true
	}
}

func part1(input string) any {
	matrix := utils.NewMatrixFromLines(input)
	var toVisit []*utils.Cell
	for _, ok := matrix.Next(); ok; _, ok = matrix.Next() {
		toVisit = append(toVisit, matrix.CurrCell())
	}

	regByCell := map[utils.Cell]*Region{}

	for len(toVisit) > 0 {

		var currCell *utils.Cell
		if len(toVisit) > 0 {
			currCell = toVisit[0]
			toVisit = toVisit[1:]
		}

		if currCell == nil {
			panic("nil cell")
		}
		currVal := matrix.GetAtCell(currCell)

		if _, ok := regByCell[*currCell]; !ok {
			newReg := &Region{
				visited:   map[utils.Cell]bool{*currCell: true},
				perimeter: 0,
				seed:      *currVal,
			}
			regByCell[*currCell] = newReg
		}
		regByCell[*currCell].visited[*currCell] = true

		nextCells := []utils.Cell{
			currCell.Up(1),
			currCell.Right(1),
			currCell.Down(1),
			currCell.Left(1)}

		for _, cell := range nextCells {
			val := matrix.GetAtCell(&cell)

			//out of bounds
			if val == nil {
				regByCell[*currCell].perimeter += 1
				continue
			}

			//another region
			if *val != regByCell[*currCell].seed {
				regByCell[*currCell].perimeter += 1
				continue
			}

			if _, ok := regByCell[cell]; !ok {
				//regByCell[cell] = regByCell[*currCell]
			} else if regByCell[cell] != regByCell[*currCell] {
				oldReg := regByCell[*currCell]
				regByCell[cell].merge(oldReg)
				regByCell[*currCell] = regByCell[cell]
				for otherRegCell := range oldReg.visited {
					regByCell[otherRegCell] = regByCell[cell]
				}
			} else {
				//same region nothing to do
			}

		}
	}

	values := maps.Values(regByCell)
	uqRegions := map[*Region]bool{}

	for reg := range values {
		uqRegions[reg] = true
	}
	sum := 0
	for reg, _ := range uqRegions {
		sum += reg.price()
	}

	return sum
}

type Region2 struct {
	visited   map[utils.Cell]bool
	perimeter map[utils.Cell]bool
	seed      rune
}

func (r *Region2) corners(x *utils.Cell) int {
	corners := 0
	_, l := r.visited[x.Left(1)]
	_, u := r.visited[x.Up(1)]
	_, rx := r.visited[x.Right(1)]
	_, d := r.visited[x.Down(1)]

	if !l && !u {
		corners++ // top left
	}
	if !u && !rx {
		corners++ // top right
	}
	if !rx && !d {
		corners++ // bottom right
	}
	if !d && !l {
		corners++ // bottom left
	}

	// concave angles
	_, ld := r.visited[x.Left(1).Down(1)]
	_, lu := r.visited[x.Left(1).Up(1)]
	_, ru := r.visited[x.Right(1).Up(1)]
	_, rd := r.visited[x.Right(1).Down(1)]

	if rx && d && !rd {
		corners++
	}
	if l && d && !ld {
		corners++
	}
	if u && l && !lu {
		corners++
	}
	if u && rx && !ru {
		corners++
	}

	return corners
}

func (r *Region2) sides() int {

	tot := 0
	for c, _ := range r.visited {
		corners := r.corners(&c)
		tot += corners
	}

	return tot
}
func (r *Region2) price() int {
	return len(r.visited) * r.sides()
}

func (r *Region2) merge(o *Region2) {
	for cell := range o.perimeter {
		r.perimeter[cell] = true
	}
	for cell := range o.visited {
		r.visited[cell] = true
	}
}

func (r *Region2) String() string {
	cells := ""
	for cell, _ := range r.perimeter {
		cells += cell.String()
	}

	return fmt.Sprintf("[seed: %s, area: %d, perimeter: %d, sides: %d, price: %d]\n%s\n",
		string(r.seed), len(r.visited), len(r.perimeter), r.sides(), r.price(), cells)
}

func part2(input string) any {
	matrix := utils.NewMatrixFromLines(input)
	var toVisit []*utils.Cell
	for _, ok := matrix.Next(); ok; _, ok = matrix.Next() {
		toVisit = append(toVisit, matrix.CurrCell())
	}

	regByCell := map[utils.Cell]*Region2{}

	for len(toVisit) > 0 {

		var currCell *utils.Cell
		if len(toVisit) > 0 {
			currCell = toVisit[0]
			toVisit = toVisit[1:]
		}

		if currCell == nil {
			panic("nil cell")
		}
		currVal := matrix.GetAtCell(currCell)

		if _, ok := regByCell[*currCell]; !ok {
			newReg := &Region2{
				visited:   map[utils.Cell]bool{*currCell: true},
				perimeter: map[utils.Cell]bool{},
				seed:      *currVal,
			}
			regByCell[*currCell] = newReg
		}
		regByCell[*currCell].visited[*currCell] = true

		nextCells := []utils.Cell{
			currCell.Up(1),
			currCell.Right(1),
			currCell.Down(1),
			currCell.Left(1)}

		for _, cell := range nextCells {
			val := matrix.GetAtCell(&cell)

			//out of bounds
			if val == nil {
				regByCell[*currCell].perimeter[*currCell] = true
				continue
			}

			//another region
			if *val != regByCell[*currCell].seed {
				regByCell[*currCell].perimeter[*currCell] = true
				continue
			}

			if _, ok := regByCell[cell]; ok && regByCell[cell] != regByCell[*currCell] {
				oldReg := regByCell[*currCell]
				regByCell[cell].merge(oldReg)
				regByCell[*currCell] = regByCell[cell]
				for otherRegCell := range oldReg.visited {
					regByCell[otherRegCell] = regByCell[cell]
				}
			}
		}
	}

	values := maps.Values(regByCell)
	uqRegions := map[*Region2]bool{}

	for reg := range values {
		uqRegions[reg] = true
	}
	sum := 0
	for reg, _ := range uqRegions {
		sum += reg.price()
	}

	return sum

}
