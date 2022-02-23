package main

type Cell struct {
	p Position
	F *Forklift
	T *Truck
	P *Package
}

func (c Cell) IsEmpty() bool {
	return c.F != nil && c.T != nil && c.P != nil
}

func NewCell(p Position) *Cell {
	return &Cell{
		p: Position{
			x: p.x,
			y: p.y,
		},
		F: nil,
		T: nil,
		P: nil,
	}
}
