package main

import "sync"

type ItemGraph struct {
	nodes []*Cell
	edges map[Cell][]*Cell
	lock  sync.RWMutex
}

func NewGraph() *ItemGraph {
	return &ItemGraph{
		nodes: make([]*Cell, 0),
		edges: nil,
		lock:  sync.RWMutex{},
	}
}

func (g *ItemGraph) AddNode(n *Cell) {
	g.lock.Lock()
	if g.nodes == nil {
		g.nodes = make([]*Cell, 1)
	}
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *ItemGraph) AddEdge(n1, n2 *Cell) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Cell][]*Cell)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	//g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

func (g *ItemGraph) GetNodeFromPosition(x, y int) *Cell {
	for i, node := range g.nodes {
		if node.p.x == x && node.p.y == y {
			return g.nodes[i]
		}
	}
	return nil
}
