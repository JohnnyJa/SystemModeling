package Lab3

import (
	Model2 "Model/Model"
	"Model/Model/Interfaces"
	"fmt"
	"math"
)

type Model struct {
	list  []Interfaces.IProcess
	tnext float64
	tcurr float64
	event int
}

func NewModel(elements []Interfaces.IProcess) *Model {
	return &Model{
		list:  elements,
		tnext: 0.0,
		tcurr: 0.0,
		event: 0,
	}
}

func (m *Model) Simulate(time float64) {

	for m.tcurr < time {
		m.tnext = math.MaxFloat64
		for _, e := range m.list {
			if e.GetTNext() < m.tnext {
				m.tnext = e.GetTNext()
				m.event = e.GetId()
			}
		}

		for _, e := range m.list {
			e.DoStatistic(m.tnext - m.tcurr)
		}

		m.tcurr = m.tnext

		fmt.Printf("*********************************************\nCurrent Model time = %f\n-------------------\n", m.tcurr)

		for _, e := range m.list {
			e.SetTCurr(m.tcurr)
		}

		m.list[m.event].OutAct()
		for _, e := range m.list {
			if e.GetTNext() == m.tcurr {
				e.OutAct()
			}
		}
		m.PrintInfo()
	}
	m.PrintResult()
}

func (m *Model) PrintInfo() {
	for _, e := range m.list {
		e.PrintInfo()
	}
}

func (m *Model) PrintResult() {
	fmt.Println("Result:")
	for _, e := range m.list {
		e.PrintResult()
		if p, ok := e.(*Model2.Process); ok {
			fmt.Printf(" mean length of queue = %f\n failure probability = %f\n---------------------------------------------\n", p.GetMeanQueue()/m.tcurr, float64(p.GetFailure())/float64(p.GetFailure()+p.GetQuantity()))
		}
	}
}
