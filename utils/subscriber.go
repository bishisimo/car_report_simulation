/*
@author '彼时思默'
@time 2020/3/19 9:03
@describe:
*/
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Subscriber struct {
	wtP *sync.WaitGroup
	chP *chan DateTravel
	url string
}

func NewSubscriber(wtP *sync.WaitGroup, chP *chan DateTravel) *Subscriber {
	return &Subscriber{
		wtP: wtP,
		chP: chP,
		url: "http://localhost:1234/api/car_report/v1.0/",
	}
}

func (s Subscriber) Subscribe(subNum int) {
	for i := 0; i < subNum; i++ {
		s.wtP.Add(1)
		go s._baseSub()
	}
}

func (s Subscriber) _baseSub() {
	defer s.wtP.Done()
	client := &http.Client{}
	for data := range *s.chP {
		//fmt.Println("data",data)
		//_=data
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println("序列化错误")
		}
		req, _ := http.NewRequest("POST", s.url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(req)
		status := resp.Status
		resp.Body.Close()
		fmt.Println(Str(len(*s.chP))+":\tresponse Status:", status)
		if strings.Contains(status, "500") {
			fmt.Println("_baseSub", data)
		}
	}
}
