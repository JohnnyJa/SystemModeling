package Elements

import (
	"Model/Model/Statistic"
	"strconv"
)

var NumOFElements int = 0

type IElement interface {
	Start(Statistic.Marker)
	Finish()
	GetId() int

	GetName() string
}

type BasicElement struct {
	id   int
	name string
}

func NewBasicElement() *BasicElement {
	el := &BasicElement{
		id:   NumOFElements,
		name: "Element" + strconv.Itoa(NumOFElements),
	}
	NumOFElements++

	return el
}

func NewAnonymusElement() *BasicElement {
	el := &BasicElement{
		id:   -1,
		name: "anonymus",
	}

	return el
}

func (e *BasicElement) GetId() int {
	return e.id
}

func (e *BasicElement) SetName(name string) {
	e.name = name
}

func (e *BasicElement) GetName() string {
	return e.name
}

type ElementState int

const (
	Free ElementState = iota
	Busy
)
