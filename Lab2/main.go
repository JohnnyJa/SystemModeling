package main

import (
	"Model/Lab2/Model"
	Model2 "Model/Model"
)

func main() {
	c := Model2.NewCreate(0.5)
	p1 := Model2.NewProcess(1.0, 2)
	p2 := Model2.NewProcess(2.0, 1)
	p3 := Model2.NewProcess(3.0, 1)
	p4 := Model2.NewProcess(1.0, 1)

	d := Model2.NewDispose()
	//fmt.Printf("id0 = %d\n, id1 = %d", c.GetId(), p1.GetId())

	c.SetNextElement(p1)
	c.SetName("Create")
	c.SetDistribution("exp")

	p1.SetMaxQueue(100)
	p1.SetName("Process 1")
	p1.SetDistribution("exp")
	p1.SetNextElement(p2)

	p2.SetMaxQueue(100)
	p2.SetName("Process 2")
	p2.SetDistribution("exp")
	p2.SetNextElement(p3)

	p3.SetMaxQueue(100)
	p3.SetName("Process 3")
	p3.SetDistribution("exp")
	p3.SetNextElement(p4)
	p3.SetNextElement(d)
	p3.SetCondition(Model2.NewFullRandomCondition(2))

	p4.SetMaxQueue(1)
	p4.SetName("Process 4")
	p4.SetDistribution("exp")
	p4.SetNextElement(p3)
	p4.SetNextElement(d)
	p4.SetCondition(Model2.NewFullRandomCondition(2))

	elements := []Model2.IElement{c, p1, p2, p3, p4, d}
	m := Lab2.NewModel(elements)
	m.Simulate(10.0)
}
