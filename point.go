package joejson

import "encoding/json"

// GeometryTypePoint is the value for a Point's 'type' member.
const GeometryTypePoint = "Point"

// Point is a single position geometry.
type Point Position

// Raw exposes the data for this geometry as primitive types.
func (p Point) Raw() []float64 {
	return p
}

// MarshalJSON is a custom JSON marshaller.
func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	}{
		p,
		GeometryTypePoint,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (p *Point) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Position Position `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*p = Point(tmp.Position)
	return nil
}
