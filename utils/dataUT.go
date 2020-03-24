/*
@author '彼时思默'
@time 2020/3/19 15:15
@describe:
*/
package utils

type DateUnit struct {
	Vin     string  `json:"0000"`
	Ts      int64   `json:"0001"`
	Mileage float64 `json:"2101"`
	Speed   float64 `json:"2103"`
	VAcc    float64 `json:"232F"`
	LAcc    float64 `json:"2330"`
	Fuel    float64 `json:"234E"`
}

type DateTravel []DateUnit
