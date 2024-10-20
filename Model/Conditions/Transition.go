package Conditions

import "Model/Model/Elements"

type ITransition interface {
	StartNextElement()
	SetNextElement(Elements.IElement)
	SetNextElements([]Elements.IElement)
	SetCondition(ICondition)
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

func (t *Transition) SetNextElement(el Elements.IElement) {
	t.nextElements = append(t.nextElements, el)
}

func (t *Transition) SetNextElements(elements []Elements.IElement) {
	t.nextElements = elements
}

func (t *Transition) StartNextElement() {
	t.nextElements[t.condition.MakeCondition()].Start()
}

func (t *Transition) SetCondition(condition ICondition) {
	t.condition = condition
}
