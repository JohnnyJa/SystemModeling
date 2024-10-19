package Model

import (
	"fmt"
	"math"
)

type Process struct {
	element    *Element
	processors []*Element
	queue      int
	maxQueue   int
	failure    int
	meanQueue  float64

	condition ICondition
}

func NewProcess(delay float64, numOfProcessors int) *Process {
	elements := make([]*Element, numOfProcessors)
	for i := range elements {
		elements[i] = NewElementWithDelay(delay)
	}

	c := &Process{
		element:    NewElementWithDelay(0),
		processors: elements,
		queue:      0,
		maxQueue:   math.MaxInt,
		meanQueue:  0.0,
		condition:  NewFullRandomCondition(1),
	}
	c.element.tnext = math.MaxFloat64
	return c
}

func (c *Process) InAct() {
	isFreeProcess := false

	for _, el := range c.processors {
		if el.state == 0 {
			el.state = 1
			el.tnext = c.element.tcurr + el.GetDelay()
			isFreeProcess = true

			break
		}
	}

	c.element.tnext = math.MaxFloat64
	for _, el := range c.processors {
		if el.tnext < c.element.tnext {
			c.element.tnext = el.tnext
		}
	}

	if !isFreeProcess {
		if c.queue < c.maxQueue {
			c.queue++
		} else {
			c.failure++
		}
	}
}

func (c *Process) SingleInAct() {
	if c.element.state == 0 {
		c.element.state = 1
		c.element.tnext = c.element.tcurr + c.element.delayMean
	} else {
		if c.queue < c.maxQueue {
			c.queue++
		} else {
			c.failure++
		}
	}
}

func (c *Process) OutAct() {

	for _, el := range c.processors {
		if el.tnext == c.element.tcurr {
			el.OutAct()
			el.tnext = math.MaxFloat64
			el.state = 0
			if c.queue > 0 {
				c.queue--
				el.state = 1
				el.tnext = c.element.tcurr + el.GetDelay()

			}

			if c.element.nextElements != nil {
				c.element.nextElements[c.condition.MakeCondition()].InAct()
			}
		}
	}

	c.element.tnext = math.MaxFloat64
	for _, el := range c.processors {
		if el.tnext < c.element.tnext {
			c.element.tnext = el.tnext
		}
	}
}

func (c *Process) SingleOutAct() {
	c.element.OutAct()
	c.element.tnext = math.MaxFloat64
	c.element.state = 0

	if c.queue > 0 {
		c.queue--
		c.element.state = 1
		c.element.tnext = c.element.tcurr + c.element.delayMean
	}

	if c.element.nextElements != nil {
		c.element.nextElements[c.condition.MakeCondition()].InAct()
	}
}

func (c *Process) PrintInfo() {
	states := "["
	c.element.quantity = 0

	for _, el := range c.processors {
		if el.state == 0 {
			states += "0"
		} else {
			states += "1"
		}
		states += ", "
		c.element.quantity += el.quantity
	}

	states = states[:len(states)-2]
	states += "]"

	if c.element.tnext == math.MaxFloat64 {
		fmt.Printf("Element %s:\n tnext = max\n tcurr = %f\n queue = %d\n state = %s\n quantity = %d\n failure = %d\n-------------------\n", c.element.name, c.element.tcurr, c.queue, states, c.element.quantity, c.failure)

	} else {
		fmt.Printf("Element %s:\n tnext = %f\n tcurr = %f\n queue = %d\n state = %s\n quantity = %d\n failure = %d\n-------------------\n", c.element.name, c.element.tnext, c.element.tcurr, c.queue, states, c.element.quantity, c.failure)
	}
}

func (c *Process) DoStatistic(delta float64) {
	c.meanQueue += float64(c.queue) * delta
}

func (c *Process) GetBaseElement() *Element {
	return c.element
}
func (c *Process) PrintResult() {
	c.element.PrintResult()
}

func (c *Process) SetMaxQueue(maxQueue int) {
	c.maxQueue = maxQueue

}

func (c *Process) GetId() int {
	return c.element.id
}

func (c *Process) SetName(name string) {
	c.element.name = name
}

func (c *Process) SetDistribution(distribution string) {
	c.element.distribution = distribution
}

func (c *Process) SetNextElement(el IElement) {
	c.element.nextElements = append(c.element.nextElements, el)
}

func (c *Process) SetNextElements(elements []IElement) {
	c.element.nextElements = elements
}

func (c *Process) SetCondition(condition ICondition) {
	c.condition = condition
}

func (c *Process) GetQueueSize() int {
	return c.queue
}

func (c *Process) GetNextElements() []IElement {
	return c.element.nextElements
}

func (c *Process) GetTNext() float64 {
	return c.element.GetTNext()
}

func (c *Process) SetTNext(tNext float64) {
	c.element.SetTNext(tNext)
}

func (c *Process) GetTCurr() float64 {
	return c.element.GetTCurr()
}

func (c *Process) SetTCurr(tCurr float64) {
	c.element.SetTCurr(tCurr)
}

func (c *Process) GetMeanQueue() float64 {
	return c.meanQueue
}

func (c *Process) GetFailure() int {
	return c.failure
}

func (c *Process) GetQuantity() int {
	return c.element.quantity
}
