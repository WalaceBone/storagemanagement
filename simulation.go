package main

import (
	"fmt"
)

//Simulation Simulate the warehouse and it's actions
func (w *Warehouse) Simulation() error {

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	for w.IsSimulationComplete() == false {
		for i, _ := range w.Forklifts {
			if w.PackageLeft() > 0 {
				w.ForkliftSimulation(&w.Forklifts[i])
			}
		}
		for i, _ := range w.Trucks {
			w.TruckSimulation(&w.Trucks[i])
		}
		w.decountLifeTime()
		w.DumpTurn()
		w.DumpMap()
		fmt.Printf("\n")
		//time.Sleep(1 * time.Second)
	}
	return nil
}

func (w *Warehouse) ForkliftSimulation(f *Forklift) {
	fmt.Println(w.PackageTargeted(), f.IsTargetSelected())
	//TODO
	// Trouver comme dire a une forklift de rien faire si pas de package dispo a prendre aka + de fork que de paquet
	//Forklift ended action
	//fmt.Println(f)
	if f.IsTargetSelected() == false && f.Status != LEAVE {
		w.SelectForkliftTarget(f)
	}
	if w.PackageTargeted() == len(w.Packages) && !f.IsTargetSelected() {
		return
	}
	// Is at Target
	if f.Path == nil && f.Status != LEAVE {
		path := w.FindPath(f.ID, f.Target)
		f.AddPath(path)
	}
	// Forklift is at truck
	if f.Package != nil && f.Pos.x == f.TargetPos.x && f.Pos.y == f.TargetPos.y && w.Map[f.TargetPos.x][f.TargetPos.y].T != nil {
		if f.Package != nil && w.Map[f.TargetPos.x][f.TargetPos.y].T.CanReceive(f.Package.Weight) {
			f.ResetPath()
			f.ResetTarget()
			f.updateStatus(LEAVE)
		}
	} else if f.Status == LEAVE {
		w.GetPackageByID(f.Package.ID).Load()
		w.GetCellById(f.Target).T.loadPackage(f.Package)
		f.Reset()
	}
	//TODO
	// add check can move else calc path
	if f.Path != nil && len(f.Path) > 0 {
		if f.Package == nil && len(f.Path) == 1 {
			f.Reset()
			f.updateStatus(TAKE)
			if w.GetCellById(f.Target).P != nil {
				f.Package = w.GetCellById(f.Target).GetPackage()
				w.GetCellById(f.ID).P = nil
			}
		} else {
			f.updateStatus(GO)
			oldCell := f.ID
			f.Move()
			f.Pos = w.GetCellById(f.ID).p
			w.GetCellById(oldCell).F = nil
			w.GetCellById(f.ID).F = f
		}
	}
}

func (w *Warehouse) TruckSimulation(t *Truck) {
	if t.Status == GONE {
		t.empty()
		t.updateCD()
		return
	}
	packageCanFit := false
	for _, p := range w.Packages {
		if t.CanReceive(p.Weight) == true && !p.Loaded {
			packageCanFit = true
		}
	}
	load, _ := t.IsFull()
	if load > 0 && packageCanFit == false {
		t.updateStatus(GONE)
	}
}
