package Model

import (
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
	element IElement
}

func NewPriorityCondition(element IElement, queueSizeToSkip int) *PriorityCondition {
	return &PriorityCondition{element: element}
}

func (p *PriorityCondition) MakeCondition() int {
	elements := p.element.GetNextElements()
	minQueueIndex := 0
	minQueue := math.MaxInt
	for i, el := range elements {
		if s, ok := el.(*Process); ok {
			if s.GetQueueSize() < minQueue {
				minQueue = s.GetQueueSize()
				minQueueIndex = i
			}
		} else {
			panic("element is not a process")
		}
	}
	return minQueueIndex
}
