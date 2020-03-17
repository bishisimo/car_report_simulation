package main

import (
	"git/car_report_simulation/utils"
	"sync"
)

func main() {
	wt:=sync.WaitGroup{}
	simulation := utils.NewSimulation(&wt)
	simulation.Simulate(1)
	wt.Wait()
}
