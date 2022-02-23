package main

import (
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
	Map         [][]WarehouseCell
	Packages    []Package
	Forklifts   []Forklift
	Trucks      []Truck
}

func initMap(x, y int) [][]WarehouseCell {
	w := make([][]WarehouseCell, y)
	for i := 0; i < x; i++ {
		w[i] = make([]WarehouseCell, x)
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
		Graph:       nil,
		Map:         initMap(int(x), int(y)),
		Packages:    nil,
		Forklifts:   nil,
		Trucks:      nil,
	}
}

func (w Warehouse) IsSimulationComplete() bool {
	return w.CurrentTurn >= w.Lifetime
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
			}
			if c.T != nil {
				fmt.Printf("[T] ")
			}
			if c.P != nil {
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
		fmt.Printf("%s %s [%d,%d]\n", f.Name, f.Status, f.Pos.x, f.Pos.y)
	}
	for _, t := range w.Trucks {
		c, _ := t.IsFull()
		fmt.Printf("%s %s %d/%d\n", t.Name, t.Status, c, t.Capacity)
	}
}

func (w *Warehouse) move(d int, f *Forklift) error {
	switch d {
	case 0:
		if f.Pos.x-1 >= 0 {
			w.Map[f.Pos.x][f.Pos.y].F = nil
			f.Pos.Up()
			w.Map[f.Pos.x][f.Pos.y].F = f
		}
	case 1:
		if f.Pos.y+1 < w.Size.y {
			w.Map[f.Pos.x][f.Pos.y].F = nil
			f.Pos.Right()
			w.Map[f.Pos.x][f.Pos.y].F = f
		}
	case 2:
		if f.Pos.x+1 < w.Size.x {
			w.Map[f.Pos.x][f.Pos.y].F = nil
			f.Pos.Down()
			w.Map[f.Pos.x][f.Pos.y].F = f
		}
	case 3:
		if f.Pos.y-1 >= 0 {
			w.Map[f.Pos.x][f.Pos.y].F = nil
			f.Pos.Left()
			w.Map[f.Pos.x][f.Pos.y].F = f
		}
	}
	return nil
}

func (w Warehouse) SelectForkliftObjective(f *Forklift) {
	if f.Package == nil {
		dist := -1.0
		target := 0
		for i, p := range w.Packages {
			xd := p.Pos.x - f.Pos.x
			yd := p.Pos.y - f.Pos.y
			if dist < 0 {
				dist = math.Sqrt(float64(xd*xd + yd*yd))
				target = i
			} else {
				if dist > math.Sqrt(float64(xd*xd+yd*yd)) {
					dist = math.Sqrt(float64(xd*xd + yd*yd))
					target = i
				}
			}
		}
		f.TargetPos = w.Packages[target].Pos
	} else {
		dist := -1.0
		target := 0
		for i, t := range w.Trucks {
			xd := t.Pos.x - f.Pos.x
			yd := t.Pos.y - f.Pos.y
			if dist < 0 {
				dist = math.Sqrt(float64(xd*xd + yd*yd))
				target = i
			} else {
				if dist > math.Sqrt(float64(xd*xd+yd*yd)) {
					dist = math.Sqrt(float64(xd*xd + yd*yd))
					target = i
				}
			}
		}
		f.TargetPos = w.Trucks[target].Pos
	}
}

func (w *Warehouse) CreateGraph() {
	for i := 0; i < w.Size.x; i++ {
		for j := 0; j < w.Size.y; j++ {
			w.Graph.AddNode(&Node{
				p: Position{
					x: i,
					y: j,
				},
				F: nil,
				T: nil,
				P: nil,
			})
		}
	}
}
