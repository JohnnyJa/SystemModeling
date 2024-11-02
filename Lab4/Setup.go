package Lab4

import (
	Lab4 "Model/Lab4/Model"
	"Model/Model/Interface"
	"Model/Model/Marker"
	"Model/Model/Processes"
	"Model/Model/Transitions"
	"math/rand"
)

func Setup(n int, time float64) *Lab4.Model {
	c := Processes.NewCreate(1, "Create")
	c.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}

	processes := make([]*Processes.Process, n)
	for i := 0; i < n; i++ {
		processes[i] = Processes.NewProcess(i+2, "Process", 10000, 1)
		processes[i].GetDelay = func(m *Marker.Marker) float64 {
			return 1
		}

		if i > 0 {
			processes[i-1].SetTransition(Transitions.NewTransition([]Interface.IElement{processes[i]}))
		}
	}

	c.SetTransition(Transitions.NewTransition([]Interface.IElement{processes[0]}))
	d := Processes.NewDispose(n+2, "Dispose")
	processes[n-1].SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	c.MarkerGenerator = func(time float64) *Marker.Marker {
		return Marker.NewMarker(time)
	}

	els := make([]Interface.IElement, n+2)
	els[0] = c
	for i := 0; i < n; i++ {
		els[i+1] = processes[i]
	}
	els[n+1] = d
	return Lab4.NewModel(time, els)
}

func Setup2(time float64) *Lab4.Model {
	processes := make([]Interface.IElement, 0)

	c := Processes.NewCreate(1, "Create")
	c.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, c)

	p1 := Processes.NewProcess(2, "Process1", 10000, 1)
	p1.GetDelay = func(m *Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, p1)

	p2 := Processes.NewProcess(3, "Process2", 10000, 1)
	p2.GetDelay = func(m *Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, p2)

	t := Transitions.NewTransition([]Interface.IElement{p1, p2})
	t.SetCondition(func(m Marker.Marker) int {
		if rand.Float64() < 0.5 {
			return 0
		}
		return 1
	})

	c.SetTransition(t)

	prev1 := p1
	prev2 := p2
	for i := 0; i < 10; i++ {
		pr1 := Processes.NewProcess(i+2, "Process", 10000, 1)
		pr1.GetDelay = func(m *Marker.Marker) float64 {
			return 1
		}
		prev1.SetTransition(Transitions.NewTransition([]Interface.IElement{pr1}))

		prev1 = pr1

		processes = append(processes, pr1)

		pr2 := Processes.NewProcess(i+3, "Process", 10000, 1)
		pr2.GetDelay = func(m *Marker.Marker) float64 {
			return 1
		}
		prev2.SetTransition(Transitions.NewTransition([]Interface.IElement{pr2}))

		prev2 = pr2

		processes = append(processes, pr2)
	}

	d := Processes.NewDispose(0, "Dispose")
	prev1.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))
	prev2.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	c.MarkerGenerator = func(time float64) *Marker.Marker {
		return Marker.NewMarker(time)
	}

	return Lab4.NewModel(time, processes)
}

func Setup3(time float64) *Lab4.Model {
	processes := make([]Interface.IElement, 0)

	c := Processes.NewCreate(1, "Create")
	c.GetDelay = func(*Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, c)

	p1 := Processes.NewProcess(2, "Process1", 10000, 1)
	p1.GetDelay = func(m *Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, p1)

	p2 := Processes.NewProcess(3, "Process2", 10000, 1)
	p2.GetDelay = func(m *Marker.Marker) float64 {
		return 1
	}
	processes = append(processes, p2)

	t := Transitions.NewTransition([]Interface.IElement{p1, p2})
	t.SetCondition(func(m Marker.Marker) int {
		if rand.Float64() < 0.5 {
			return 0
		}
		return 1
	})

	c.SetTransition(t)

	prev1 := p1
	prev2 := p2
	for i := 0; i < 10; i++ {

		pr1 := Processes.NewProcess(i+2, "Process", 10000, 1)
		pr1.GetDelay = func(m *Marker.Marker) float64 {
			return 1
		}

		prev1.SetTransition(Transitions.NewTransition([]Interface.IElement{pr1}))

		prev1 = pr1

		processes = append(processes, pr1)

		pr2 := Processes.NewProcess(i+3, "Process", 10000, 1)
		pr2.GetDelay = func(m *Marker.Marker) float64 {
			return 1
		}
		prev2.SetTransition(Transitions.NewTransition([]Interface.IElement{pr2}))

		prev2 = pr2

		processes = append(processes, pr2)
	}

	d := Processes.NewDispose(0, "Dispose")

	t1 := Transitions.NewTransition([]Interface.IElement{prev1, prev2})
	t1.SetCondition(func(m Marker.Marker) int {
		if rand.Float64() < 0.5 {
			return 0
		}
		return 1
	})
	prev1.SetTransition(t1)

	prev2.SetTransition(Transitions.NewTransition([]Interface.IElement{d}))

	c.MarkerGenerator = func(time float64) *Marker.Marker {
		return Marker.NewMarker(time)
	}

	return Lab4.NewModel(time, processes)
}
