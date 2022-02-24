package main

import (
	"fmt"
	"time"
)

//Simulation Simulate the warehouse and it's actions
func (w *Warehouse) Simulation() error {

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	for w.IsSimulationComplete() == false {
		for i, f := range w.Forklifts {
			w.SelectForkliftObjective(&w.Forklifts[i])
			node := w.FindPath(w.Map[f.Pos.x][f.Pos.y].ID, w.Map[w.Forklifts[i].TargetPos.x][w.Forklifts[i].TargetPos.y].ID)
			fmt.Println(node)
			//w.move(r.Intn(4), &w.Forklifts[i])
		}
		for _, t := range w.Trucks {
			w.TruckSimulation(&t)
		}
		w.decountLifeTime()
		w.DumpTurn()
		w.DumpMap()
		fmt.Printf("\n")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (w *Warehouse) ForkliftSimulation(f *Forklift) {
	f.Dump()
}

func (w *Warehouse) TruckSimulation(t *Truck) {
	packageCanFit := false
	if t.Status == "GONE" {
		t.empty()
		t.updateStatus("WAITING")
		return
	}
	for _, p := range w.Packages {
		if t.CanReceive(p.Weight) == true {
			packageCanFit = true
		}
	}
	load, _ := t.IsFull()
	if load > 0 && packageCanFit == false {
		t.updateStatus("GONE")
	}
}
