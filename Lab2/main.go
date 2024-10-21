package main

import (
	Lab2 "Model/Lab2/Model"
	"Model/Model/Interface"
	Processes2 "Model/Model/Processes"
	"Model/Model/Transitions"
	"math/rand"
)

func main() {
	c := Processes2.NewCreate(1, "Create")
	c.SetDelay(0.5, 0)
	c.SetDistribution(Processes2.None)

	p1 := Processes2.NewProcess(2, "Process1", 0, 2)
	p1.SetDelay(1, 0)
	p1.SetDistribution(Processes2.None)

	p2 := Processes2.NewProcess(3, "Process2", 0, 1)
	p2.SetDelay(2, 0)
	p2.SetDistribution(Processes2.None)

	p3 := Processes2.NewProcess(4, "Process3", 0, 1)
	p3.SetDelay(3, 0)
	p3.SetDistribution(Processes2.None)

	d := Processes2.NewDispose(5, "Dispose")

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
