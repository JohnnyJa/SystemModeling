package Processes

import (
	"Model/Model/Conditions"
	"Model/Model/Elements"
	Statistic2 "Model/Model/Statistic"
)

type SingleProcess struct {
	*Elements.ProcessElement
	*Statistic2.ElementStatistic
	*Conditions.Transition
}

func NewSingleProcess(delay float64) *SingleProcess {
	return &SingleProcess{
		ProcessElement:   Elements.NewProcessElementWithDelay(delay),
		ElementStatistic: Statistic2.NewElementStatistic(),
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
