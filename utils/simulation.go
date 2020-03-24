package utils

/*
@author '彼时思默'
@time 2020/3/17 13:35
@describe:
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Simulation struct {
	CarLabel []string
	StartTS  int64
	wt       *sync.WaitGroup
	ch       chan DateTravel
}

func NewSimulation(wt *sync.WaitGroup) *Simulation {
	rand.Seed(time.Now().Unix())
	return &Simulation{
		CarLabel: []string{"normal", "speed_fast", "unstable", "energy_waste", "far"},
		//CarLabel:  []string{"normal"},
		StartTS: time.Now().Unix(),
		wt:      wt,
		ch:      make(chan DateTravel, 1e5),
	}
}

func (s Simulation) Simulate(carNum int, customNum int) {
	for i := 0; i < customNum; i++ {
		s.wt.Add(1)
		go s.sub()
	}
	for i := 0; i < carNum; i++ {
		for _, label := range s.CarLabel {
			s.wt.Add(1)
			go s.emit(label + "_" + Str(i))
		}
	}
}

func (s Simulation) sub() {
	defer s.wt.Done()
	url := "http://localhost:1234/api/car_report/v1.0/"
	for data := range s.ch {
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println("序列化错误")
		}
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		status := resp.Status
		resp.Body.Close()
		fmt.Println(Str(len(s.ch))+":\tresponse Status:", status)
		if strings.Contains(status, "500") {
			fmt.Println(data)
		}
	}
}

func (s Simulation) emit(label string) {
	//defer fmt.Println(label + "_ok")
	defer s.wt.Done()
	trTotal := 2 //单辆车产生多少次行程
	curTime := time.Now().Unix()
	curTS := curTime - s.StartTS
	num := rand.Intn(80) + 20
	du := new(DateUnit)
	du.Vin = label
	du.Ts = curTS  //时间戳
	du.Mileage = 0 // 里程
	du.Speed = 0   // 速度
	du.VAcc = 0    // 纵加
	du.LAcc = 0    // 横加
	du.Fuel = 0    // 油耗
	for r := 0; r < trTotal; r++ {
		travel := make(DateTravel, num)
		if strings.Contains(label, "normal") {
			for i := 0; i < num; i++ {
				du.Ts += 10 // 时间戳
				du.Mileage += du.Speed / 360
				du.Speed = RandFloat(1, 60)
				du.VAcc = RandFloat(-3, 3)
				du.LAcc = RandFloat(-3, 3)
				du.Fuel = RandFloat(5, 8)
				travel[i] = *du
			}
		} else if strings.Contains(label, "speed_fast") {
			for i := 0; i < num; i++ {
				du.Ts += 10                   // 时间戳
				du.Mileage += du.Speed / 360  // 里程
				du.Speed = RandFloat(40, 120) // 速度
				du.VAcc = RandFloat(-5, 5)    // 纵加
				du.LAcc = RandFloat(-5, 5)    // 横加
				du.Fuel = RandFloat(6, 11)    // 油耗
				travel[i] = *du
			}
		} else if strings.Contains(label, "unstable") {
			for i := 0; i < num; i++ {
				du.Ts += 10                  // 时间戳
				du.Mileage += du.Speed / 360 // 里程
				du.Speed = RandFloat(1, 60)  // 速度
				du.VAcc = RandFloat(-10, 10) // 纵加
				du.LAcc = RandFloat(-10, 10) // 横加
				du.Fuel = RandFloat(5, 9)    // 油耗
				travel[i] = *du
			}
		} else if strings.Contains(label, "energy_waste") {
			for i := 0; i < num; i++ {
				du.Ts += 10                  // 时间戳
				du.Mileage += du.Speed / 360 // 里程
				du.Speed = RandFloat(1, 60)  // 速度
				du.VAcc = RandFloat(-3, 3)   // 纵加
				du.LAcc = RandFloat(-3, 3)   // 横加
				du.Fuel = RandFloat(6, 14)   // 油耗
				travel[i] = *du
			}
		} else if strings.Contains(label, "far") {
			for i := 0; i < num; i++ {
				du.Ts += 10 // 时间戳
				if du.Mileage == 0 {
					du.Mileage = 10000
				}
				du.Mileage += du.Speed / 360 // 里程
				du.Speed = RandFloat(40, 80) // 速度
				du.VAcc = RandFloat(-3, 3)   // 纵加
				du.LAcc = RandFloat(-3, 3)   // 横加
				du.Fuel = RandFloat(5, 8)    // 油耗
				travel[i] = *du
			}
		} else {
			fmt.Printf("%#v\n", label)
		}
		//PrintA(travel)
		s.ch <- travel
		//time.Sleep(1)
		du.Ts += 100
	}
}
func RandInt(start int, stop int) int {
	return rand.Intn(stop-start) + start
}
func RandFloat(start int, stop int) float64 {
	result := rand.Float64()*float64(stop-start) + float64(start)
	return Round(result, 2)
}
func PrintS(arr interface{}) {
	switch it := arr.(type) {
	case DateTravel:
		for i := range it {
			fmt.Printf("%d: %#v\n", i, it[i])
		}
	case int:
		fmt.Printf("%d", it)
	}
}
