package Queue

import (
	"Model/Model/Marker"
	"sort"
)

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

func (q *Queue) PopBack() *Marker.Marker {
	if len(q.elements) == 0 {
		return nil
	}

	element := q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
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

func (q *Queue) OrderQueue(criteria func(*Marker.Marker, *Marker.Marker) bool) {
	if len(q.elements) == 0 {
		return
	}

	sort.Slice(q.elements, func(i, j int) bool {
		return criteria(q.elements[i], q.elements[j])
	})
}
