package Processes

import (
	"Model/Model/Marker"
	"Model/Model/Transitions"
	"fmt"
)

type CreateStatistic struct {
	Created int
}

func (c *CreateStatistic) GetLog() string {
	return fmt.Sprintf("Created: %d", c.Created)

}

type Create struct {
	BaseElement
	CreateStatistic

	transition *Transitions.Transition

	markerGenerator func(float64) *Marker.Marker
}

func NewCreate(id int, name string) *Create {
	return &Create{
		BaseElement: BaseElement{
			Id:                 id,
			Name:               name,
			currentTime:        0,
			nextActivationTime: 0,
		},
		markerGenerator: func(time float64) *Marker.Marker {
			return Marker.NewMarker(time)
		},
	}
}

func (c *Create) SetTransition(transition *Transitions.Transition) {
	c.transition = transition
}

func (c *Create) RunToCurrentTime(currentTime float64) {
	if c.nextActivationTime <= currentTime {
		c.CreateNewMarker()
		c.nextActivationTime = currentTime + c.GetDelay()
	}
	c.currentTime = currentTime
}

func (c *Create) CreateNewMarker() {
	c.Created++
	m := c.markerGenerator(c.currentTime)

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
