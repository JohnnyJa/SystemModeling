package Statistic

import "fmt"

type IStatistic interface {
	GetResult() string
	GetLog() string
}

type ElementStatistic struct {
	totalProceeded int
	failures       int
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

func (s *ElementStatistic) GetFailure() int {
	return s.failures
}

func (s *ElementStatistic) GetResult() string {
	return fmt.Sprintf("Total proceeded: %d Failures: %d", s.totalProceeded, s.failures)
}
