package Interface

import "Model/Lab2/Marker"

type ILogger interface {
	GetLog() string

	GetResults() string
}
type IElement interface {
	RunToCurrentTime(float64)
	GetActivationTime() float64
	TakeMarker(marker *Marker.Marker)
}
