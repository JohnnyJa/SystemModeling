package Processes

import (
	"Model/funRand"
	"fmt"
)

type Distribution int

const (
	None Distribution = iota
	Exp
	Unif
	Norm
)

type BaseElement struct {
	Id   int
	Name string

	nextActivationTime float64
	currentTime        float64
	distributionType   Distribution
	delayMean          float64
	delayDev           float64
}

func (e *BaseElement) GetActivationTime() float64 {
	return e.nextActivationTime
}

func (e *BaseElement) GetLog() string {
	return fmt.Sprintf("Element: %s, Id: %d", e.Name, e.Id)
}

func (e *BaseElement) SetDistribution(distribution Distribution) {
	e.distributionType = distribution
}

func (e *BaseElement) SetDelay(delayMean float64, delayDev float64) {
	e.delayDev = delayDev
	e.delayMean = delayMean
}

func (e *BaseElement) GetDelay() float64 {
	delay := e.delayMean
	switch e.distributionType {
	case None:
		delay = e.delayMean
	case Exp:
		delay = funRand.Exp(e.delayMean)
	case Unif:
		delay = funRand.Unif(e.delayMean, e.delayDev)
	case Norm:
		delay = funRand.Norm(e.delayMean, e.delayDev)
	default:
		delay = e.delayMean
	}

	return delay
}

func (e *BaseElement) SetNextActivationTime(time float64) {
	e.nextActivationTime = time
}
