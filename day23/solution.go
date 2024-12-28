package main

import (
	"AdventOfCode2024/utils"
	"slices"
	"strings"
)

func part1(input string) any {
	network := &Graph{
		edges: make(map[Edge]bool),
		vert:  make(map[string][]*Edge),
	}
	tStarting := map[string]bool{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		connection := strings.Split(line, "-")
		network.AddEdge(connection[0], connection[1])

		if connection[0][0] == 't' {
			tStarting[connection[0]] = true
		}
		if connection[1][0] == 't' {
			tStarting[connection[1]] = true
		}
	}

	threeOrMore := 0
	compSubsets := map[string][]string{}

	for pc, _ := range tStarting {
		edges := network.vert[pc]
		var neighbors []string
		for _, edge := range edges {
			neighbors = append(neighbors, edge.Other(pc))
		}

		partSubsets := subsets(neighbors, 2)
		for _, subset := range partSubsets {
			tmp := append(subset, pc)
			slices.Sort(tmp)
			compSubsets[strings.Join(tmp, "")] = tmp
		}

	}
	for _, subset := range compSubsets {
		if network.IsFullyConnected(subset) {
			threeOrMore++
		}
	}

	return threeOrMore
}

func part2(input string) any {
	return part2_BronKerbosch(input)
	//return part2_v1(input)
}

func part2_v1(input string) any {
	network := &Graph{
		edges: make(map[Edge]bool),
		vert:  make(map[string][]*Edge),
	}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		connection := strings.Split(line, "-")
		network.AddEdge(connection[0], connection[1])
	}

	maxByVert := map[string]string{}
	currMaxSize := 1
	currMaxKey := ""
	for pc, _ := range network.vert {

		neighbors := network.Neighbors(pc)

		for i := len(neighbors); i >= currMaxSize; i-- {
			compSubsets := map[string][]string{}
			partSubsets := subsets(neighbors, i)
			for _, subset := range partSubsets {
				tmp := append(subset, pc)
				slices.Sort(tmp)
				compSubsets[strings.Join(tmp, ",")] = tmp
			}
			for key, subset := range compSubsets {
				if network.IsFullyConnected(subset) {
					maxByVert[pc] = key
					currMaxSize = len(subset)
					currMaxKey = key
					break
				}
			}
			if _, ok := maxByVert[pc]; ok {
				break
			}
		}
	}

	return currMaxKey
}

func part2_BronKerbosch(input string) any {
	network := &Graph{
		edges: make(map[Edge]bool),
		vert:  make(map[string][]*Edge),
	}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		connection := strings.Split(line, "-")
		network.AddEdge(connection[0], connection[1])
	}

	cliques := network.BronKerbosch()

	var maxClique []string
	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}
	slices.Sort(maxClique)

	return strings.Join(maxClique, ",")
}

type Edge struct {
	a string
	b string
}

func NewEdge(a, b string) Edge {
	if a > b {
		a, b = b, a
	}
	return Edge{a: a, b: b}
}

func (e Edge) Other(node string) string {
	if e.a == node {
		return e.b
	}
	return e.a
}

type Graph struct {
	edges map[Edge]bool
	vert  map[string][]*Edge
}

func (g *Graph) AddEdge(a, b string) {
	edge := NewEdge(a, b)
	g.edges[edge] = true
	g.vert[a] = append(g.vert[a], &edge)
	g.vert[b] = append(g.vert[b], &edge)
}

func (g *Graph) IsFullyConnected(nodes []string) bool {
	for _, nodeS := range nodes {
		for _, nodeT := range nodes {
			if nodeS == nodeT {
				continue
			}
			if !g.edges[NewEdge(nodeS, nodeT)] {
				return false
			}
		}
	}

	return true
}

func subsetsRec(nodes []string, minSize int, index int, current []string) [][]string {
	if len(current) >= minSize {
		return [][]string{append([]string{}, current...)}
	}

	var allSubsets [][]string
	for i := index; i < len(nodes); i++ {
		current = append(current, nodes[i])
		allSubsets = append(allSubsets, subsetsRec(nodes, minSize, i+1, current)...)
		current = current[:len(current)-1]
	}

	return allSubsets
}

func subsets(nodes []string, minSize int) [][]string {
	return subsetsRec(nodes, minSize, 0, []string{})
}

func (g *Graph) Neighbors(node string) []string {
	edges := g.vert[node]
	var neighbors []string
	for _, edge := range edges {
		neighbors = append(neighbors, edge.Other(node))
	}
	return neighbors
}
func (g *Graph) NeighborsSet(node string) map[string]bool {
	edges := g.vert[node]
	neighbors := map[string]bool{}
	for _, edge := range edges {
		neighbors[edge.Other(node)] = true
	}
	return neighbors
}

func (g *Graph) BronKerbosch() [][]string {
	return g._BronKerbosch([]string{}, utils.Keys(g.vert), []string{})
}

func (g *Graph) _BronKerbosch(R []string, P []string, X []string) [][]string {
	if len(P) == 0 && len(X) == 0 {
		return [][]string{append([]string{}, R...)}
	}

	var allCliques [][]string

	var pivot string
	if len(P) > 0 {
		pivot = P[0]
	} else {
		pivot = X[0]
	}
	pivotNeighbors := g.NeighborsSet(pivot)
	pivotedP := utils.Filter(P, func(x string) bool { return !pivotNeighbors[x] })
	for _, v := range pivotedP {
		neighbors := g.NeighborsSet(v)

		p := utils.Filter(P, func(x string) bool { return neighbors[x] })
		x := utils.Filter(X, func(x string) bool { return neighbors[x] })

		vCliques := g._BronKerbosch(append(R, v), p, x)
		allCliques = append(allCliques, vCliques...)

		P = utils.Filter(P, func(x string) bool { return x != v })
		X = append(X, v)
	}

	return allCliques
}
