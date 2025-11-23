package util

import "strconv"

func parseUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
