package Statistic

import "fmt"

type QueueStatistic struct {
	meanQueueSize float64
	clientAmount  float64
	lastTime      float64
}

func NewQueueStatistic() *QueueStatistic {
	return &QueueStatistic{
		meanQueueSize: 0}
}

func (qs *QueueStatistic) SetLastTime(lastTime float64) {
	qs.lastTime = lastTime
}
func (qs *QueueStatistic) CountMeanQueue(queueSize int, currentTime float64) {
	qs.meanQueueSize += float64(queueSize) * (currentTime - qs.lastTime)
}

func (qs *QueueStatistic) CountAverageClientAmount(clientAmount int, currentTime float64) {
	qs.clientAmount += float64(clientAmount) * (currentTime - qs.lastTime)
}

func (qs *QueueStatistic) GetResult() string {
	return fmt.Sprintf("Mean queue size: %f", qs.meanQueueSize/qs.lastTime)
}
