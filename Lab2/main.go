package main

import (
	Lab2 "Model/Lab2/Model"
	"Model/Model/Interface"
	"Model/Model/Marker"
	Processes2 "Model/Model/Processes"
	"Model/Model/Transitions"
	"math/rand"
)

func main() {
	Task3()
	//c := Processes2.NewCreate(1, "Create")
	//
	//c.GetDelay = func(*Marker.Marker) float64 {
	//	return 0.5
	//}
	//
	//p1 := Processes2.NewProcess(2, "Process1", 0, 2)
	//p1.GetDelay = func(*Marker.Marker) float64 {
	//	return 1
	//}
	//
	//p2 := Processes2.NewProcess(3, "Process2", 0, 1)
	//p2.GetDelay = func(*Marker.Marker) float64 {
	//	return 2
	//}
	//
	//p3 := Processes2.NewProcess(4, "Process3", 0, 1)
	//p3.GetDelay = func(*Marker.Marker) float64 {
	//	return 3
	//}
	//d := Processes2.NewDispose(5, "Dispose")
	//
	//c.SetTransition(Transitions.NewTransition([]Interface.IElement{p1}))
	//p1.SetTransition(Transitions.NewTransition([]Interface.IElement{p2}))
	//p2.SetTransition(Transitions.NewTransition([]Interface.IElement{p3}))
	//t := Transitions.NewTransition([]Interface.IElement{p2, d})
	//
	//t.SetCondition(func(marker Marker.Marker) int {
	//	if rand.Float64() <= 0.5 {
	//		return 0
	//	}
	//	return 1
	//})
	//
	//p3.SetTransition(t)
	//
	//m := Lab2.NewModel(10, []Interface.IElement{c, p1, p2, p3, d})
	//m.Simulate()

}

func Task3() {
	c := Processes2.NewCreate(1, "Create")

	c.GetDelay = func(*Marker.Marker) float64 {
		return 0.5
	}

	p1 := Processes2.NewProcess(2, "Process1", 0, 1)
	p1.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}

	p2 := Processes2.NewProcess(3, "Process2", 0, 1)
	p2.GetDelay = func(*Marker.Marker) float64 {
		return 2
	}

	p3 := Processes2.NewProcess(4, "Process3", 0, 1)
	p3.GetDelay = func(*Marker.Marker) float64 {
		return 3
	}
	d := Processes2.NewDispose(5, "Dispose")

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{p1}))
	p1.SetTransition(Transitions.NewTransition([]Interface.IElement{p2}))
	p2.SetTransition(Transitions.NewTransition([]Interface.IElement{p3}))
	p3.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	m := Lab2.NewModel(10, []Interface.IElement{c, p1, p2, p3, d})
	m.Simulate()
}

func Task5() {
	c := Processes2.NewCreate(1, "Create")

	c.GetDelay = func(*Marker.Marker) float64 {
		return 0.5
	}

	p1 := Processes2.NewProcess(2, "Process1", 0, 2)
	p1.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}

	p2 := Processes2.NewProcess(3, "Process2", 0, 1)
	p2.GetDelay = func(*Marker.Marker) float64 {
		return 2
	}

	p3 := Processes2.NewProcess(4, "Process3", 0, 1)
	p3.GetDelay = func(*Marker.Marker) float64 {
		return 3
	}
	d := Processes2.NewDispose(5, "Dispose")

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{p1}))
	p1.SetTransition(Transitions.NewTransition([]Interface.IElement{p2}))
	p2.SetTransition(Transitions.NewTransition([]Interface.IElement{p3}))
	p3.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	m := Lab2.NewModel(10, []Interface.IElement{c, p1, p2, p3, d})
	m.Simulate()
}

func Task6() {
	c := Processes2.NewCreate(1, "Create")

	c.GetDelay = func(*Marker.Marker) float64 {
		return 0.5
	}

	p1 := Processes2.NewProcess(2, "Process1", 0, 2)
	p1.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}

	p2 := Processes2.NewProcess(3, "Process2", 0, 1)
	p2.GetDelay = func(*Marker.Marker) float64 {
		return 2
	}

	p3 := Processes2.NewProcess(4, "Process3", 0, 1)
	p3.GetDelay = func(*Marker.Marker) float64 {
		return 3
	}
	d := Processes2.NewDispose(5, "Dispose")

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{p1}))
	p1.SetTransition(Transitions.NewTransition([]Interface.IElement{p2}))
	p2.SetTransition(Transitions.NewTransition([]Interface.IElement{p3}))
	t := Transitions.NewTransition([]Interface.IElement{p2, d})

	t.SetCondition(func(marker Marker.Marker) int {
		if rand.Float64() <= 0.5 {
			return 0
		}
		return 1
	})

	p3.SetTransition(t)

	m := Lab2.NewModel(10, []Interface.IElement{c, p1, p2, p3, d})
	m.Simulate()
}
