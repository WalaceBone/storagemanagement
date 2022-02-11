package main

type TStatus string

const (
	GONE    TStatus = "GONE"
	WAITING TStatus = "WAITING"
)

type Truck struct {
	Status   TStatus
	Packages []Package
	Capacity uint
	Pos      Position
	Cooldown uint
	Name     string
}
