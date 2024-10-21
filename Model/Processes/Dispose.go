package Processes

import (
	"Model/Model/Marker"
	"fmt"
	"math"
)

type DisposeStatistic struct {
	disposed int

	avgLeaveTime  float64
	avgClientTime float64

	lastTime float64
}

func (d *DisposeStatistic) GetLog() string {
	return fmt.Sprintf("Disposed: %d", d.disposed)
}

func (d *DisposeStatistic) GetResults() string {
	log := ""
	log += fmt.Sprintf("Average time between leaving: %f\n", d.avgLeaveTime/float64(d.disposed))
	log += fmt.Sprintf("Average client time: %f\n", d.avgClientTime/float64(d.disposed))

	return log
}

type Dispose struct {
	BaseElement
	DisposeStatistic
}

func NewDispose(id int, name string) *Dispose {
	return &Dispose{
		BaseElement: BaseElement{
			Id:                 id,
			Name:               name,
			nextActivationTime: math.MaxFloat64,
		},
	}
}

func (d *Dispose) RunToCurrentTime(currentTime float64) {
	d.currentTime = currentTime
}

func (d *Dispose) TakeMarker(m *Marker.Marker) {
	d.avgLeaveTime += d.currentTime - d.lastTime
	d.avgClientTime += d.currentTime - m.TimeStart
	d.disposed++

	d.lastTime = d.currentTime
}

func (d *Dispose) GetLog() string {
	return fmt.Sprintf("%s,\n Stats: %s", d.BaseElement.GetLog(), d.DisposeStatistic.GetLog())
}

func (d *Dispose) GetResults() string {
	return d.DisposeStatistic.GetResults()
}
