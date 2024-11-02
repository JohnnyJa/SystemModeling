package Lab4

import (
	"Model/Model/Interface"
	"Model/Model/Marker"
	"Model/Model/Processes"
	"fmt"
)

type ModelStatistic struct {
	avgAmountOfClient float64
	rebalanced        int
}

type Model struct {
	ModelStatistic
	simulationTime float64
	elements       []Interface.IElement
	currentTime    float64
	prevTime       float64
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

		//m.LogResults()
		//m.RebalanceQueue()

		//m.CountStats()
	}

	//m.GetResults()

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
	fmt.Printf("Current time: %f\n", m.currentTime)

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

	fmt.Printf("Average number of clients: %f\n", m.avgAmountOfClient/m.simulationTime)
	fmt.Printf("Failure percent: %f\n", float64(m.elements[1].(*Processes.Process).FailureCount+m.elements[2].(*Processes.Process).FailureCount)/float64(m.elements[0].(*Processes.Create).Created))
	fmt.Printf("Rebalance: %d", m.rebalanced)
	//fmt.Printf("Average time between leaving: %f\n", m.elements[3].(*Processes.Dispose).GetLeaveTime())
}

func (m *Model) CountStats() {
	p1 := m.elements[1].(*Processes.Process)
	p2 := m.elements[2].(*Processes.Process)

	m.avgAmountOfClient += float64(p1.GetQueue().Size()+p1.GetNumOfMarkersInProcess()+p2.GetQueue().Size()+p2.GetNumOfMarkersInProcess()) * (m.currentTime - m.prevTime)
	m.prevTime = m.currentTime
}

func (m *Model) RebalanceQueue() {
	wq := m.elements[1].(*Processes.Process).GetQueue()
	wq.OrderQueue(func(marker *Marker.Marker, marker2 *Marker.Marker) bool {
		return marker.Type == 1
	})
	fmt.Printf("qq")
}
