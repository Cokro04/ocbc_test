package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	u, v   int
	weight int
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	return &DSU{p, r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == rootY {
		return false
	}

	if d.rank[rootX] < d.rank[rootY] {
		d.parent[rootX] = rootY
	} else if d.rank[rootX] > d.rank[rootY] {
		d.parent[rootY] = rootX
	} else {
		d.parent[rootY] = rootX
		d.rank[rootX]++
	}

	return true
}

func getMinimumCostMST(graph_nodes int, graph_from, graph_to, graph_weight []int, source, destination int) int {
	edges := []Edge{}
	for i := range graph_from {
		edges = append(edges, Edge{graph_from[i], graph_to[i], graph_weight[i]})
	}

	// Sort edge by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(graph_nodes)
	totalCost := 0

	for _, e := range edges {
		if dsu.Union(e.u, e.v) {
			totalCost += e.weight

			// Stop early once source and destination are connected
			if dsu.Find(source) == dsu.Find(destination) {
				return totalCost
			}
		}
	}

	return -1
}

func main() {
	graph_nodes := 3
	graph_from := []int{1, 2, 1}
	graph_to := []int{2, 3, 3}
	graph_weight := []int{5, 3, 4}

	fmt.Println(getMinimumCostMST(graph_nodes, graph_from, graph_to, graph_weight, 1, 3)) // Output: 7
}
