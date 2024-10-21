package Processes

import (
	"Model/Model/Elements"
	ModelQueue "Model/Model/Queue"
	"Model/Model/Statistic"
	"fmt"
)

type SingleProcessWithQueue struct {
	*ModelQueue.ModelQueue
	*Statistic.QueueStatistic
	*SingleProcess
}

func NewSingleProcessWithQueue(delay float64, maxQueueSize int) *SingleProcessWithQueue {
	return &SingleProcessWithQueue{
		ModelQueue:     ModelQueue.NewQueue(maxQueueSize),
		QueueStatistic: Statistic.NewQueueStatistic(),
		SingleProcess:  NewSingleProcess(delay),
	}
}

func (p *SingleProcessWithQueue) Start(marker Statistic.Marker) {
	if p.GetState() == Elements.Free {
		p.SingleProcess.Start(marker)
	} else {
		if err := p.AddToQueue(marker); err != nil {
			p.ElementStatistic.AddFailure()
		}
	}
}

func (p *SingleProcessWithQueue) Finish() {
	p.SingleProcess.Finish()
}

func (p *SingleProcessWithQueue) MoveToCurrentTime() {
	if p.GetCurrentTime() >= p.GetNextTime() {
		p.Finish()
	}

	p.SingleProcess.MoveToCurrentTime()
	if p.GetCurrentQueueSize() > 0 {
		if p.SingleProcess.GetState() == Elements.Free {
			m := p.RemoveFromQueue()
			p.SingleProcess.Start(m)
		}
	}

	p.QueueStatistic.CountMeanQueue(p.GetCurrentQueueSize(), p.GetCurrentTime())
	p.QueueStatistic.SetLastTime(p.GetCurrentTime())
}

func (p *SingleProcessWithQueue) GetQueue() *ModelQueue.ModelQueue {
	return p.ModelQueue
}

func (p *SingleProcessWithQueue) GetLog() string {
	return p.ProcessElement.GetLog() + fmt.Sprintf("ModelQueue size: %d\nFailure: %d\n", p.GetCurrentQueueSize(), p.GetFailure())
}

func (p *SingleProcessWithQueue) GetResult() string {
	return p.SingleProcess.GetResult() + p.QueueStatistic.GetResult() + "\n"
}
