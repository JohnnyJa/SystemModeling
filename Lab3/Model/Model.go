package Lab3

import (
	"Model/Model/Interfaces"
	"Model/Model/Processes"
	"fmt"
	"math"
)

var rebalance = 0

type Model struct {
	list        []Interfaces.IProcess
	currentTime float64
	event       int

	averageClientInBank float64
}

func NewModel(elements []Interfaces.IProcess) *Model {
	return &Model{
		list:                elements,
		currentTime:         0.0,
		event:               0,
		averageClientInBank: 0.0,
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
		prevTime := m.currentTime
		m.currentTime = nextTime

		for _, e := range m.list {
			e.SetCurrentTime(m.currentTime)
		}

		fmt.Printf("*********************************************\nCurrent Model time = %f\n-------------------\n", m.currentTime)

		for _, e := range m.list {
			e.MoveToCurrentTime()
		}
		m.PrintLog()

		m.RebalanceQueues(m.list)

		m.CountAverageClientsInBank(prevTime)
	}
	m.PrintResult()
}

func (m *Model) PrintLog() {
	for _, e := range m.list {
		fmt.Println(e.GetLog())
	}
}

func (m *Model) PrintResult() {
	log := "*********************************************\n"
	failures := m.list[1].(*Processes.SingleProcessWithQueue).GetFailure() + m.list[2].(*Processes.SingleProcessWithQueue).GetFailure()
	total := m.list[0].(*Processes.Create).GetTotalProceeded()
	for _, e := range m.list {
		log += fmt.Sprintf("-------------------\nElement %s:\n %s\n", e.GetName(), e.GetResult())
	}
	fmt.Println(log)

	fmt.Printf("Average clients in bank: %f\n", m.averageClientInBank/m.currentTime)
	fmt.Printf("Failures percent: %f\n", float64(failures)/float64(total+6)*float64(100))
	fmt.Printf("Rebalance: %d\n", rebalance)
}

func (m *Model) RebalanceQueues(list []Interfaces.IProcess) {
	firstQueue := list[1].(*Processes.SingleProcessWithQueue).GetQueue()
	secondQueue := list[2].(*Processes.SingleProcessWithQueue).GetQueue()

	if firstQueue.GetCurrentQueueSize()-secondQueue.GetCurrentQueueSize() >= 2 {
		m := firstQueue.RemoveFromQueue()
		_ = secondQueue.AddToQueue(m)
		fmt.Printf("Rebalanced queues: %d %d\n", firstQueue.GetCurrentQueueSize(), secondQueue.GetCurrentQueueSize())
		rebalance++
	}

	if secondQueue.GetCurrentQueueSize()-firstQueue.GetCurrentQueueSize() >= 2 {
		m := secondQueue.RemoveFromQueue()
		_ = firstQueue.AddToQueue(m)
		fmt.Printf("Rebalanced queues: %d %d\n", firstQueue.GetCurrentQueueSize(), secondQueue.GetCurrentQueueSize())
		rebalance++
	}

}

func (m *Model) CountAverageClientsInBank(time float64) {
	first := m.list[1].(*Processes.SingleProcessWithQueue).GetQueue().GetCurrentQueueSize()
	second := m.list[2].(*Processes.SingleProcessWithQueue).GetQueue().GetCurrentQueueSize()

	first += int(m.list[1].(*Processes.SingleProcessWithQueue).GetState())
	second += int(m.list[2].(*Processes.SingleProcessWithQueue).GetState())

	m.averageClientInBank += float64(first+second) * (m.currentTime - time)

}
