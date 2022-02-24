package main

type Cell struct {
	ID int
	p  Position
	F  *Forklift
	T  *Truck
	P  *Package
}

func (c Cell) IsEmpty() bool {
	return c.F == nil && c.T == nil && c.P == nil
}

func NewCell(ID int, p Position) *Cell {
	return &Cell{
		ID: ID,
		p: Position{
			x: p.x,
			y: p.y,
		},
		F: nil,
		T: nil,
		P: nil,
	}
}
