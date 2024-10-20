package Elements

import "math"

type ITimeElement interface {
	MoveToCurrentTime()
	GetNextTime() float64
	SetCurrentTime(float64)
}

type TimeElement struct {
	currentTime float64
	nextTime    float64
}

func NewTimeElement() *TimeElement {
	return &TimeElement{
		currentTime: 0,
		nextTime:    math.MaxFloat64,
	}
}

func (te *TimeElement) GetCurrentTime() float64 {
	return te.currentTime
}

func (te *TimeElement) GetNextTime() float64 {
	return te.nextTime
}

func (te *TimeElement) SetNextTime(nextTime float64) {
	te.nextTime = nextTime
}

func (te *TimeElement) SetCurrentTime(currentTime float64) {
	te.currentTime = currentTime
}
