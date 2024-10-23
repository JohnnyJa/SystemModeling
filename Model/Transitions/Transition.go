package Transitions

import (
	"Model/Model/Interface"
	"Model/Model/Marker"
)

type Transition struct {
	nextElements []Interface.IElement
	condition    func(marker Marker.Marker) int
}

func NewTransition(elements []Interface.IElement) *Transition {
	return &Transition{
		nextElements: elements,
		condition: func(Marker.Marker) int {
			return 0
		},
	}
}

func (t *Transition) SetCondition(condition func(marker Marker.Marker) int) {
	t.condition = condition
}

func (t *Transition) PushMarkerToNextNode(marker *Marker.Marker) {
	nextElementIndex := t.condition(*marker)
	t.nextElements[nextElementIndex].TakeMarker(marker)
}
