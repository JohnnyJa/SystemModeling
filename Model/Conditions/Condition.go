package Conditions

import (
	ModelQueue "Model/Model/Queue"
	"errors"
	"math"
	"math/rand"
)

type ICondition interface {
	MakeCondition() int
}

type FullRandomCondition struct {
	numOfConditions int
}

func NewFullRandomCondition(numOfConditions int) *FullRandomCondition {
	return &FullRandomCondition{numOfConditions: numOfConditions}
}

func (r *FullRandomCondition) MakeCondition() int {
	return rand.Intn(r.numOfConditions)
}

type RandomCondition struct {
	conditionsArr []float64
}

func NewRandomCondition(conditionsArr []float64) (*RandomCondition, error) {
	sum := 0.0

	for _, el := range conditionsArr {
		sum += el
	}

	if sum != 1.0 {
		return nil, errors.New("sum of conditions probabilities must be equal to 1")
	}

	return &RandomCondition{conditionsArr: conditionsArr}, nil
}

func (r *RandomCondition) MakeCondition() int {
	rnd := rand.Float64()
	sum := 0.0

	for i, el := range r.conditionsArr {
		sum += el
		if rnd <= sum {
			return i
		}
	}

	return -1
}

type PriorityCondition struct {
	queues []*ModelQueue.Queue
}

func NewPriorityCondition() *PriorityCondition {
	return &PriorityCondition{queues: make([]*ModelQueue.Queue, 0)}
}

func (p *PriorityCondition) AddQueue(queue *ModelQueue.Queue) {
	p.queues = append(p.queues, queue)
}

func (p *PriorityCondition) SetQueues(queues []*ModelQueue.Queue) {
	p.queues = queues
}

func (p *PriorityCondition) MakeCondition() int {

	minQueueIndex := 0
	minQueue := math.MaxInt
	for i, q := range p.queues {
		if q.GetCurrentQueueSize() < minQueue {
			minQueue = q.GetCurrentQueueSize()
			minQueueIndex = i
		}
	}

	return minQueueIndex
}
