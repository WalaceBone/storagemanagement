package main

type weight int

const (
	YELLOW weight = 100
	GREEN  weight = 200
	BLUE   weight = 500
)

type Package struct {
	Weight weight
	Pos    Position
	Name   string
}
