package Model

import (
	"Model/funRand"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var nextId = 0

type IElement interface {
	OutAct()
	InAct()
	DoStatistic(float64)
	PrintInfo()
	PrintResult()
	GetNextElements() []IElement

	GetTNext() float64
	SetTNext(float64)

	GetTCurr() float64
	SetTCurr(float64)

	GetId() int
}

type Element struct {
	name         string
	tnext        float64
	delayMean    float64
	delayDev     float64
	distribution string
	quantity     int
	tcurr        float64
	state        int
	nextElements []IElement
	id           int
}

func NewBasicElement() *Element {
	el := &Element{
		tnext:        math.MaxFloat64,
		delayMean:    1.0,
		distribution: "exp",
		tcurr:        math.MaxFloat64,
		state:        0,
		id:           nextId,
		name:         "Element" + strconv.Itoa(nextId),
	}
	nextId++

	return el
}

func NewElementWithDelay(delay float64) *Element {
	el := &Element{
		name:         "anonymus",
		tnext:        math.MaxFloat64,
		delayMean:    delay,
		distribution: "",
		tcurr:        0.0,
		state:        0,
		nextElements: make([]IElement, 0),
		id:           nextId,
	}
	nextId++

	return el
}

func NewElement(name string, delay float64) *Element {
	el := &Element{
		name:         name,
		tnext:        0.0,
		delayMean:    delay,
		distribution: "exp",
		tcurr:        0.0,
		state:        0,
		nextElements: make([]IElement, 0),
		id:           nextId,
	}

	nextId++
	return el
}

func (e *Element) GetDelay() float64 {
	delay := e.delayMean
	if strings.EqualFold(e.distribution, "exp") {
		delay = funRand.Exp(e.delayMean)
	} else if strings.EqualFold(e.distribution, "unif") {
		delay = funRand.Unif(e.delayMean, e.delayDev)
	} else if strings.EqualFold(e.distribution, "norm") {
		delay = funRand.Norm(e.delayMean, e.delayDev)
	} else if strings.EqualFold(e.distribution, "") {
		delay = e.delayMean
	}

	return delay
}

func (e *Element) OutAct() {
	e.quantity++
}

func (e *Element) InAct() {

}

func (e *Element) PrintInfo() {
	if e.tnext == math.MaxFloat64 {
		fmt.Printf("Element %s:\n tnext = max\n tcurr = %f\n state = %d\n quantity = %d\n-------------------\n", e.name, e.tcurr, e.state, e.quantity)

	} else {
		fmt.Printf("Element %s:\n tnext = %f\n tcurr = %f\n state = %d\n quantity = %d\n-------------------\n", e.name, e.tnext, e.tcurr, e.state, e.quantity)
	}

}

func (e *Element) PrintResult() {
	fmt.Printf("Element %s:\n quantity = %d\n", e.name, e.quantity)
}

func (e *Element) GetTNext() float64 {
	return e.tnext
}

func (e *Element) SetTNext(tnext float64) {
	e.tnext = tnext
}

func (e *Element) GetTCurr() float64 {
	return e.tcurr
}

func (e *Element) SetTCurr(tcurr float64) {
	e.tcurr = tcurr
}

func (e *Element) GetId() int {
	return e.id
}
