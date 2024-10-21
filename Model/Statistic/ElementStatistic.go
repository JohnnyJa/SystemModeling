package Statistic

import "fmt"

type IStatistic interface {
	GetResult() string
	GetLog() string
}

type ElementStatistic struct {
	totalProceeded  int
	failures        int
	timeBetweenLeft float64
	lastTime        float64
	clientTime      float64
}

func NewElementStatistic() *ElementStatistic {
	return &ElementStatistic{totalProceeded: 0}
}

func (s *ElementStatistic) AddTotalProceeded() {
	s.totalProceeded++
}

func (s *ElementStatistic) AddFailure() {
	s.failures++
}

func (s *ElementStatistic) CountTimeBetweenLeft(currentTime float64) {
	s.timeBetweenLeft += currentTime - s.lastTime
	s.lastTime = currentTime

}
func (s *ElementStatistic) SetLastTime(lastTime float64) {
	s.lastTime = lastTime
}

func (s *ElementStatistic) GetFailure() int {
	return s.failures
}

func (s *ElementStatistic) GetTotalProceeded() int {
	return s.totalProceeded
}

func (s *ElementStatistic) GetClientTime() float64 {
	return s.clientTime

}

func (s *ElementStatistic) GetResult() string {
	return fmt.Sprintf("Total proceeded: %d Failures: %d \nAverage time between left: %f\n", s.totalProceeded, s.failures, s.timeBetweenLeft/float64(s.totalProceeded))
}

func (s *ElementStatistic) CountMarkerTime(marker Marker, time float64) {
	s.clientTime += time - marker.GetTimeStart()
}
