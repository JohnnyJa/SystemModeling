package Lab2

import (
	"Model/Model/Interfaces"
	"fmt"
	"math"
)

type Model struct {
	list        []Interfaces.IProcess
	currentTime float64
	event       int
}

func NewModel(elements []Interfaces.IProcess) *Model {
	return &Model{
		list:        elements,
		currentTime: 0.0,
		event:       0,
	}
}

func (m *Model) Simulate(time float64) {
	for m.currentTime < time {
		nextTime := math.MaxFloat64
		for _, e := range m.list {
			if e.GetNextTime() < nextTime {
				nextTime = e.GetNextTime()
				m.event = e.GetId()
			}
		}

		m.currentTime = nextTime

		for _, e := range m.list {
			e.SetCurrentTime(m.currentTime)
		}

		fmt.Printf("*********************************************\nCurrent Model time = %f\n-------------------\n", m.currentTime)

		for _, e := range m.list {
			e.MoveToCurrentTime()
		}
		m.PrintLog()
	}
	m.PrintResult()
}

func (m *Model) PrintLog() {
	for _, e := range m.list {
		fmt.Println(e.GetLog())
	}
}

func (m *Model) PrintResult() {
	log := ""
	for _, e := range m.list {
		log += e.GetResult() + "\n"
	}
	fmt.Println(log)
}
