package main

import (
	Lab2 "Model/Lab2/Model"
	"Model/Model/Interface"
	Processes2 "Model/Model/Processes"
	"Model/Model/Transitions"
	"math/rand"
)

func main() {
	c := Processes2.NewCreate(1, "Create", 0.5)
	p1 := Processes2.NewProcess(2, "Process", 0, 1, 2)
	p2 := Processes2.NewProcess(2, "Process", 0, 2, 1)
	p3 := Processes2.NewProcess(2, "Process", 0, 3, 1)

	d := Processes2.NewDispose(3, "Dispose")

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{p1}))
	p1.SetTransition(Transitions.NewTransition([]Interface.IElement{p2}))
	p2.SetTransition(Transitions.NewTransition([]Interface.IElement{p3}))
	t := Transitions.NewTransition([]Interface.IElement{p2, d})

	t.SetCondition(func() int {
		if rand.Float64() <= 0.5 {
			return 0
		}
		return 1
	})

	p3.SetTransition(t)

	m := Lab2.NewModel(10, []Interface.IElement{c, p1, p2, p3, d})
	m.Simulate()

}
