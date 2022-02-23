package main

import "sync"

type ItemGraph struct {
	nodes []*Cell
	edges map[Cell][]*Cell
	lock  sync.RWMutex
}

func (g *ItemGraph) AddNode(n *Cell) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *ItemGraph) AddEdge(n1, n2 *Cell) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Cell][]*Cell)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}
