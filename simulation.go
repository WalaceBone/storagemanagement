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
		w.decountLifeTime()
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
	t.Dump()
}
