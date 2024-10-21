package Processes

import (
	"Model/Model/Elements"
	"Model/Model/Statistic"
	"fmt"
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

func (d *Dispose) Start(marker Statistic.Marker) {
	d.ElementStatistic.CountTimeBetweenLeft(d.GetCurrentTime())
	d.ElementStatistic.SetLastTime(d.GetCurrentTime())
	d.ElementStatistic.CountMarkerTime(marker, d.GetCurrentTime())
	d.Finish()
}

func (d *Dispose) GetLog() string {
	return "No log"

}

func (d *Dispose) GetResult() string {
	return d.ElementStatistic.GetResult() + fmt.Sprintf("Average Client time: %f\n", d.ElementStatistic.GetClientTime()/float64(d.ElementStatistic.GetTotalProceeded()))
}
