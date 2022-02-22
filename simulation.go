package main

import "time"

//Simulation Simulate the warehouse and it's actions
func Simulation(w *Warehouse) error {

	for w.IsSimulationComplete() == false {

		for _, fork := range w.Forklifts {
			go ForkliftSimulation(fork)
		}
		for _, truck := range w.Trucks {
			go TruckSimulation(truck)
		}

		w.decountLifeTime()
		w.DumpTurn()
		time.Sleep(500)
	}

	return nil
}

func ForkliftSimulation(f *Forklift) {
	f.Dump()
}

func TruckSimulation(t *Truck) {
	t.Dump()
}
