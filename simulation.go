package main

func Simulation(w *Warehouse) error {

	for w.IsSimulationComplete() == false {
		w.decountLifeTime()
		w.DumpTurn()
	}

	return nil
}
