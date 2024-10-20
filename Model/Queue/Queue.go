package ModelQueue

import "errors"

type Queue struct {
	currentQueueSize int
	maxQueueSize     int
}

func NewQueue(maxQueueSize int) *Queue {
	return &Queue{maxQueueSize: maxQueueSize}
}

func (q *Queue) RemoveFromQueue() {
	if q.currentQueueSize > 0 {
		q.currentQueueSize--
	}
}

func (q *Queue) AddToQueue() error {
	if q.currentQueueSize >= q.maxQueueSize {
		return errors.New("queue is full")
	}
	q.currentQueueSize++

	return nil
}

func (q *Queue) GetCurrentQueueSize() int {
	return q.currentQueueSize
}
