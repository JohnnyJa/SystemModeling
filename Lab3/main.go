package main

import (
	Lab3 "Model/Lab3/Model"
	"Model/Model/Conditions"
	"Model/Model/Elements"
	"Model/Model/Interfaces"
	"Model/Model/Processes"
	"Model/Model/Statistic"
	"Model/funRand"
)

func main() {
	c := Processes.NewCreate(0.5)
	p1 := Processes.NewSingleProcessWithQueue(0.3, 3)
	p2 := Processes.NewSingleProcessWithQueue(0.3, 3)
	d := Processes.NewDispose()

	c.SetNextElement(p1)
	c.SetNextElement(p2)
	c.SetName("Create")
	c.SetDistributionType("exp")
	c.SetNextTime(0.1)

	cond := Conditions.NewPriorityCondition()
	cond.AddQueue(p1.GetQueue())
	cond.AddQueue(p2.GetQueue())

	c.SetCondition(cond)

	p1.SetNextElement(d)
	p1.SetName("Process1")
	p1.SetDistributionType("exp")
	p1.SetNextTime(funRand.Norm(1, 0.3))
	m1 := *Statistic.NewMarker(0.0)
	m2 := *Statistic.NewMarker(0.0)
	m3 := *Statistic.NewMarker(0.0)
	m4 := *Statistic.NewMarker(0.0)

	_ = p1.AddToQueue(m1)
	_ = p1.AddToQueue(m2)
	p1.SetState(Elements.Busy)
	p1.SetCondition(Conditions.NewFullRandomCondition(1))

	p2.SetNextElement(d)
	p2.SetName("Process2")
	p2.SetDistributionType("exp")
	p2.SetNextTime(funRand.Norm(1, 0.3))
	_ = p2.AddToQueue(m3)
	_ = p2.AddToQueue(m4)
	p2.SetState(Elements.Busy)
	p2.SetCondition(Conditions.NewFullRandomCondition(1))

	elems := []Interfaces.IProcess{c, p1, p2, d}
	m := Lab3.NewModel(elems)
	m.Simulate(10.0)
	//fmt.Println("-------------------\nDispose: " + d.GetResult())
}
