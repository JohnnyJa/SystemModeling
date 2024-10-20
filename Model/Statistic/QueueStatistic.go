package Statistic

import "fmt"

type QueueStatistic struct {
	meanQueueSize float64
	lastTime      float64
}

func NewQueueStatistic() *QueueStatistic {
	return &QueueStatistic{
		meanQueueSize: 0}
}

func (qs *QueueStatistic) CountMeanQueue(queueSize int, currentTime float64) {
	qs.meanQueueSize += float64(queueSize) * (currentTime - qs.lastTime)
	qs.lastTime = currentTime
}

func (qs *QueueStatistic) GetResult() string {
	return fmt.Sprintf("Mean queue size: %f", qs.meanQueueSize/qs.lastTime)
}
