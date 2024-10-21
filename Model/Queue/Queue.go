package Queue

import (
	"Model/Model/Statistic"
	"errors"
)

type ModelQueue struct {
	currentQueueSize int
	maxQueueSize     int
	markers          []Statistic.Marker
}

func NewQueue(maxQueueSize int) *ModelQueue {
	return &ModelQueue{maxQueueSize: maxQueueSize}
}

func (q *ModelQueue) RemoveFromQueue() Statistic.Marker {
	if q.currentQueueSize > 0 {
		q.currentQueueSize--
	}

	res := q.markers[0]
	q.markers = q.markers[1:]
	return res
}

func (q *ModelQueue) AddToQueue(marker Statistic.Marker) error {
	if q.currentQueueSize >= q.maxQueueSize {
		return errors.New("queue is full")
	}
	q.currentQueueSize++
	q.markers = append(q.markers, marker)
	return nil
}

func (q *ModelQueue) GetCurrentQueueSize() int {
	return q.currentQueueSize
}

func (q *ModelQueue) GetFirst() Statistic.Marker {
	return q.markers[0]

}
