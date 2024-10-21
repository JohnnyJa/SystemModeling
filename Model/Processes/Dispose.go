package Processes

import (
	"Model/Model/Elements"
	"Model/Model/Statistic"
)

type Dispose struct {
	*Elements.ProcessElement
	*Statistic.ElementStatistic
}

func NewDispose() *Dispose {
	return &Dispose{
		ProcessElement:   Elements.NewProcessElementWithDelay(0),
		ElementStatistic: Statistic.NewElementStatistic(),
	}
}

func (d *Dispose) Finish() {
	d.ProcessElement.Finish()
	d.ElementStatistic.AddTotalProceeded()
}

func (d *Dispose) Start() {
	d.ElementStatistic.CountTimeBetweenLeft(d.GetCurrentTime())
	d.ElementStatistic.SetLastTime(d.GetCurrentTime())

	d.Finish()
}

func (d *Dispose) GetLog() string {
	return "No log"

}
