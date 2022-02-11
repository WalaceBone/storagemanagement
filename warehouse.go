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

func (p Position) Add(x, y uint) {
	p.x += x
	p.y += y
}

func (p Position) Sub(x, y uint) {
	if p.x > 0 && p.y > 0 {
		p.x -= x
		p.y -= y
	}
}
