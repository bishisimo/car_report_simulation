package utils

import "strconv"

func Str(a interface{}) string {
	switch v := a.(type) {
	case int:
		return strconv.Itoa(v)
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case uint:
		return strconv.Itoa(int(v))
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 32)
	default:
		return "change to String error"
	}
}
