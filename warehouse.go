package main

type Position struct {
	x, y uint
}

type Size struct {
	w, l uint
}

type Warehouse struct {
	Size      Size
	Lifetime  uint
	Packages  []Package
	Forklifts []Forklift
	Trucks    []Truck
}
