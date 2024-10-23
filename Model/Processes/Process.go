package Processes

import (
	"Model/Model/Marker"
	"Model/Model/Queue"
	"Model/Model/Transitions"
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
}

func (p *Processor) ProcessMarker(marker *Marker.Marker, delay float64) {
	p.state = Busy
	p.markerInProcess = marker
	p.nextActivationTime = p.currentTime + delay
}

func (p *Processor) FinishProcess(transition *Transitions.Transition) {
	transition.PushMarkerToNextNode(p.markerInProcess)
}

type ProcessStatistic struct {
	totalCount   int
	FailureCount int
	meanLoad     float64
	meanQueue    float64
	enterDelay   float64

	lastTime      float64
	lastTimeEnter float64
}

func (p *ProcessStatistic) SetMeanLoad(state int, currentTime float64) {
	p.meanLoad += float64(state) * (currentTime - p.lastTime)
}

func (p *ProcessStatistic) GetMeanLoad(currentTime float64) float64 {
	return p.meanLoad / currentTime
}

func (p *ProcessStatistic) GetLog() string {
	return fmt.Sprintf("Total count: %d, Failure count: %d", p.totalCount, p.FailureCount)
}

type Process struct {
	BaseElement
	ProcessStatistic

	transition     *Transitions.Transition
	queue          *Queue.Queue
	maxQueueSize   int
	numOfProcesses int
	processors     []*Processor

	ChangeMarker func(marker *Marker.Marker)
}

func NewProcess(id int, name string, maxQueueSize int, numOfProcesses int) *Process {
	var processors []*Processor
	if numOfProcesses == -1 {
		processors = make([]*Processor, 0)

	} else {
		processors = make([]*Processor, numOfProcesses)

		for i := range processors {
			processors[i] = &Processor{
				markerInProcess:    nil,
				nextActivationTime: math.MaxFloat64,
				state:              Free,
				currentTime:        0,
			}
		}
	}

	return &Process{
		BaseElement: BaseElement{
			Id:                 id,
			Name:               name,
			currentTime:        0,
			nextActivationTime: math.MaxFloat64,
		},
		queue:          Queue.NewQueue(),
		maxQueueSize:   maxQueueSize,
		numOfProcesses: numOfProcesses,
		processors:     processors,
		ChangeMarker: func(marker *Marker.Marker) {
			// do nothing
		},
	}
}

func (p *Process) RunToCurrentTime(currentTime float64) {
	p.currentTime = currentTime
	if p.numOfProcesses == -1 {
		toDelete := make([]int, 0)
		for i, pr := range p.processors {
			if pr.state == Busy && pr.nextActivationTime <= currentTime {
				pr.FinishProcess(p.transition)
				p.totalCount++
				toDelete = append(toDelete, i)
			}
			pr.currentTime = currentTime
		}

		for i, j := 0, len(toDelete)-1; i < j; i, j = i+1, j-1 {
			toDelete[i], toDelete[j] = toDelete[j], toDelete[i]
		}

		for _, i := range toDelete {
			p.processors = append(p.processors[:i], p.processors[i+1:]...)
		}
		return
	}

	for _, pr := range p.processors {
		if pr.state == Busy && pr.nextActivationTime <= currentTime {
			pr.FinishProcess(p.transition)
			p.totalCount++

			if p.queue.Size() > 0 {
				marker := p.queue.Pop()
				pr.ProcessMarker(marker, p.GetDelay(marker))
				p.ChangeMarker(marker)
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
	p.enterDelay += p.currentTime - p.lastTimeEnter
	p.lastTimeEnter = p.currentTime

	if p.numOfProcesses == -1 {
		p.processors = append(p.processors, &Processor{
			markerInProcess:    marker,
			nextActivationTime: p.currentTime + p.GetDelay(marker),
			state:              Busy,
			currentTime:        p.currentTime,
		})

		return
	}

	wasFree := false
	for _, processor := range p.processors {
		if processor.state == Free {
			processor.ProcessMarker(marker, p.GetDelay(marker))
			p.ChangeMarker(marker)
			wasFree = true
			break
		}
	}

	if !wasFree {
		if p.queue.Size() < p.maxQueueSize {
			p.queue.Push(marker)
		} else {
			p.FailureCount++
		}
	}
}

func (p *Process) SetTransition(transition *Transitions.Transition) {
	p.transition = transition
}

func (p *Process) GetLog() string {
	processorsLog := ""
	for i, pr := range p.processors {
		if pr.nextActivationTime == math.MaxFloat64 {
			processorsLog += fmt.Sprintf("\n  +++++++++++++++++\n   Processor %d: State: %d Next activation at: never, current time: %f \n +++++++++++++++++", i, pr.state, pr.currentTime)
		} else {
			processorsLog += fmt.Sprintf("\n  +++++++++++++++++\n   Processor %d: State: %d Next activation at: %f, current time: %f \n +++++++++++++++++", i, pr.state, pr.nextActivationTime, pr.currentTime)
		}
	}

	return fmt.Sprintf("%s,\n Stats: Queue size: %d %s\n %s", p.BaseElement.GetLog(), p.queue.Size(), processorsLog, p.ProcessStatistic.GetLog())
}

func (p *Process) GetResults() string {
	return fmt.Sprintf("Element name %s\nStats:\n Mean load: %f\n Mean queue: %f\nEnter delay: %f",
		p.Name,
		p.GetMeanLoad(p.currentTime),
		p.meanQueue/p.currentTime,
		p.enterDelay/float64(p.totalCount))
}

func (p *Process) DoStatistics(time float64) {

	for _, pr := range p.processors {
		p.SetMeanLoad(int(pr.state), time)
		p.meanQueue += float64(p.queue.Size()) * (time - p.lastTime)
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

func (p *Process) GetQueue() *Queue.Queue {
	return p.queue
}

func (p *Process) SetNextActivationTime(time float64) {
	for _, pr := range p.processors {
		pr.nextActivationTime = time
		pr.state = Busy
		pr.markerInProcess = Marker.NewMarker(p.currentTime)
	}
}

func (p *Process) GetNumOfMarkersInProcess() int {
	res := 0
	for _, processor := range p.processors {
		if processor.state == Busy {
			res++
		}
	}
	return res
}
