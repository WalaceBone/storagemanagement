package main

type Position struct {
	x, y int
}

type Size struct {
	w, l int
}

type Warehouse struct {
	Size      Size
	Lifetime  uint
	Packages  []Package
	Forklifts []Forklift
	Trucks    []Truck
}
