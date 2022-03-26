package joejson

import "encoding/json"

// GeometryTypeLineString is the value for a LineString's 'type' member.
const GeometryTypeLineString = "LineString"

// LineString is a slice of 2 or more Positions.
type LineString []Position

// Raw exposes the data for this geometry as primitive types.
func (g LineString) Raw() [][]float64 {
	out := make([][]float64, len(g))
	for i, pt := range g {
		out[i] = pt
	}
	return out
}

// MarshalJSON is a custom JSON marshaller.
func (g LineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Positions []Position `json:"coordinates"`
		Type      string     `json:"type"`
	}{
		g,
		GeometryTypeLineString,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (g *LineString) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Positions []Position `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*g = LineString(tmp.Positions)
	return nil
}
