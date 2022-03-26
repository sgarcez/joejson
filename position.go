package joejson

// Position is the fundamental geometry construct.
type Position []float64

// Lon is the Longitude coordinate.
func (p Position) Lon() float64 {
	if len(p) < 1 {
		return 0
	}
	return p[0]
}

// Lat is the latitude coordinate.
func (p Position) Lat() float64 {
	if len(p) < 2 {
		return 0
	}
	return p[1]
}

// Elevation is an optional third Position element.
func (p Position) Elevation() float64 {
	if len(p) < 3 {
		return 0
	}
	return p[2]
}
