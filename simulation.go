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
		for i, _ := range w.Trucks {
			w.TruckSimulation(&w.Trucks[i])
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
	if f.IsTargetSelected() == false && f.Package == nil {
		w.SelectForkliftObjective(f)
	}
	if f.Pos.x == f.TargetPos.x && f.Pos.y == f.TargetPos.y {
		if f.Package != nil {
			f.updateStatus("LEAVE")
		} else if f.Status == "LEAVE" {
			f.Package = nil
			// charge truck
			f.updateStatus("WAITING")
		}
	}
}

func (w *Warehouse) TruckSimulation(t *Truck) {
	if t.Status == "GONE" {
		t.empty()
		t.updateCD()
		return
	}
	packageCanFit := false
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
