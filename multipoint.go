package joejson

import (
	"encoding/json"
	"fmt"
)

// GeometryTypeMultiPoint is the value for a MultiPoint's 'type' member.
const GeometryTypeMultiPoint = "MultiPoint"

// MultiPoint is a slice of position geometries.
type MultiPoint LineString

// Raw exposes the data for this geometry as primitive types.
func (g MultiPoint) Raw() [][]float64 {
	out := make([][]float64, len(g))
	for i, pt := range g {
		out[i] = pt
	}
	return out
}

// MarshalJSON is a custom JSON marshaller.
func (g MultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Positions []Position `json:"coordinates"`
		Type      string     `json:"type"`
	}{
		g,
		GeometryTypeMultiPoint,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (g *MultiPoint) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Positions []Position `json:"coordinates"`
		Type      string     `json:"type"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if tmp.Type != GeometryTypeMultiPoint {
		return fmt.Errorf("invalid type %q, expected %q", tmp.Type, GeometryTypeMultiPoint)
	}

	*g = MultiPoint(tmp.Positions)
	return nil
}
