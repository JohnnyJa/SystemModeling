package main

import (
	Lab3 "Model/Lab3_1/Model"
	"Model/Model/Interface"
	"Model/Model/Marker"
	Processes2 "Model/Model/Processes"
	"Model/Model/Transitions"
	"Model/funRand"
)

func main() {
	c := Processes2.NewCreate(1, "Create")
	p1 := Processes2.NewProcess(2, "Cashier1", 3, 1)
	p2 := Processes2.NewProcess(3, "Cashier2", 3, 1)
	d := Processes2.NewDispose(4, "Dispose")

	q1 := p1.GetQueue()
	q2 := p2.GetQueue()

	c.GetDelay = func(*Marker.Marker) float64 {
		return funRand.Exp(0.5)
	}

	t := Transitions.NewTransition([]Interface.IElement{p1, p2})
	t.SetCondition(func(Marker.Marker) int {
		if q1.Size() > q2.Size() {
			return 1
		}
		return 0
	})

	c.SetTransition(t)
	c.SetNextActivationTime(0.1)

	p1.GetDelay = func(*Marker.Marker) float64 {
		return funRand.Exp(0.3)
	}
	p1.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	m1 := Marker.NewMarker(0.0)
	m2 := Marker.NewMarker(0.0)
	q1.Push(m1)
	q1.Push(m2)

	p1.SetNextActivationTime(funRand.Norm(1, 0.3))

	p2.GetDelay = func(*Marker.Marker) float64 {
		return funRand.Exp(0.3)
	}

	p2.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	m3 := Marker.NewMarker(0.0)
	m4 := Marker.NewMarker(0.0)
	q2.Push(m3)
	q2.Push(m4)

	p2.SetNextActivationTime(funRand.Norm(1, 0.3))

	m := Lab3.NewModel(10, []Interface.IElement{c, p1, p2, d})
	m.Simulate()

}
