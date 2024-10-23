package Processes

import (
	"Model/Model/Marker"
	"fmt"
)

type BaseElement struct {
	Id   int
	Name string

	nextActivationTime float64
	currentTime        float64
	delayMean          float64
	delayDev           float64

	GetDelay func(marker *Marker.Marker) float64
}

func (e *BaseElement) GetActivationTime() float64 {
	return e.nextActivationTime
}

func (e *BaseElement) GetLog() string {
	return fmt.Sprintf("Element: %s, Id: %d", e.Name, e.Id)
}

func (e *BaseElement) SetDelay(delayMean float64, delayDev float64) {
	e.delayDev = delayDev
	e.delayMean = delayMean
}

func (e *BaseElement) SetNextActivationTime(time float64) {
	e.nextActivationTime = time
}
