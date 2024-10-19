package Model

import (
	"fmt"
	"math"
)

type Dispose struct {
	element  *Element
	disposed int
}

func NewDispose() *Dispose {
	d := &Dispose{element: NewElementWithDelay(0)}
	d.element.tnext = math.MaxFloat64
	d.disposed = 0
	return d
}

func (d *Dispose) OutAct() {
}

func (d *Dispose) InAct() {
	d.element.OutAct()
	d.disposed++

}

func (d *Dispose) DoStatistic(time float64) {

}

func (d *Dispose) GetBaseElement() *Element {
	return d.element
}

func (d *Dispose) PrintInfo() {
}

func (d *Dispose) PrintResult() {
	d.element.PrintResult()
	fmt.Printf(" disposed = %d\n", d.disposed)
}

func (d *Dispose) GetId() int {
	return d.element.id
}

func (d *Dispose) SetName(name string) {
	d.element.name = name
}

func (d *Dispose) GetNextElements() []IElement {
	return nil
}

func (d *Dispose) GetTNext() float64 {
	return d.element.GetTNext()
}

func (d *Dispose) SetTNext(tNext float64) {
	d.element.SetTNext(tNext)
}

func (d *Dispose) GetTCurr() float64 {
	return d.element.GetTCurr()
}

func (d *Dispose) SetTCurr(tCurr float64) {
	d.element.SetTCurr(tCurr)
}
