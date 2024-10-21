package Processes

import (
	"Model/Model/Conditions"
	"Model/Model/Elements"
	"Model/Model/Statistic"
	"fmt"
)

type SingleProcess struct {
	*Elements.ProcessElement
	*Statistic.ElementStatistic
	*Conditions.Transition
}

func NewSingleProcess(delay float64) *SingleProcess {
	return &SingleProcess{
		ProcessElement:   Elements.NewProcessElementWithDelay(delay),
		ElementStatistic: Statistic.NewElementStatistic(),
		Transition:       Conditions.NewTransition(),
	}
}

func (p *SingleProcess) Start() {
	if p.GetState() == Elements.Busy {
		p.ElementStatistic.AddFailure()
		return
	}
	p.ProcessElement.Start()
}

func (p *SingleProcess) Finish() {
	p.ProcessElement.Finish()

	p.ElementStatistic.AddTotalProceeded()

	p.Transition.StartNextElement()
}

func (p *SingleProcess) GetResult() string {
	return fmt.Sprintf("Element %s, result: %s\n", p.BasicElement.GetName(), p.ElementStatistic.GetResult())
}
