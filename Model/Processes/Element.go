package Processes

import (
	"fmt"
)

type BaseElement struct {
	Id   int
	Name string

	nextActivationTime float64
	currentTime        float64
}

func (e *BaseElement) GetActivationTime() float64 {
	return e.nextActivationTime
}

func (e *BaseElement) GetLog() string {
	return fmt.Sprintf("Element: %s, Id: %d", e.Name, e.Id)
}
