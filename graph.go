package main

import "sync"

type Node struct {
	ID      int
	visited bool
}

type ItemGraph struct {
	nodes []Node
	edges map[int][]*Node
	lock  sync.RWMutex
}

func NewGraph() *ItemGraph {
	return &ItemGraph{
		nodes: make([]Node, 0),
		edges: nil,
		lock:  sync.RWMutex{},
	}
}

func (g *ItemGraph) getNodeByID(ID int) *Node {
	for i, node := range g.nodes {
		if node.ID == ID {
			return &g.nodes[i]
		}
	}
	return nil
}

func (n *Node) Visited() {
	n.visited = true
}

func (g *ItemGraph) AddNode(n int) {
	g.lock.Lock()
	if g.nodes == nil {
		g.nodes = make([]Node, 1)
	}
	g.nodes = append(g.nodes, Node{n, false})
	g.lock.Unlock()
}

func (g *ItemGraph) AddEdge(n1, n2 int) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[int][]*Node)
	}
	g.edges[n1] = append(g.edges[n1], g.getNodeByID(n2))
	//g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}
