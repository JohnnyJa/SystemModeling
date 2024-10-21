package main

import (
	"Model/Lab2/Interface"
	Lab2 "Model/Lab2/Model"
	"Model/Lab2/Processes"
	"Model/Lab2/Transitions"
	"math/rand"
)

func main() {
	c := Processes.NewCreate(1, "Create", 0.5)
	p1 := Processes.NewProcess(2, "Process", 0, 1, 2)
	p2 := Processes.NewProcess(2, "Process", 0, 2, 1)
	p3 := Processes.NewProcess(2, "Process", 0, 3, 1)

	d := Processes.NewDispose(3, "Dispose")

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
