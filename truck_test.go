package main

import "testing"

func TestNewTruck(t *testing.T) {
	truck := NewTruck(1000, 100, 100, 1000, "truckname")
	testTruck := Truck{
		Status:   WAITING,
		Packages: make([]Package, 1),
		Capacity: 1000,
		Pos: Position{
			x: 100,
			y: 100,
		},
		Cooldown:        1000,
		CurrentCooldown: 1000,
		Name:            "truckname",
	}
	if truck.Name != testTruck.Name {
		t.Errorf("Name not equal")
	}
	if truck.Packages == nil || len(truck.Packages) != len(testTruck.Packages) {
		t.Errorf("Package not initalized")
	}
	if truck.Cooldown != testTruck.Cooldown {
		t.Errorf("Cooldown not equal")
	}
	if truck.CurrentCooldown != testTruck.CurrentCooldown {
		t.Errorf("CurrentCooldown not equal")
	}
	if truck.Capacity != testTruck.Capacity {
		t.Errorf("Capacity not equal")
	}
	if truck.Status != testTruck.Status {
		t.Errorf("Status not equal")
	}
	if truck.Pos.x != testTruck.Pos.x {
		t.Errorf("Pos not equal")
	}
	if truck.Pos.y != testTruck.Pos.y {
		t.Errorf("Pos not equal")
	}
}

func TestUpdateStatus(t *testing.T) {
	truck := NewTruck(1000, 100, 100, 1000, "truckname")
	truck.updateStatus(GONE)
	if truck.Status != GONE {
		t.Errorf("Status not set")
	}
}

func TestLoadPackage(t *testing.T) {

}