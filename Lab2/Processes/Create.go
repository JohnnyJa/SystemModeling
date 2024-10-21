package Processes

import (
	"Model/Lab2/Marker"
	"Model/Lab2/Transitions"
	"fmt"
)

type CreateStatistic struct {
	created int
}

func (c *CreateStatistic) GetLog() string {
	return fmt.Sprintf("Created: %d", c.created)

}

type Create struct {
	BaseElement
	CreateStatistic

	delay      float64
	transition *Transitions.Transition

	markerGenerator func() *Marker.Marker
}

func NewCreate(id int, name string, delay float64) *Create {
	return &Create{
		BaseElement: BaseElement{
			id,
			name,
			0,
			0,
		},
		delay: delay,
		markerGenerator: func() *Marker.Marker {
			return Marker.NewMarker()
		},
	}
}

func (c *Create) SetTransition(transition *Transitions.Transition) {
	c.transition = transition
}

func (c *Create) RunToCurrentTime(currentTime float64) {
	if c.nextActivationTime <= currentTime {
		c.CreateNewMarker()
		c.nextActivationTime = currentTime + c.delay
	}
	c.currentTime = currentTime
}

func (c *Create) CreateNewMarker() {
	c.created++
	m := c.markerGenerator()

	m.SetTimeStart(c.currentTime)
	c.transition.PushMarkerToNextNode(m)
}

func (c *Create) TakeMarker(*Marker.Marker) {
	// do nothing
}

func (c *Create) GetLog() string {
	return fmt.Sprintf("%s,\n Stats: %s", c.BaseElement.GetLog(), c.CreateStatistic.GetLog())
}

func (c *Create) GetResults() string {
	return ""
}
