package Transitions

import (
	"Model/Lab2/Interface"
	"Model/Lab2/Marker"
)

type Transition struct {
	nextElements []Interface.IElement
	condition    func() int
}

func NewTransition(elements []Interface.IElement) *Transition {
	return &Transition{
		nextElements: elements,
		condition: func() int {
			return 0
		},
	}
}

func (t *Transition) SetCondition(condition func() int) {
	t.condition = condition
}

func (t *Transition) PushMarkerToNextNode(marker *Marker.Marker) {
	nextElementIndex := t.condition()
	t.nextElements[nextElementIndex].TakeMarker(marker)
}
