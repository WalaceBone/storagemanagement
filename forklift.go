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
	Pos       Position
	TargetPos Position
	Package   *Package
	Status    FStatus
	Name      string
}

//NewForklift
func NewForklift(x, y int, name string) Forklift {
	return Forklift{
		Pos: Position{
			x: x,
			y: y,
		},
		TargetPos: Position{
			x: -1,
			y: -1,
		},
		Package: nil,
		Status:  WAIT,
		Name:    name,
	}
}

func (f *Forklift) updateStatus(s FStatus) {
	f.Status = s
}

func (f Forklift) IsTargetSelected() bool {
	return f.TargetPos.x != -1 && f.TargetPos.y != -1
}

func (f *Forklift) ResetTarget() {
	f.TargetPos.x = -1
	f.TargetPos.y = -1
}

func (f *Forklift) move(d int) {
	switch d {
	case 0:
		f.Pos.Up()
	case 1:
		f.Pos.Right()
	case 2:
		f.Pos.Down()
	case 3:
		f.Pos.Left()
	}
}

//Dump
func (f Forklift) Dump() {
	fmt.Printf("\tName: %s\n", f.Name)
	fmt.Printf("\tPosition: [%d,%d]\n", f.Pos.x, f.Pos.y)
	fmt.Printf("\tStatus: %s\n", f.Status)
	fmt.Printf("Target Pos: [%d,%d]\n", f.TargetPos.x, f.TargetPos.y)
	fmt.Printf("\tPackage: \n")
	if f.Package != nil {
		f.Package.Dump()
	} else {
		fmt.Printf("empty\n")
	}
}
