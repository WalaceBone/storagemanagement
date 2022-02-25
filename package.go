package main

import "fmt"

const (
	YELLOW = 100
	GREEN  = 200
	BLUE   = 500
)

type Package struct {
	ID       int
	Targeted bool
	Loaded   bool
	Weight   int
	Pos      Position
	Name     string
}

func NewPackage(ID, w, x, y int, name string) Package {
	return Package{
		ID:       ID,
		Targeted: false,
		Loaded:   false,
		Weight:   w,
		Pos: Position{
			x: x,
			y: y,
		},
		Name: name,
	}
}

func (p *Package) Target() {
	p.Targeted = true
}

func (p Package) updatePosition(x, y int) {
	p.Pos.x += x
	p.Pos.y += y
}

func (p Package) WeightToColor() string {
	colors := map[int]string{
		100: "YELLOW",
		200: "GREEN",
		500: "BLUE",
	}
	return colors[p.Weight]
}

func (p Package) Dump() {
	fmt.Printf("\tName: %s\n", p.Name)
	fmt.Printf("\tWeight: %d\n", p.Weight)
	fmt.Printf("\tPosition: [%d,%d]\n", p.Pos.x, p.Pos.y)
}

func (p *Package) Load() {
	p.Loaded = true
}
