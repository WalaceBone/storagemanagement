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
	Capacity        int
	Pos             Position
	Cooldown        int
	CurrentCooldown int
	Name            string
}

func NewTruck(cd, x, y, cap int, name string) Truck {
	return Truck{
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

func (t *Truck) empty() {
	t.Packages = nil
}

func (t Truck) IsFull() (int, bool) {
	var currentCap int
	for _, p := range t.Packages {
		currentCap = currentCap + p.Weight
	}
	return currentCap, currentCap < t.Capacity
}

func (t Truck) CanReceive(weight int) bool {
	var total int
	for _, p := range t.Packages {
		total += p.Weight
	}
	return total + weight < t.Capacity
}

func (t *Truck) updateCD() {
	if t.CurrentCooldown > 0 {
		t.CurrentCooldown--
	} else {
		t.CurrentCooldown = t.Cooldown
		t.updateStatus("WAITING")
	}
}

func (t Truck) Dump() {
	fmt.Printf("%+v\n\n", t)
}
