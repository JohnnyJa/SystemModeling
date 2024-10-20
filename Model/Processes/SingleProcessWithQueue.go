package Processes

import (
	"Model/Model/Elements"
	ModelQueue "Model/Model/Queue"
	"Model/Model/Statistic"
)

type SingleProcessWithQueue struct {
	*ModelQueue.Queue
	*Statistic.QueueStatistic
	*SingleProcess
}

func NewSingleProcessWithQueue(delay float64, maxQueueSize int) *SingleProcessWithQueue {
	return &SingleProcessWithQueue{
		Queue:         ModelQueue.NewQueue(maxQueueSize),
		SingleProcess: NewSingleProcess(delay),
	}
}

func (p *SingleProcessWithQueue) Start() {
	if p.GetState() == Elements.Free {
		p.SingleProcess.Start()
	} else {
		if err := p.AddToQueue(); err == nil {
			p.ElementStatistic.AddFailure()
		}
	}
}

func (p *SingleProcessWithQueue) Finish() {
	p.SingleProcess.Finish()
}

func (p *SingleProcessWithQueue) MoveToCurrentTime() {
	p.SingleProcess.MoveToCurrentTime()

	if p.GetCurrentQueueSize() > 0 {
		if p.SingleProcess.GetState() == Elements.Free {
			p.RemoveFromQueue()
			p.SingleProcess.Start()
		}

	}

	p.QueueStatistic.CountMeanQueue(p.GetCurrentQueueSize(), p.GetCurrentTime())
}
