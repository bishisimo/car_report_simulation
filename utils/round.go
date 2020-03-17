package utils

import (
	"fmt"
	"strconv"
)

func Round(f interface{}, n int) float64 {
	switch f.(type) {
	case float32:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case float64:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case int:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case int8:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case int16:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case int32:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case int64:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	case uint:
		str := fmt.Sprintf("%."+strconv.Itoa(n)+"d", f)
		inst, _ := strconv.ParseFloat(str, 64)
		return inst
	default:
		return 0.0
	}
}