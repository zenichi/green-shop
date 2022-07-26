package utils

import "math"

func RoundPrice(p float64) float64 {
	return math.Round(p*100) / 100
}
