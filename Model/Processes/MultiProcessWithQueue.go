package Processes

import (
	ModelQueue "Model/Model/Queue"
	"Model/Model/Statistic"
	"fmt"
)

type MultiProcessWithQueue struct {
	*MultiProcess
	*ModelQueue.Queue
	*Statistic.QueueStatistic
}

func NewMultiProcessWithQueue(numOfProcessors int, delay float64, maxQueueSize int) *MultiProcessWithQueue {
	return &MultiProcessWithQueue{
		MultiProcess:   NewMultiProcessWithDelay(numOfProcessors, delay),
		Queue:          ModelQueue.NewQueue(maxQueueSize),
		QueueStatistic: Statistic.NewQueueStatistic(),
	}
}

func (p *MultiProcessWithQueue) Start() {
	if p.StartInProcessor() != nil {
		if p.AddToQueue() != nil {
			p.AddFailure()
		}
	}
}

func (p *MultiProcessWithQueue) MoveToCurrentTime() {
	p.MultiProcess.MoveToCurrentTime()

	for p.GetCurrentQueueSize() > 0 {
		if p.StartInProcessor() != nil {
			break
		} else {
			p.RemoveFromQueue()
		}
	}

	p.QueueStatistic.CountMeanQueue(p.GetCurrentQueueSize(), p.processors[0].GetCurrentTime())
}

func (p *MultiProcessWithQueue) GetLog() string {
	return p.MultiProcess.GetLog() + fmt.Sprintf("Queue size: %d\nFailure: %d\n", p.GetCurrentQueueSize(), p.GetFailure())

}

func (p *MultiProcessWithQueue) GetResult() string {
	return p.MultiProcess.GetResult() + p.QueueStatistic.GetResult() + "\n"
}
