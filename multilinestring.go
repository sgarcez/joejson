package joejson

import "encoding/json"

// GeometryTypeMultiLineString is the value for a MultiLineString's 'type' member.
const GeometryTypeMultiLineString = "MultiLineString"

// MultiLineString is a slice of LineStrings.
type MultiLineString []LineString

// Raw exposes the data for this geometry as primitive types.
func (g MultiLineString) Raw() [][][]float64 {
	out := make([][][]float64, len(g))
	for i, ls := range g {
		out[i] = ls.Raw()
	}
	return out
}

// MarshalJSON is a custom JSON marshaller.
func (g MultiLineString) MarshalJSON() ([]byte, error) {
	positions := make([][]Position, 0, len(g))
	for _, ls := range g {
		positions = append(positions, ls)
	}
	return json.Marshal(&struct {
		LineStrings [][]Position `json:"coordinates"`
		Type        string       `json:"type"`
	}{
		positions,
		GeometryTypeMultiLineString,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (g *MultiLineString) UnmarshalJSON(b []byte) error {
	var tmp struct {
		LineStrings [][]Position `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*g = make([]LineString, len(tmp.LineStrings))
	for i, pl := range tmp.LineStrings {
		[]LineString(*g)[i] = pl
	}

	return nil
}
