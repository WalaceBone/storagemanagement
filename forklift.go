package main

import "fmt"

type FStatus string

const (
	GO    FStatus = "GO"
	WAIT  FStatus = "WAIT"
	TAKE  FStatus = "TAKE"
	LEAVE FStatus = "LEAVE"
)

type Forklift struct {
	Pos     Position
	Package *Package
	Status  FStatus
	Name    string
}

//NewForklift
func NewForklift(x, y uint, name string) *Forklift {
	return &Forklift{
		Pos: Position{
			x: x,
			y: y,
		},
		Package: nil,
		Status:  WAIT,
		Name:    name,
	}
}

//Dump
func (f Forklift) Dump() {
	fmt.Printf("\tName: %s\n", f.Name)
	fmt.Printf("\tPosition: [%d,%d]\n", f.Pos.x, f.Pos.y)
	fmt.Printf("\tStatus: %s\n", f.Status)
	fmt.Printf("\tPackage: \n")
	if f.Package != nil {
		f.Package.Dump()
	} else {
		fmt.Printf("empty\n")
	}
}
