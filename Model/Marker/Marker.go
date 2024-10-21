package Marker

type Marker struct {
	id        int
	TimeStart float64
}

func NewMarker(timeStart float64) *Marker {
	return &Marker{
		TimeStart: timeStart,
	}
}

func (m *Marker) SetTimeStart(timeStart float64) {
	m.TimeStart = timeStart
}
