package utils

import "strconv"

func ToInt(txt string) (int, error) {
	return strconv.Atoi(txt)
}

func ToInt64(txt string) (int64, error) {
	return strconv.ParseInt(txt, 10, 64)
}

func ToFloat64(txt string) (float64, error) {
	return strconv.ParseFloat(txt, 64)
}
