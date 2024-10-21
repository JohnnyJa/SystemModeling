package Conditions

import (
	"Model/Model/Elements"
	"Model/Model/Statistic"
)

type ITransition interface {
	StartNextElement(marker Statistic.Marker)
	SetNextElement(Elements.IElement)
	SetNextElements([]Elements.IElement)
	SetCondition(ICondition)
	GetTransition() *Transition
}

type Transition struct {
	condition    ICondition
	nextElements []Elements.IElement
}

func NewTransition() *Transition {
	return &Transition{
		nextElements: make([]Elements.IElement, 0),
	}
}

func (t *Transition) GetTransition() *Transition {
	return t
}

func (t *Transition) SetNextElement(el Elements.IElement) {
	t.nextElements = append(t.nextElements, el)
}

func (t *Transition) SetNextElements(elements []Elements.IElement) {
	t.nextElements = elements
}

func (t *Transition) StartNextElement(marker Statistic.Marker) {
	t.nextElements[t.condition.MakeCondition()].Start(marker)
}

func (t *Transition) SetCondition(condition ICondition) {
	t.condition = condition
}
