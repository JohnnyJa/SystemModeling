package Marker

type Marker struct {
	id        int
	TimeStart float64
	Type      int
}

func NewMarker(timeStart float64) *Marker {
	return &Marker{
		TimeStart: timeStart,
	}
}

func (m *Marker) SetTimeStart(timeStart float64) {
	m.TimeStart = timeStart
}

func NewMarKerWithType(timeStart float64, Type int) *Marker {
	return &Marker{
		TimeStart: timeStart,
		Type:      Type,
	}
}
