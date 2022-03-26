package joejson

import "encoding/json"

// GeometryTypeMultiPolygon is the value for a MultiPolygon's 'type' member.
const GeometryTypeMultiPolygon = "MultiPolygon"

// MultiPolygon is a slice of Polygon geometries.
type MultiPolygon []Polygon

// Raw exposes the data for this geometry as primitive types.
func (p MultiPolygon) Raw() [][][][]float64 {
	out := make([][][][]float64, len(p))
	for i, pl := range p {
		out[i] = pl.Raw()
	}
	return out
}

// MarshalJSON is a custom JSON marshaller.
func (p MultiPolygon) MarshalJSON() ([]byte, error) {
	lrs := make([][]LinearRing, 0, len(p))
	for _, lr := range p {
		lrs = append(lrs, lr)
	}
	return json.Marshal(&struct {
		Polygons [][]LinearRing `json:"coordinates"`
		Type     string         `json:"type"`
	}{
		lrs,
		GeometryTypeMultiPolygon,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (p *MultiPolygon) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Polygons [][]LinearRing `json:"coordinates"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*p = make(MultiPolygon, len(tmp.Polygons))
	for i, pl := range tmp.Polygons {
		[]Polygon(*p)[i] = pl
	}

	return nil
}
