package main

import (
	"AdventOfCode2024/utils"
	"strings"
)

func part1(input string) any {

	antennas := map[rune][]*utils.Cell{}
	split := strings.Split(input, "\n")
	maxRow := len(split)
	maxCol := 0
	for r, line := range split {
		maxCol = len(line)
		for i, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], utils.NewCell(r, i))
			}
		}
	}

	antinodes := map[utils.Cell]bool{}

	bounds := &Bounds{
		MaxRow: maxRow,
		MaxCol: maxCol,
	}
	for freq, cells := range antennas {

		for _, cell1 := range cells {
			for _, cell2 := range cells {
				if cell1 == cell2 {
					continue
				}
				_ = freq

				leftCell := cell1
				rightCell := cell2
				if cell1.C > cell2.C {
					leftCell = cell2
					rightCell = cell1
				}

				a1 := utils.NewCell(leftCell.R+(leftCell.R-rightCell.R), leftCell.C-(rightCell.C-leftCell.C))
				if bounds.IsValid(a1) {
					antinodes[*a1] = true
				}

				a2 := utils.NewCell(rightCell.R-(leftCell.R-rightCell.R), rightCell.C+(rightCell.C-leftCell.C))
				if bounds.IsValid(a2) {
					antinodes[*a2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func part2(input string) any {
	antennas := map[rune][]*utils.Cell{}
	split := strings.Split(input, "\n")
	maxRow := len(split)
	maxCol := 0
	for r, line := range split {
		maxCol = len(line)
		for i, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], utils.NewCell(r, i))
			}
		}
	}

	antinodes := map[utils.Cell]bool{}

	bounds := &Bounds{
		MaxRow: maxRow,
		MaxCol: maxCol,
	}
	for _, cells := range antennas {

		for _, cell1 := range cells {
			antinodes[*cell1] = true

			for _, cell2 := range cells {
				if cell1 == cell2 {
					continue
				}

				leftCell := cell1
				rightCell := cell2
				if cell1.C > cell2.C {
					leftCell = cell2
					rightCell = cell1
				}

				//left
				a1 := utils.NewCell(leftCell.R+(leftCell.R-rightCell.R), leftCell.C-(rightCell.C-leftCell.C))
				for bounds.IsValid(a1) {
					antinodes[*a1] = true

					rightCell = leftCell
					leftCell = a1
					a1 = utils.NewCell(leftCell.R+(leftCell.R-rightCell.R), leftCell.C-(rightCell.C-leftCell.C))
				}

				//right
				a2 := utils.NewCell(rightCell.R-(leftCell.R-rightCell.R), rightCell.C+(rightCell.C-leftCell.C))
				for bounds.IsValid(a2) {
					antinodes[*a2] = true

					leftCell = rightCell
					rightCell = a2
					a2 = utils.NewCell(rightCell.R-(leftCell.R-rightCell.R), rightCell.C+(rightCell.C-leftCell.C))
				}
			}
		}
	}

	return len(antinodes)

}

type Bounds struct {
	MaxRow, MaxCol int
}

func (b *Bounds) IsValid(x *utils.Cell) bool {
	return utils.Between(x.R, 0, b.MaxRow) && utils.Between(x.C, 0, b.MaxCol)
}

func part1_line(input string) any {

	antennas := map[rune][]*utils.Cell{}
	split := strings.Split(input, "\n")
	maxRow := len(split)
	maxCol := 0
	for r, line := range split {
		maxCol = len(line)
		for i, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], utils.NewCell(r, i))
			}
		}
	}

	antinodes := map[utils.Cell]bool{}

	_ = maxRow
	for freq, cells := range antennas {

		for _, cell1 := range cells {
			for _, cell2 := range cells {
				if cell1 == cell2 {
					continue
				}
				_ = freq
				line := utils.NewLine(cell1, cell2)

				if line.IsVertical() {
					//go along the column
					minCellRow := min(cell1.R, cell2.R)
					maxCellRow := max(cell1.R, cell2.R)

					for row := minCellRow - 1; row >= 0; row-- {
						x := utils.NewCell(row, cell1.C)

						if isAntinode(line, x) {
							antinodes[*x] = true
							//antinodesByFreq[*x] = true
							//break
						}
					}

					for row := maxCellRow + 1; row < maxRow; row++ {
						x := utils.NewCell(row, cell1.C)

						if isAntinode(line, x) {
							antinodes[*x] = true
							//antinodesByFreq[*x] = true
							//break
						}
					}

				} else {

					minCellCol := min(cell1.C, cell2.C)
					maxCellCol := max(cell1.C, cell2.C)

					//antinode 1
					for col := minCellCol - 1; col >= 0; col-- {
						x := line.CellAtCol(col)
						if !utils.Between(x.R, 0, maxRow) {
							continue
						}

						if isAntinode(line, x) {
							antinodes[*x] = true
							break
						}
					}

					//antinode 2
					for col := maxCellCol + 1; col < maxCol; col++ {
						x := line.CellAtCol(col)
						if !utils.Between(x.R, 0, maxRow) {
							continue
						}

						if isAntinode(line, x) {
							antinodes[*x] = true
							break
						}
					}

				}

			}
		}

		//fmt.Printf("freq %s: %d\n", string(freq), len(antinodes))
	}

	return len(antinodes)
}

func part2_line(input string) any {

	antennas := map[rune][]*utils.Cell{}
	split := strings.Split(input, "\n")
	maxRow := len(split)
	maxCol := 0
	for r, line := range split {
		maxCol = len(line)
		for i, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], utils.NewCell(r, i))
			}
		}
	}

	antinodes := map[utils.Cell]bool{}

	_ = maxRow
	for freq, cells := range antennas {
		for _, cell1 := range cells {
			antinodes[*cell1] = true
			for _, cell2 := range cells {
				if cell1 == cell2 {
					continue
				}
				_ = freq
				line := utils.NewLine(cell1, cell2)

				if line.IsVertical() {
					//go along the column
					minCellRow := min(cell1.R, cell2.R)
					maxCellRow := max(cell1.R, cell2.R)

					for row := minCellRow - 1; row >= 0; row-- {
						x := utils.NewCell(row, cell1.C)

						if isAntinode(line, x) {
							antinodes[*x] = true
							//antinodesByFreq[*x] = true
							//break
						}
					}

					for row := maxCellRow + 1; row < maxRow; row++ {
						x := utils.NewCell(row, cell1.C)

						if isAntinode(line, x) {
							antinodes[*x] = true
							//antinodesByFreq[*x] = true
							//break
						}
					}

				} else {

					minCellCol := min(cell1.C, cell2.C)
					maxCellCol := max(cell1.C, cell2.C)

					//antinode 1
					for col := minCellCol - 1; col >= 0; col-- {
						x := line.CellAtCol(col)
						if !utils.Between(x.R, 0, maxRow) {
							continue
						}

						if isResonant(line, x) {
							antinodes[*x] = true
						}
					}

					//antinode 2
					for col := maxCellCol + 1; col < maxCol; col++ {
						x := line.CellAtCol(col)
						if !utils.Between(x.R, 0, maxRow) {
							continue
						}

						if isResonant(line, x) {
							antinodes[*x] = true
						}
					}
				}
			}
		}
	}

	return len(antinodes)
}

func isAntinode(line *utils.Line, x *utils.Cell) bool {

	dist1 := x.DistManhattan(line.A)
	dist2 := x.DistManhattan(line.B)

	return dist1 == 2*dist2
}

func isResonant(line *utils.Line, x *utils.Cell) bool {
	dist := line.A.DistManhattan(line.B)
	distP := x.DistManhattan(line.A)

	return distP%dist == 0
}
