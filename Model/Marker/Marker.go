package Marker

type Marker struct {
	Id        int
	timeStart float64
}

func NewMarker() *Marker {
	return &Marker{}
}

func (m *Marker) SetTimeStart(timeStart float64) {
	m.timeStart = timeStart
}
