package main

import (
	"Model/Lab2/Model"
	Model2 "Model/Model/Conditions"
	"Model/Model/Interfaces"
	"Model/Model/Processes"
	"fmt"
)

func main() {
	c := Processes.NewCreate(0.5)

	p1 := Processes.NewMultiProcessWithQueue(2, 1.0, 1)
	p2 := Processes.NewMultiProcessWithQueue(1, 2.0, 1)
	p3 := Processes.NewMultiProcessWithQueue(1, 3.0, 1)
	p4 := Processes.NewMultiProcessWithQueue(1, 1.0, 1)
	d := Processes.NewDispose()

	c.SetNextElement(p1)
	c.SetName("Create")
	c.SetDistributionType("")
	c.SetNextTime(0.0)
	c.SetCondition(Model2.NewFullRandomCondition(1))

	p1.SetName("Process 1")
	p1.SetDistributionType("")
	p1.SetNextElement(p2)
	p1.SetCondition(Model2.NewFullRandomCondition(1))

	p2.SetName("Process 2")
	p2.SetDistributionType("")
	p2.SetNextElement(p3)
	p2.SetCondition(Model2.NewFullRandomCondition(1))

	p3.SetName("Process 3")
	p3.SetDistributionType("")
	p3.SetNextElement(p4)
	p3.SetNextElement(d)
	p3.SetCondition(Model2.NewFullRandomCondition(2))

	p4.SetName("Process 4")
	p4.SetDistributionType("")
	p4.SetNextElement(p3)
	p4.SetNextElement(d)
	p4.SetCondition(Model2.NewFullRandomCondition(2))

	elements := []Interfaces.IProcess{c, p1, p2, p3, p4}
	m := Lab2.NewModel(elements)
	m.Simulate(10.0)
	fmt.Println(d.GetResult())
}
