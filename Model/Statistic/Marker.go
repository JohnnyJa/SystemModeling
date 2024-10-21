package Statistic

type Marker struct {
	TimeStart float64
	TimeEnd   float64
}

func NewMarker(timeStart float64) *Marker {
	return &Marker{
		TimeStart: timeStart,
	}
}

func (m *Marker) SetTimeEnd(timeEnd float64) {
	m.TimeEnd = timeEnd
}

func (m *Marker) GetTimeStart() float64 {
	return m.TimeStart
}
