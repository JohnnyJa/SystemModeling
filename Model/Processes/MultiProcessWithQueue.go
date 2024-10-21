package Processes

import (
	ModelQueue "Model/Model/Queue"
	"Model/Model/Statistic"
	"fmt"
)

type MultiProcessWithQueue struct {
	*MultiProcess
	*ModelQueue.ModelQueue
	*Statistic.QueueStatistic
}

func NewMultiProcessWithQueue(numOfProcessors int, delay float64, maxQueueSize int) *MultiProcessWithQueue {
	return &MultiProcessWithQueue{
		MultiProcess:   NewMultiProcessWithDelay(numOfProcessors, delay),
		ModelQueue:     ModelQueue.NewQueue(maxQueueSize),
		QueueStatistic: Statistic.NewQueueStatistic(),
	}
}

func (p *MultiProcessWithQueue) Start(marker Statistic.Marker) {
	p.marker = marker

	if p.StartInProcessor(marker) != nil {
		if p.AddToQueue(marker) != nil {
			p.AddFailure()
		}
	}
}

func (p *MultiProcessWithQueue) MoveToCurrentTime() {
	p.MultiProcess.MoveToCurrentTime()
	nextMarker := p.GetFirst()
	for p.GetCurrentQueueSize() > 0 {
		if p.StartInProcessor(nextMarker) != nil {
			break
		} else {
			p.RemoveFromQueue()
		}
	}

	p.QueueStatistic.CountMeanQueue(p.GetCurrentQueueSize(), p.processors[0].GetCurrentTime())
	p.QueueStatistic.SetLastTime(p.processors[0].GetCurrentTime())
}

func (p *MultiProcessWithQueue) GetLog() string {
	return p.MultiProcess.GetLog() + fmt.Sprintf("ModelQueue size: %d\nFailure: %d\n", p.GetCurrentQueueSize(), p.GetFailure())

}

func (p *MultiProcessWithQueue) GetResult() string {
	return p.MultiProcess.GetResult() + p.QueueStatistic.GetResult() + "\n"
}
