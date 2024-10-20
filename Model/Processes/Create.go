package Processes

import (
	"Model/Model/Conditions"
	"Model/Model/Elements"
	Statistic2 "Model/Model/Statistic"
)

type Create struct {
	*Elements.ProcessElement
	*Statistic2.ElementStatistic
	*Conditions.Transition
}

//
//type Create struct {
//	element *Element
//
//	condition ICondition
//}

func NewCreate(delay float64) *Create {
	c := &Create{
		ProcessElement:   Elements.NewProcessElementWithDelay(delay),
		ElementStatistic: Statistic2.NewElementStatistic(),
		Transition:       Conditions.NewTransition(),
	}
	return c
}

func (c *Create) Finish() {
	c.AddTotalProceeded()

	c.SetNextTime(c.GetCurrentTime() + c.GetDelay())

	c.Transition.StartNextElement()
}

func (c *Create) Start() {

}

func (c *Create) MoveToCurrentTime() {
	if c.TimeElement.GetNextTime() <= c.GetCurrentTime() {
		c.Finish()
	}

}
