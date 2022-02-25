package main

import (
	"container/list"
	"fmt"
	"math"
)

type Size struct {
	x, y int
}

type Warehouse struct {
	Size        Size
	Lifetime    int
	CurrentTurn int
	Graph       *ItemGraph
	Map         [][]Cell
	Packages    []Package
	Forklifts   []Forklift
	Trucks      []Truck
}

func initMap(x, y int) [][]Cell {
	w := make([][]Cell, x)
	id := 0
	for i := 0; i < x; i++ {
		w[i] = make([]Cell, y)
		for j := 0; j < y; j++ {
			w[i][j] = *NewCell(id, Position{
				x: i,
				y: j,
			})
			id += 1
		}
	}
	return w
}

//NewWarehouse
func NewWarehouse(x, y, lifetime int) Warehouse {
	return Warehouse{
		Size: Size{
			x: x,
			y: y,
		},
		Lifetime:    lifetime,
		CurrentTurn: 0,
		Graph:       NewGraph(),
		Map:         initMap(x, y),
		Packages:    nil,
		Forklifts:   nil,
		Trucks:      nil,
	}
}

func (w Warehouse) IsSimulationComplete() bool {

	if len(w.Packages) == 0 || w.PackageLeft() == 0 {
		fmt.Println("😎")
		return true
	} else if w.CurrentTurn >= w.Lifetime {
		fmt.Println("🙂")
		return true
	}
	return false
}

func (w *Warehouse) decountLifeTime() {
	w.CurrentTurn++
}

func (w *Warehouse) addTruck(t Truck) error {
	if w.Trucks == nil {
		w.Trucks = make([]Truck, 0)
	}
	w.Trucks = append(w.Trucks, t)
	w.Map[t.Pos.x][t.Pos.y].T = &w.Trucks[len(w.Trucks)-1]
	return nil
}

func (w *Warehouse) addPackage(p Package) error {
	if w.Packages == nil {
		w.Packages = make([]Package, 0)
	}
	w.Packages = append(w.Packages, p)
	w.Map[p.Pos.x][p.Pos.y].P = &w.Packages[len(w.Packages)-1]
	return nil
}

func (w *Warehouse) addForklift(f Forklift) error {
	if w.Forklifts == nil {
		w.Forklifts = make([]Forklift, 0)
	}
	w.Forklifts = append(w.Forklifts, f)
	w.Map[f.Pos.x][f.Pos.y].F = &w.Forklifts[len(w.Forklifts)-1]
	return nil
}

//Dump
func (w Warehouse) Dump() {
	fmt.Printf("Warehouse\n")
	fmt.Printf("\tSize: [%d,%d]\n", w.Size.x, w.Size.y)
	fmt.Printf("\tLifetime: %d\n", w.Lifetime)
	fmt.Printf("\tCurrent Turn: %d\n", w.CurrentTurn)
	fmt.Printf("\nForklifts:\n")
	for _, f := range w.Forklifts {
		f.Dump()
	}
	fmt.Printf("\nTrucks:\n")
	for _, t := range w.Trucks {
		t.Dump()
		fmt.Printf("\n")
	}
	fmt.Printf("\nPackages:\n")
	for _, p := range w.Packages {
		p.Dump()
	}
	w.DumpMap()
}

func (w Warehouse) DumpMap() {
	for _, cells := range w.Map {
		for _, c := range cells {
			if c.F != nil {
				fmt.Printf("[F] ")
			} else if c.T != nil {
				fmt.Printf("[T] ")
			} else if c.P != nil {
				fmt.Printf("[P] ")
			} else {
				fmt.Printf("[ ] ")
			}
		}
		fmt.Printf("\n")
	}
}

//DumpTurn
func (w Warehouse) DumpTurn() {
	fmt.Printf("Turn %d\n", w.CurrentTurn)
	for _, f := range w.Forklifts {
		switch f.Status {
		case "GO":
			fmt.Printf("%s %s [%d,%d]\n", f.Name, f.Status, f.Pos.x, f.Pos.y)
		case "WAIT":
			fmt.Printf("%s %s\n", f.Name, f.Status)
		case "TAKE", "LEAVE":
			color := f.Package.WeightToColor()
			fmt.Printf("%s %s %s %s\n", f.Name, f.Status, f.Package.Name, color)
		}
	}
	for _, t := range w.Trucks {
		c, _ := t.IsFull()
		fmt.Printf("%s %s %d/%d\n", t.Name, t.Status, c, t.Capacity)
	}
}

//func (w *Warehouse) move(d int, f *Forklift) error {
//	switch d {
//	case 0:
//		if f.Pos.x-1 >= 0 {
//			w.Map[f.Pos.x][f.Pos.y].F = nil
//			f.Pos.Up()
//			w.Map[f.Pos.x][f.Pos.y].F = f
//		}
//	case 1:
//		if f.Pos.y+1 < w.Size.y {
//			w.Map[f.Pos.x][f.Pos.y].F = nil
//			f.Pos.Right()
//			w.Map[f.Pos.x][f.Pos.y].F = f
//		}
//	case 2:
//		if f.Pos.x+1 < w.Size.x {
//			w.Map[f.Pos.x][f.Pos.y].F = nil
//			f.Pos.Down()
//			w.Map[f.Pos.x][f.Pos.y].F = f
//		}
//	case 3:
//		if f.Pos.y-1 >= 0 {
//			w.Map[f.Pos.x][f.Pos.y].F = nil
//			f.Pos.Left()
//			w.Map[f.Pos.x][f.Pos.y].F = f
//		}
//	}
//	return nil
//}

func (w Warehouse) SelectForkliftTarget(f *Forklift) {
	if f.Package == nil {
		closest := -1.0
		target := 0
		for i, p := range w.Packages {
			if !p.Loaded && !p.Targeted {
				xd := p.Pos.x - f.Pos.x
				yd := p.Pos.y - f.Pos.y
				dist := math.Sqrt(float64(xd*xd + yd*yd))
				if closest < 0 || closest > dist {
					closest = dist
					target = i
				}
			}
		}
		f.TargetPos = w.Packages[target].Pos
		f.Target = w.Packages[target].ID
		w.GetPackageByID(f.Target).Target()
		//TODO
		// change when nothing left to pick
	} else {
		closest := -1.0
		target := 0
		for i, t := range w.Trucks {
			xd := t.Pos.x - f.Pos.x
			yd := t.Pos.y - f.Pos.y
			dist := math.Sqrt(float64(xd*xd + yd*yd))
			if closest < 0 || closest > dist {
				closest = dist
				target = i
			}
			if t.CanReceive(f.Package.Weight) {
				target = i
				break
			}
		}
		f.TargetPos = w.Trucks[target].Pos
		f.Target = w.Trucks[target].ID
	}
}

func (w *Warehouse) CreateGraph() {
	id := 0
	for i := 0; i < w.Size.x; i++ {
		for j := 0; j < w.Size.y; j++ {
			w.Graph.AddNode(id)
			id++
		}
	}
}

func (w *Warehouse) GetPackageByID(ID int) *Package {
	for i, p := range w.Packages {
		if p.ID == ID {
			return &w.Packages[i]
		}
	}
	return nil
}

func (w Warehouse) GetCellById(id int) *Cell {
	for i, cells := range w.Map {
		for j, cell := range cells {
			if cell.ID == id {
				return &w.Map[i][j]
			}
		}
	}
	return nil
}

func (w Warehouse) GetCellIDFromPosition(x, y int) int {
	for _, cells := range w.Map {
		for _, cell := range cells {
			if cell.p.x == x && cell.p.y == y {
				return cell.ID
			}
		}
	}
	return -1
}

func (w *Warehouse) CreateEdges() {
	for i, node := range w.Graph.nodes {
		//add edge up
		cell := w.GetCellById(node.ID)
		if cell.p.x-1 >= 0 {
			w.Graph.AddEdge(w.Graph.nodes[i].ID, w.GetCellIDFromPosition(cell.p.x-1, cell.p.y))
		}
		//add edge down
		if cell.p.x+1 < w.Size.x {
			w.Graph.AddEdge(w.Graph.nodes[i].ID, w.GetCellIDFromPosition(cell.p.x+1, cell.p.y))
		}
		//add edge left
		if cell.p.y-1 >= 0 {
			w.Graph.AddEdge(w.Graph.nodes[i].ID, w.GetCellIDFromPosition(cell.p.x, cell.p.y-1))
		}
		//add edge right
		if cell.p.y+1 < w.Size.y {
			w.Graph.AddEdge(w.Graph.nodes[i].ID, w.GetCellIDFromPosition(cell.p.x, cell.p.y+1))
		}
	}
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (g *ItemGraph) GetNodeByID(ID int) *Node {
	for i, node := range g.nodes {
		if node.ID == ID {
			return &g.nodes[i]
		}
	}
	return nil
}

func GetPath(src, tgt int, path map[int]int) []int {
	spath := make([]int, 0)
	tmp := tgt
	spath = append(spath, tgt)
	for path[tmp] != src {
		spath = append(spath, path[tmp])
		tmp = path[tmp]
	}
	for i, j := 0, len(spath)-1; i < j; i, j = i+1, j-1 {
		spath[i], spath[j] = spath[j], spath[i]
	}
	return spath
}

func (w Warehouse) FindPath(src, tgt int) []int {
	for i := range w.Graph.nodes {
		w.Graph.nodes[i].Reset()
	}
	fmt.Println(src, tgt)
	queue := list.New()
	queue.PushBack(w.Graph.getNodeByID(src))

	camefrom := make(map[int]int)

	camefrom[src] = -1
	w.Graph.GetNodeByID(src).Visited()
	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(queue.Front())

		if current.Value.(*Node).ID == tgt {
			return GetPath(src, tgt, camefrom)
		}
		neighbours := w.Graph.edges[current.Value.(*Node).ID]
		for _, neighbour := range neighbours {
			if !w.GetCellById(neighbour.ID).IsEmpty() && neighbour.ID != tgt {
				w.Graph.GetNodeByID(neighbour.ID).Visited()
			}
			if !neighbour.visited {
				w.Graph.GetNodeByID(neighbour.ID).Visited()
				queue.PushBack(neighbour)
				camefrom[neighbour.ID] = current.Value.(*Node).ID
			}
		}
	}
	return nil
}

func (w Warehouse) PackageTargeted() int {
	pleft := 0
	for _, p := range w.Packages {
		if p.Targeted {
			pleft++
		}
	}
	return pleft
}

func (w Warehouse) PackageLeft() int {
	pleft := 0
	for _, p := range w.Packages {
		if !p.Loaded {
			pleft++
		}
	}
	return pleft
}
