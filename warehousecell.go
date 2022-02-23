package main

type WarehouseCell struct {
	p Position
	F *Forklift
	T *Truck
	P *Package
}

func (c WarehouseCell) IsEmpty() bool {
	return c.F != nil && c.T != nil && c.P != nil
}
