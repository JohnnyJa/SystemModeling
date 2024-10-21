package Processes

import (
	"Model/Lab2/Marker"
	"Model/Lab2/Queue"
	"Model/Lab2/Transitions"
	"fmt"
	"math"
)

type State int

const (
	Free State = iota
	Busy
)

type Processor struct {
	markerInProcess    *Marker.Marker
	nextActivationTime float64
	state              State
	currentTime        float64
	delay              float64
}

func (p *Processor) ProcessMarker(marker *Marker.Marker) {
	p.state = Busy
	p.markerInProcess = marker
	p.nextActivationTime = p.currentTime + p.delay
}

func (p *Processor) FinishProcess(transition *Transitions.Transition) {
	transition.PushMarkerToNextNode(p.markerInProcess)
}

type ProcessStatistic struct {
	totalCount   int
	failureCount int
	meanLoad     float64

	lastTime float64
}

func (p *ProcessStatistic) SetMeanLoad(state int, currentTime float64) {
	p.meanLoad += float64(state) * (currentTime - p.lastTime)
}

func (p *ProcessStatistic) GetMeanLoad(currentTime float64) float64 {
	return p.meanLoad / currentTime
}

func (p *ProcessStatistic) GetLog() string {
	return fmt.Sprintf("Total count: %d, Failure count: %d", p.totalCount, p.failureCount)
}

type Process struct {
	BaseElement
	ProcessStatistic

	transition     *Transitions.Transition
	queue          *Queue.Queue
	maxQueueSize   int
	numOfProcesses int
	processors     []*Processor
}

func NewProcess(id int, name string, maxQueueSize int, delay float64, numOfProcesses int) *Process {
	processors := make([]*Processor, numOfProcesses)

	for i := range processors {
		processors[i] = &Processor{
			markerInProcess:    nil,
			nextActivationTime: math.MaxFloat64,
			state:              Free,
			delay:              delay,
			currentTime:        0,
		}
	}

	return &Process{
		BaseElement: BaseElement{
			id,
			name,
			0,
			0,
		},
		queue:          Queue.NewQueue(),
		maxQueueSize:   maxQueueSize,
		numOfProcesses: numOfProcesses,
		processors:     processors,
	}
}

func (p *Process) RunToCurrentTime(currentTime float64) {

	for _, pr := range p.processors {
		if pr.state == Busy && pr.nextActivationTime <= currentTime {
			pr.FinishProcess(p.transition)
			p.totalCount++

			if p.queue.Size() > 0 {
				marker := p.queue.Pop()
				pr.ProcessMarker(marker)
			} else {
				pr.state = Free
				pr.nextActivationTime = math.MaxFloat64
			}

		}

		pr.currentTime = currentTime
	}

	p.DoStatistics(currentTime)
}

func (p *Process) TakeMarker(marker *Marker.Marker) {

	wasFree := false
	for _, processor := range p.processors {
		if processor.state == Free {
			processor.ProcessMarker(marker)
			wasFree = true
			break
		}
	}

	if !wasFree {
		if p.queue.Size() < p.maxQueueSize {
			p.queue.Push(marker)
		} else {
			p.failureCount++
		}
	}
}

func (p *Process) SetTransition(transition *Transitions.Transition) {
	p.transition = transition
}

func (p *Process) GetLog() string {
	processorsLog := ""
	for i, pr := range p.processors {
		processorsLog += fmt.Sprintf("\n  +++++++++++++++++\n   Processor %d: State: %d\n  +++++++++++++++++", i, pr.state)
	}

	return fmt.Sprintf("%s,\n Stats: %s\n %s", p.BaseElement.GetLog(), processorsLog, p.ProcessStatistic.GetLog())
}

func (p *Process) GetResults() string {
	return fmt.Sprintf("Mean load: %f\n", p.GetMeanLoad(p.processors[0].currentTime))
}

func (p *Process) DoStatistics(time float64) {

	for _, pr := range p.processors {
		p.SetMeanLoad(int(pr.state), time)
	}
	p.lastTime = time
}

func (p *Process) GetActivationTime() float64 {
	minTime := math.MaxFloat64
	for _, processor := range p.processors {
		if processor.nextActivationTime < minTime {
			minTime = processor.nextActivationTime
		}
	}

	return minTime
}
