package funRand

import (
	"math"
	"math/rand"
)

func Exp(timeMean float64) float64 {
	a := 0.0
	for a == 0.0 {
		a = rand.Float64()
	}
	return -timeMean * math.Log(a)
}

func Unif(timeMin float64, timeMax float64) float64 {
	a := 0.0
	for a == 0.0 {
		a = rand.Float64()
	}

	return timeMin + a*(timeMax-timeMin)
}

func Norm(timeMean float64, timeDeviation float64) float64 {
	return timeMean + timeDeviation*(rand.NormFloat64())
}
