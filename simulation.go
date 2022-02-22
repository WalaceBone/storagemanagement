package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Simulation Simulate the warehouse and it's actions
func (w *Warehouse) Simulation() error {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for w.IsSimulationComplete() == false {
		for i, _ := range w.Forklifts {
			w.move(r.Intn(4), &w.Forklifts[i])
		}
		w.decountLifeTime()
		w.DumpMap()
		fmt.Printf("\n")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (w *Warehouse) findNextPackage(f *Forklift) {

}

func (w *Warehouse) ForkliftSimulation(f *Forklift) {
	f.Dump()
}

func (w *Warehouse) TruckSimulation(t *Truck) {
	t.Dump()
}
