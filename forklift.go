package main

type FStatus string

const (
	GO    FStatus = "GO"
	WAIT  FStatus = "WAIT"
	TAKE  FStatus = "TAKE"
	LEAVE FStatus = "LEAVE"
)

type Forklift struct {
	Pos     Position
	Package Package
	Status  FStatus
	Name    string
}
