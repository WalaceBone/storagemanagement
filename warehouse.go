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
