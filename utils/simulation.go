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

//type DateUnit map[string]interface{}
type DateUnit struct {
	Vin string `json:"0000"`
	Ts int64 `json:"0001"`
	Mileage float64 `json:"2101"`
	Speed float64 `json:"2103"`
	VAcc float64 `json:"232F"`
	LAcc float64 `json:"2330"`
	Fuel float64 `json:"234E"`
}

type DateTravel []DateUnit

type Simulation struct {
	CarLabel  []string
	StartTime int64
	wt        *sync.WaitGroup
	ch        chan DateTravel
}

func NewSimulation(wt *sync.WaitGroup) *Simulation {
	rand.Seed(time.Now().Unix())
	return &Simulation{
		//CarLabel:  []string{"normal", "speed_fast", "unstable", "energy_waste", "far"},
		CarLabel:  []string{"normal"},
		StartTime: time.Now().Unix(),
		wt:        wt,
		ch:        make(chan DateTravel, 1024),
	}
}

func (s Simulation) Simulate(carNum int) {
	for i:=0;i<3;i++{
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
	jsonStr,err:=json.Marshal(<-s.ch)
	if err!=nil{
		fmt.Println("序列化错误")
	}
	req,_ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
}

func (s Simulation) emit(label string) {
	defer s.wt.Done()
	trTotal := 1
	curTime := time.Now().Unix()
	curTS := curTime - s.StartTime
	num := rand.Intn(80) + 20
	du := new(DateUnit)
	du.Vin=label
	du.Ts = curTS  //时间戳
	du.Mileage = 0 // 里程
	du.Speed = 0   // 速度
	du.VAcc = 0    // 纵加
	du.LAcc = 0    // 横加
	du.Fuel = 0  // 油耗

	for r := 0; r < trTotal; r++ {
		travel := make(DateTravel, 100)
		if strings.Contains(label, "normal") {
			for i := 0; i < num; i++ {
				du.Ts += 10 // 时间戳
				du.Mileage += du.Speed / 360
				du.Speed = RandFloat(1, 60)
				du.VAcc = RandFloat(-3, 3)
				du.LAcc = RandFloat(-3, 3)
				du.Fuel = Round(RandFloat(5,8), 2)
				travel = append(travel, *du)
			}
		} else if strings.Contains(label, "du.Speed_fast") {
			for i := 0; i < num; i++ {
				du.Ts += 10                              // 时间戳
				du.Mileage += du.Speed / 360                // 里程
				du.Speed = RandFloat(40, 120)              // 速度
				du.VAcc = RandFloat(-5, 5)                 // 纵加
				du.LAcc = RandFloat(-5, 5)                 // 横加
				du.Fuel = Round(RandFloat(6,11), 2)// 油耗
				travel = append(travel, *du)
			}
		} else if strings.Contains(label, "unstable") {
			for i := 0; i < num; i++ {
				du.Ts += 10                              // 时间戳
				du.Mileage += du.Speed / 360                // 里程
				du.Speed = RandFloat(1, 60)                // 速度
				du.VAcc = RandFloat(-10, 10)               // 纵加
				du.LAcc = RandFloat(-10, 10)               // 横加
				du.Fuel = Round(RandFloat(5,9), 2)  // 油耗
				travel = append(travel, *du)
			}
		} else if strings.Contains(label, "energy_waste") {
			for i := 0; i < num; i++ {
				du.Ts += 10                              // 时间戳
				du.Mileage += du.Speed / 360                // 里程
				du.Speed = RandFloat(1, 60)                // 速度
				du.VAcc = RandFloat(-3, 3)                 // 纵加
				du.LAcc = RandFloat(-3, 3)                 // 横加
				du.Fuel = Round(RandFloat(6,14), 2)// 油耗
				travel = append(travel, *du)
			}
		} else if strings.Contains(label, "far") {
			for i := 0; i < num; i++ {
				du.Ts += 10 // 时间戳
				if du.Mileage == 0 {
					du.Mileage = 10000
				}
				du.Mileage += du.Speed / 360                // 里程
				du.Speed = RandFloat(40, 80)               // 速度
				du.VAcc = RandFloat(-3, 3)                 // 纵加
				du.LAcc = RandFloat(-3, 3)                 // 横加
				du.Fuel = Round(RandFloat(5,8), 2) + 5 // 油耗
				travel = append(travel, *du)
			}
		}
		fmt.Printf("%#v", travel)
		s.ch <- travel
	}
}
func RandInt(start int, stop int) int {
	return rand.Intn(stop-start) + start
}
func RandFloat(start int, stop int) float64 {
	result:=rand.Float64()*float64(stop-start) + float64(start)
	return result
}
