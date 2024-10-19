package main

import (
	"fmt"
	"math/rand"
	"time"
)

var data = [][2]float64{
	{0, 0},
	{2, 0.3},
	{4, 0.7},
	{6, 0.9},
	{10, 1.0},
}

func generateRandomValue(cdfValues []float64, xValues []float64) float64 {

	randomNumber := rand.Float64()

	for i := 0; i < len(cdfValues)-1; i++ {
		if cdfValues[i] <= randomNumber && randomNumber < cdfValues[i+1] {
			return xValues[i]
		}
	}

	return xValues[len(xValues)-1]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	xValues := make([]float64, len(data))
	cdfValues := make([]float64, len(data))

	for i, point := range data {
		xValues[i] = point[0]
		cdfValues[i] = point[1]
	}

	randomValues := make([]float64, 10)
	for i := 0; i < 10; i++ {
		randomValues[i] = generateRandomValue(cdfValues, xValues)
	}

	fmt.Println("Сгенеровані випадкові величини:", randomValues)
}
