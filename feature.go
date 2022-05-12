package joejson

import (
	"encoding/json"
	"fmt"
)

// TypeFeature is the value for a Feature's 'type' member.
const TypeFeature string = "Feature"

// Feature represents a spatially bounded 'thing'.
type Feature struct {
	// ID is an optional Feature identifier ('id').
	ID any
	//  Properties is an optional JSON object ('properties').
	Properties map[string]any
	// Bbox optionally includes information on the coordinate range for the Feature Geometry.
	Bbox []Position
	// geometry is an unexported field representing one of
	// (Point|LineString|MultiPoint|MultiLineString|Polygon|MultiPolygon|GeometryCollection).
	geometry any
}

// GeometryType is the type of the Feature's Geometry.
func (f Feature) GeometryType() string {
	switch f.geometry.(type) {
	case *Point:
		return GeometryTypePoint
	case *MultiPoint:
		return GeometryTypeMultiPoint
	case *LineString:
		return GeometryTypeLineString
	case *MultiLineString:
		return GeometryTypeMultiLineString
	case *Polygon:
		return GeometryTypePolygon
	case *MultiPolygon:
		return GeometryTypeMultiPolygon
	case GeometryCollection:
		return GeometryTypeGeometryCollection
	default:
		return ""
	}
}

// WithPoint sets the Feature's Geometry to the provided Point.
func (f Feature) WithPoint(g *Point) Feature {
	f.geometry = g
	return f
}

// AsPoint casts the Feature's Geometry to a Point.
func (f Feature) AsPoint() (*Point, bool) {
	p, ok := f.geometry.(*Point)
	return p, ok
}

// WithMultiPoint sets the Feature's Geometry to the provided MultiPoint.
func (f Feature) WithMultiPoint(g *MultiPoint) Feature {
	f.geometry = g
	return f
}

// AsMultiPoint casts the Feature's Geometry to a MultiPoint.
func (f Feature) AsMultiPoint() (*MultiPoint, bool) {
	p, ok := f.geometry.(*MultiPoint)
	return p, ok
}

// WithLineString sets the Feature's Geometry to the provided LineString.
func (f Feature) WithLineString(g *LineString) Feature {
	f.geometry = g
	return f
}

// AsLineString casts the Feature's Geometry to a LineString.
func (f Feature) AsLineString() (*LineString, bool) {
	p, ok := f.geometry.(*LineString)
	return p, ok
}

// WithMultiLineString sets the Feature's Geometry to the provided LineMultiString.
func (f Feature) WithMultiLineString(g *MultiLineString) Feature {
	f.geometry = g
	return f
}

// AsMultiLineString casts the Feature's Geometry to a MultiLineString.
func (f Feature) AsMultiLineString() (*MultiLineString, bool) {
	p, ok := f.geometry.(*MultiLineString)
	return p, ok
}

// WithPolygon sets the Feature's Geometry to the provided Polygon.
func (f Feature) WithPolygon(g *Polygon) Feature {
	f.geometry = g
	return f
}

// AsPolygon casts the Feature's Geometry to a Polygon.
func (f Feature) AsPolygon() (*Polygon, bool) {
	p, ok := f.geometry.(*Polygon)
	return p, ok
}

// WithMultiPolygon sets the Feature's Geometry to the provided MultiPolygon.
func (f Feature) WithMultiPolygon(g *MultiPolygon) Feature {
	f.geometry = g
	return f
}

// AsMultiPolygon casts the Feature's Geometry to a MultiPolygon.
func (f *Feature) AsMultiPolygon() (*MultiPolygon, bool) {
	p, ok := f.geometry.(*MultiPolygon)
	return p, ok
}

// WithGeometryCollection sets the Feature's Geometry to the provided GeometryCollection.
func (f Feature) WithGeometryCollection(g GeometryCollection) Feature {
	f.geometry = g
	return f
}

// AsGeometryCollection casts Feature's Geometry to a GeometryCollection.
func (f *Feature) AsGeometryCollection() (GeometryCollection, bool) {
	p, ok := f.geometry.(GeometryCollection)
	return p, ok
}

// MarshalJSON is a custom JSON marshaller.
func (f Feature) MarshalJSON() ([]byte, error) {
	switch f.ID.(type) {
	case string, uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, nil:
	default:
		return nil, fmt.Errorf(`invalid type "%T" for id, expected string or numeric`, f.ID)
	}

	return json.Marshal(&struct {
		ID         any            `json:"id,omitempty"`
		Type       string         `json:"type"`
		Geometry   any            `json:"geometry"`
		Properties map[string]any `json:"properties,omitempty"`
		BBox       []Position     `json:"bbox,omitempty"`
	}{
		f.ID,
		TypeFeature,
		f.geometry,
		f.Properties,
		f.Bbox,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (f *Feature) UnmarshalJSON(b []byte) error {
	var tmp struct {
		ID         any             `json:"id"`
		Geometry   json.RawMessage `json:"geometry"`
		Properties map[string]any  `json:"properties"`
		Bbox       []Position      `json:"bbox"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	switch tmp.ID.(type) {
	case string, float64, nil:
	default:
		return fmt.Errorf(`invalid type "%T" for id, expected string or numeric`, tmp.ID)
	}

	f.ID = tmp.ID
	f.Properties = tmp.Properties
	f.Bbox = tmp.Bbox

	var err error
	f.geometry, err = unmarshalGeometry(tmp.Geometry)
	return err
}

func unmarshalGeometry(bs []byte) (any, error) {
	var tmp struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(bs, &tmp); err != nil {
		return nil, err
	}

	switch tmp.Type {
	case GeometryTypePoint:
		var g Point
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypeMultiPoint:
		var g MultiPoint
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypeLineString:
		var g LineString
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypeMultiLineString:
		var g MultiLineString
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypePolygon:
		var g Polygon
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypeMultiPolygon:
		var g MultiPolygon
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return &g, nil
	case GeometryTypeGeometryCollection:
		var g GeometryCollection
		if err := json.Unmarshal(bs, &g); err != nil {
			return nil, err
		}
		return g, nil
	default:
		return nil, fmt.Errorf("unknown geometry type: %q", tmp.Type)
	}
}
