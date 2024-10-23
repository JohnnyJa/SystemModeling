package main

import (
	Lab3_2 "Model/Lab3_2/Model"
	"Model/Model/Interface"
	"Model/Model/Marker"
	"Model/Model/Processes"
	"Model/Model/Transitions"
	"Model/funRand"
	"math/rand"
)

func main() {
	c := Processes.NewCreate(1, "Create")
	c.GetDelay = func(*Marker.Marker) float64 {
		return funRand.Exp(15)
	}

	wr := Processes.NewProcess(2, "Waiting room", 10000, 2)
	wr.GetDelay = func(m *Marker.Marker) float64 {
		switch m.Type {
		case 1:
			return 15
		case 2:
			return 40
		case 3:
			return 30
		}
		return 0
	}

	ch := Processes.NewProcess(3, "To chamber", 10000, 3)
	ch.GetDelay = func(m *Marker.Marker) float64 {
		return funRand.Unif(3, 8)
	}
	rtl := Processes.NewProcess(4, "To lab", 10000, -1)
	rtl.GetDelay = func(m *Marker.Marker) float64 {
		return funRand.Unif(2, 5)
	}

	lr := Processes.NewProcess(5, "Lab register", 10000, 1)
	lr.GetDelay = func(m *Marker.Marker) float64 {
		return funRand.Erlang(4.5, 3)
	}

	la := Processes.NewProcess(6, "Lab work", 10000, 2)
	la.GetDelay = func(m *Marker.Marker) float64 {
		return funRand.Erlang(4, 2)
	}

	rtw := Processes.NewProcess(7, "Back to waiting room", 10000, -1)
	rtw.GetDelay = func(m *Marker.Marker) float64 {
		return funRand.Unif(2, 5)
	}
	rtw.ChangeMarker = func(m *Marker.Marker) {
		m.Type = 1
	}

	d := Processes.NewDispose(8, "Dispose")

	c.MarkerGenerator = func(time float64) *Marker.Marker {
		r := rand.Float64()

		if r < 0.5 {
			return Marker.NewMarKerWithType(time, 1)
		} else if r >= 0.5 && r < 0.6 {
			return Marker.NewMarKerWithType(time, 2)
		} else {
			return Marker.NewMarKerWithType(time, 3)
		}
	}

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{wr}))

	t1 := Transitions.NewTransition([]Interface.IElement{ch, rtl})

	t1.SetCondition(func(marker Marker.Marker) int {
		if marker.Type == 1 {
			return 0
		}
		return 1
	})

	wr.SetTransition(t1)

	ch.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	rtl.SetTransition(Transitions.NewTransition([]Interface.IElement{lr}))
	lr.SetTransition(Transitions.NewTransition([]Interface.IElement{la}))
	rtw.SetTransition(Transitions.NewTransition([]Interface.IElement{wr}))
	t2 := Transitions.NewTransition([]Interface.IElement{rtw, d})
	t2.SetCondition(func(marker Marker.Marker) int {
		if marker.Type == 2 {
			return 0
		}
		return 1
	})

	la.SetTransition(t2)

	m := Lab3_2.NewModel(10000, []Interface.IElement{c, wr, ch, rtl, lr, la, rtw, d})
	m.Simulate()
}
