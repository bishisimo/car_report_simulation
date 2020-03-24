package main

import (
	"fmt"
	"git/car_report_simulation/utils"
	"sync"
)

func app(wtP *sync.WaitGroup) {
	labels := []string{"normal", "speedFast", "unstable", "energyWaste", "far"}
	carNum := 100
	ch := make(chan utils.DateTravel, 1024)
	subscriber := utils.NewSubscriber(wtP, &ch)
	subscriber.Subscribe(16)
	for i := 0; i < carNum; i++ {
		for _, l := range labels {
			vin := l + "_" + utils.Str(i)
			cs := utils.NewCarSim(vin, wtP, &ch)
			cs.Simulate(cs, 2)
		}
	}
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	wt := sync.WaitGroup{}
	//simulation := utils.NewSimulation(&wt)
	//simulation.Simulate(1e4,16)
	app(&wt)
	wt.Wait()
	fmt.Println("Over!")
}
