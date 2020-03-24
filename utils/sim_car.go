package utils

/*
@author '彼时思默'
@time 2020/3/19 8:59
@describe:
*/

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Simulator interface {
	Simulate(s Simulator, travelNum int)
	emit(s Simulator, travelNum int)
	duRule()
}

type CarSim struct {
	Vin   string
	CurTS int64
	wtP   *sync.WaitGroup
	chP   *chan DateTravel
	duP   *DateUnit
}

func NewCarSim(vin string, wtP *sync.WaitGroup, chP *chan DateTravel) Simulator {
	rand.Seed(time.Now().Unix())
	du := new(DateUnit)
	du.Vin = vin
	du.Ts = 0      //时间戳
	du.Mileage = 0 // 里程
	du.Speed = 0   // 速度
	du.VAcc = 0    // 纵加
	du.LAcc = 0    // 横加
	du.Fuel = 0    // 油耗
	c := CarSim{
		Vin:   vin,
		CurTS: 0,
		wtP:   wtP,
		chP:   chP,
		duP:   du,
	}
	v := strings.Split(vin, "_")[0]
	switch v {
	case "normal":
		return &CarNormal{
			c,
		}
	case "speedFast":
		return &CarSpeedFast{
			c,
		}
	case "unstable":
		return &CarUnstable{
			c,
		}
	case "energyWaste":
		return &CarEnergyWaste{
			c,
		}
	case "far":
		return &CarFar{
			c,
		}
	default:
		fmt.Println("Type error!")
		return c
	}
}

func (c CarSim) Simulate(s Simulator, travelNum int) {
	c.wtP.Add(1)
	go c.emit(s, travelNum)
}

func (c CarSim) emit(s Simulator, travelNum int) {
	defer c.wtP.Done()
	for r := 0; r < travelNum; r++ {
		unitNum := rand.Intn(80) + 20
		dt := make(DateTravel, unitNum)
		for i := 0; i < unitNum; i++ {
			s.duRule()
			dt[i] = *c.duP
		}
		*c.chP <- dt
	}
}

func (c CarSim) duRule() {
}

type CarNormal struct {
	CarSim
}

func (c CarNormal) duRule() {
	c.duP.Ts += 10 // 时间戳
	c.duP.Mileage += c.duP.Speed / 360
	c.duP.Speed = RandFloat(1, 60)
	c.duP.VAcc = RandFloat(-3, 3)
	c.duP.LAcc = RandFloat(-3, 3)
	c.duP.Fuel = RandFloat(5, 8)
}

type CarSpeedFast struct {
	CarSim
}

func (c CarSpeedFast) duRule() {
	c.duP.Ts += 10 // 时间戳
	c.duP.Mileage += c.duP.Speed / 360
	c.duP.Speed = RandFloat(40, 120)
	c.duP.VAcc = RandFloat(-5, 5)
	c.duP.LAcc = RandFloat(-5, 5)
	c.duP.Fuel = RandFloat(6, 11)
}

type CarUnstable struct {
	CarSim
}

func (c CarUnstable) duRule() {
	c.duP.Ts += 10 // 时间戳
	c.duP.Mileage += c.duP.Speed / 360
	c.duP.Speed = RandFloat(1, 60)
	c.duP.VAcc = RandFloat(-10, 10)
	c.duP.LAcc = RandFloat(-10, 10)
	c.duP.Fuel = RandFloat(5, 9)
}

type CarEnergyWaste struct {
	CarSim
}

func (c CarEnergyWaste) duRule() {
	c.duP.Ts += 10 // 时间戳
	c.duP.Mileage += c.duP.Speed / 360
	c.duP.Speed = RandFloat(1, 60)
	c.duP.VAcc = RandFloat(-3, 3)
	c.duP.LAcc = RandFloat(-3, 3)
	c.duP.Fuel = RandFloat(6, 14)
}

type CarFar struct {
	CarSim
}

func (c CarFar) duRule() {
	c.duP.Ts += 10 // 时间戳
	if c.duP.Mileage == 0 {
		c.duP.Mileage = 10000
	}
	c.duP.Mileage += c.duP.Speed / 360
	c.duP.Speed = RandFloat(40, 80)
	c.duP.VAcc = RandFloat(-3, 3)
	c.duP.LAcc = RandFloat(-3, 3)
	c.duP.Fuel = RandFloat(5, 8)
}
