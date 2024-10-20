package Processes

import (
	"Model/Model/Conditions"
	"Model/Model/Elements"
	"Model/Model/Statistic"
	"errors"
	"fmt"
	"math"
)

type MultiProcess struct {
	*Elements.BasicElement
	processors []*Elements.ProcessElement
	*Statistic.ElementStatistic
	*Conditions.Transition
}

func NewMultiProcessWithDelay(numOfProcessors int, delay float64) *MultiProcess {
	pr := make([]*Elements.ProcessElement, numOfProcessors)
	for i := 0; i < numOfProcessors; i++ {
		pr[i] = Elements.NewAnonymusProcessElementWithDelay(delay)
	}

	return &MultiProcess{
		BasicElement:     Elements.NewBasicElement(),
		processors:       pr,
		ElementStatistic: Statistic.NewElementStatistic(),
		Transition:       Conditions.NewTransition(),
	}
}

func (p *MultiProcess) Start() {
	if err := p.StartInProcessor(); err != nil {
		p.AddFailure()
	}
}

func (p *MultiProcess) StartInProcessor() error {
	for _, pr := range p.processors {
		if pr.GetState() == Elements.Free {
			pr.Start()
			return nil
		}
	}

	return errors.New("all processors are busy")
}

func (p *MultiProcess) Finish() {
	p.AddTotalProceeded()
	p.Transition.StartNextElement()
}

func (p *MultiProcess) MoveToCurrentTime() {
	for _, pr := range p.processors {
		if (pr.GetNextTime() <= pr.GetCurrentTime()) && (pr.GetState() == Elements.Busy) {
			p.Finish()
		}

		pr.MoveToCurrentTime()
	}
}

func (p *MultiProcess) GetLog() string {

	log := fmt.Sprintf("MultiProcess Element %s:\n-------------------\n ", p.BasicElement.GetName())
	for i, pr := range p.processors {
		log += fmt.Sprintf("Processor %d:\n ", i)
		log += pr.GetLog()
	}
	return log
}

func (p *MultiProcess) GetNextTime() float64 {
	nextTime := math.MaxFloat64
	for _, pr := range p.processors {
		if pr.GetNextTime() < nextTime {
			nextTime = pr.GetNextTime()
		}
	}
	return nextTime
}

func (p *MultiProcess) SetCurrentTime(currentTime float64) {
	for _, pr := range p.processors {
		pr.SetCurrentTime(currentTime)
	}
}

func (p *MultiProcess) SetDistributionType(distributionType string) {
	for _, pr := range p.processors {
		pr.SetDistributionType(distributionType)
	}
}

func (p *MultiProcess) GetResult() string {
	return fmt.Sprintf("Element %s, result: %s\n", p.BasicElement.GetName(), p.ElementStatistic.GetResult())
}
