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
	ID        int
	Path      []int
	Pos       Position
	TargetPos Position
	Target    int
	Package   *Package
	Status    FStatus
	Name      string
}

//TODO
// Change target for id
// comparator position
// getmapcellby position

//NewForklift
func NewForklift(ID, x, y int, name string) Forklift {
	return Forklift{
		ID:   ID,
		Path: nil,
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

func (f *Forklift) Reset() {
	f.ResetPath()
	f.ResetTarget()
	f.Package = nil
	f.updateStatus(WAIT)
}

func (f *Forklift) ResetPath() {
	f.Path = nil
}

func (f *Forklift) AddPath(path []int) {
	f.Path = path
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

/*func (f *Forklift) move(d int) {
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
}*/

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

func (f *Forklift) Move() {
	f.ID = f.Path[0]
	f.Path = remove(f.Path, 0)
}
