package Processes

import (
	"Model/Model/Marker"
	"fmt"
	"math"
)

type DisposeStatistic struct {
	disposed int
}

func (d *DisposeStatistic) GetLog() string {
	return fmt.Sprintf("Disposed: %d", d.disposed)
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

func (d *Dispose) TakeMarker(*Marker.Marker) {
	d.disposed++
}

func (d *Dispose) GetLog() string {
	return fmt.Sprintf("%s,\n Stats: %s", d.BaseElement.GetLog(), d.DisposeStatistic.GetLog())
}

func (d *Dispose) GetResults() string {
	return ""
}
