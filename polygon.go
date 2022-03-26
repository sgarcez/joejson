package joejson

import "encoding/json"

// GeometryTypePolygon is the value for a Polygon's 'type' member.
const GeometryTypePolygon = "Polygon"

// LinearRing is a closed LineString with four or more positions.
type LinearRing []Position

// Raw exposes the data for this geometry as primitive types.
func (l LinearRing) Raw() [][]float64 {
	out := make([][]float64, len(l))
	for i, pos := range l {
		out[i] = pos
	}
	return out
}

// Polygon is a slice of LinearRing.
type Polygon []LinearRing

// Raw exposes the data for this geometry as primitive types.
func (p Polygon) Raw() [][][]float64 {
	out := make([][][]float64, len(p))
	for i, lr := range p {
		out[i] = lr.Raw()
	}
	return out
}

// MarshalJSON is a custom JSON marshaller.
func (p Polygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Rings []LinearRing `json:"coordinates"`
		Type  string       `json:"type"`
	}{
		p,
		GeometryTypePolygon,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (p *Polygon) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Rings []LinearRing `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*p = tmp.Rings

	return nil
}
