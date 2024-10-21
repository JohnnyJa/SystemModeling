package Elements

import (
	"Model/Model/Statistic"
	"Model/funRand"
	"fmt"
	"math"
	"strings"
)

type IProcessElement interface {
	IElement
	ITimeElement
}

type ProcessElement struct {
	*BasicElement
	*TimeElement
	delayMean        float64
	delayDev         float64
	distributionType string
	state            ElementState
	marker           Statistic.Marker
}

func NewProcessElementWithDelay(delay float64) *ProcessElement {
	el := &ProcessElement{
		BasicElement:     NewBasicElement(),
		TimeElement:      NewTimeElement(),
		delayMean:        delay,
		distributionType: "",
		state:            0,
	}
	NumOFElements++

	return el
}

func NewAnonymusProcessElementWithDelay(delay float64) *ProcessElement {
	el := &ProcessElement{
		BasicElement:     NewAnonymusElement(),
		TimeElement:      NewTimeElement(),
		delayMean:        delay,
		distributionType: "",
		state:            0,
	}

	return el
}

func (e *ProcessElement) GetDelay() float64 {
	delay := e.delayMean
	if strings.EqualFold(e.distributionType, "exp") {
		delay = funRand.Exp(e.delayMean)
	} else if strings.EqualFold(e.distributionType, "unif") {
		delay = funRand.Unif(e.delayMean, e.delayDev)
	} else if strings.EqualFold(e.distributionType, "norm") {
		delay = funRand.Norm(e.delayMean, e.delayDev)
	} else if strings.EqualFold(e.distributionType, "") {
		delay = e.delayMean
	}

	return delay
}

func (e *ProcessElement) Start(marker Statistic.Marker) {
	e.marker = marker
	e.state = Busy
	e.nextTime = e.currentTime + e.GetDelay()
}

func (e *ProcessElement) Finish() {
	e.state = Free
	e.nextTime = math.MaxFloat64
}

func (e *ProcessElement) MoveToCurrentTime() {
	if e.currentTime >= e.nextTime {
		e.Finish()
	}
}

func (e *ProcessElement) GetLog() string {
	if e.nextTime == math.MaxFloat64 {
		return fmt.Sprintf("Element %s:\n tnext = max\n tcurr = %f\n state = %d\n-------------------\n", e.name, e.currentTime, e.state)
	} else {
		return fmt.Sprintf("Element %s:\n tnext = %f\n tcurr = %f\n state = %d\n-------------------\n", e.name, e.nextTime, e.currentTime, e.state)
	}
}

func (e *ProcessElement) SetDistributionType(distributionType string) {
	e.distributionType = distributionType
}

func (e *ProcessElement) GetState() ElementState {
	return e.state
}

func (e *ProcessElement) SetState(state ElementState) {
	e.state = state
}

func (e *ProcessElement) GetMarker() Statistic.Marker {
	return e.marker
}
