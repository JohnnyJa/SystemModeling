package Processes

import (
	"Model/Model/Elements"
	"Model/Model/Statistic"
)

type Dispose struct {
	*Elements.BasicElement
	*Statistic.ElementStatistic
}

func NewDispose() *Dispose {
	return &Dispose{
		BasicElement:     Elements.NewBasicElement(),
		ElementStatistic: Statistic.NewElementStatistic(),
	}
}

func (d *Dispose) Finish() {
	d.ElementStatistic.AddTotalProceeded()
}

func (d *Dispose) Start() {
	d.Finish()
}

func (d *Dispose) GetLog() string {
	return "No log"

}
