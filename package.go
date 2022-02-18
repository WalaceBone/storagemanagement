package main

import "fmt"

const (
	YELLOW = 100
	GREEN  = 200
	BLUE   = 500
)

type Package struct {
	Weight uint
	Pos    Position
	Name   string
}

func NewPackage(w, x, y uint, name string) Package {
	return Package{
		Weight: w,
		Pos: Position{
			x: x,
			y: y,
		},
		Name: name,
	}
}

func (p Package) updatePosition(x, y uint) {
	p.Pos.x += x
	p.Pos.y += y
}

func (p Package) Dump() {
	fmt.Printf("\tName: %s\n", p.Name)
	fmt.Printf("\tWeight: %d\n", p.Weight)
	fmt.Printf("\tPosition: [%d,%d]\n", p.Pos.x, p.Pos.y)
}
