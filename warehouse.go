package main

import "fmt"

type Position struct {
	x, y uint
}

type Size struct {
	w, l uint
}

type Warehouse struct {
	Size        Size
	Lifetime    uint
	CurrentTurn uint
	Packages    []Package
	Forklifts   []Forklift
	Trucks      []Truck
}

//NewWarehouse
func NewWarehouse(w, l, lifetime uint) Warehouse {
	return Warehouse{
		Size: Size{
			w: w,
			l: l,
		},
		Lifetime:    lifetime,
		CurrentTurn: 0,
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

func (p Position) Add(pos Position) {
	p.x += pos.x
	p.y += pos.y
}

func (p Position) Sub(pos Position) {
	if p.x > 0 && p.y > 0 {
		p.x -= pos.x
		p.y -= pos.y
	}
}

func (w *Warehouse) addTruck(t Truck) error {
	if w.Trucks == nil {
		w.Trucks = make([]Truck, 0)
	}
	w.Trucks = append(w.Trucks, t)
	return nil
}

func (w *Warehouse) addPackage(p Package) error {
	if w.Packages == nil {
		w.Packages = make([]Package, 0)
	}
	w.Packages = append(w.Packages, p)
	return nil
}

func (w *Warehouse) addForklift(f Forklift) error {
	if w.Forklifts == nil {
		w.Forklifts = make([]Forklift, 0)
	}
	w.Forklifts = append(w.Forklifts, f)
	return nil
}

//Dump
func (w Warehouse) Dump() {
	fmt.Printf("Warehouse\n")
	fmt.Printf("\tSize: [%d,%d]\n", w.Size.w, w.Size.l)
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
