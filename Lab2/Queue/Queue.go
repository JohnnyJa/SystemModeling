package Queue

import "Model/Lab2/Marker"

type Queue struct {
	elements []*Marker.Marker
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(element *Marker.Marker) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Pop() *Marker.Marker {
	if len(q.elements) == 0 {
		return nil
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Head() *Marker.Marker {
	if len(q.elements) == 0 {
		return nil
	}
	return q.elements[0]
}
