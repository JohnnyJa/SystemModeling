package Lab2

import (
	"Model/Model/Interface"
	"fmt"
)

type Model struct {
	simulationTime float64
	elements       []Interface.IElement
	currentTime    float64
}

func NewModel(simulationTime float64, elements []Interface.IElement) *Model {
	return &Model{
		simulationTime: simulationTime,
		elements:       elements,
		currentTime:    0.0,
	}
}

func (m *Model) Simulate() {
	for m.currentTime < m.simulationTime {
		m.currentTime = m.FindNextActivationTime()
		m.RunToCurrentTime(m.currentTime)
		m.LogResults()
	}

	m.GetResults()

}

func (m *Model) FindNextActivationTime() float64 {
	minTime := m.simulationTime
	for _, element := range m.elements {
		if element.GetActivationTime() < minTime {
			minTime = element.GetActivationTime()
		}
	}
	return minTime
}

func (m *Model) RunToCurrentTime(time float64) {
	for _, element := range m.elements {
		element.RunToCurrentTime(time)
	}
}

func (m *Model) LogResults() {
	fmt.Printf("Current time: %f\n", m.FindNextActivationTime())

	for _, element := range m.elements {
		log := element.(Interface.ILogger).GetLog()
		fmt.Println("---------------------------")
		fmt.Println(log)
	}
	fmt.Println("******************************************************")

}

func (m *Model) GetResults() {
	for _, element := range m.elements {
		log := element.(Interface.ILogger).GetResults()
		fmt.Println(log)
	}
}
