package main

import "fmt"

type TStatus string

const (
	GONE    TStatus = "GONE"
	WAITING TStatus = "WAITING"
)

type Truck struct {
	Status          TStatus
	Packages        []*Package
	Capacity        uint
	Pos             Position
	Cooldown        uint
	CurrentCooldown uint
	Name            string
}

func NewTruck(cd, x, y, cap uint, name string) *Truck {
	return &Truck{
		Status:   WAITING,
		Packages: nil,
		Capacity: cap,
		Pos: Position{
			x: x,
			y: y,
		},
		Cooldown:        cd,
		CurrentCooldown: cd,
		Name:            name,
	}
}

func (t *Truck) updateStatus(s TStatus) {
	t.Status = s
}

func (t *Truck) loadPackage(p *Package) {
	t.Packages = append(t.Packages, p)
}

func (t Truck) IsFull() (uint, bool) {
	var currentCap uint
	for _, p := range t.Packages {
		currentCap = currentCap + p.Weight
	}
	return currentCap, currentCap < t.Capacity
}

func (t *Truck) updateCD() {
	if t.CurrentCooldown > 0 {
		t.CurrentCooldown--
	} else {
		t.CurrentCooldown = t.Cooldown
	}
}

func (t Truck) Dump() {
	fmt.Printf("%+v\n\n", t)
}
