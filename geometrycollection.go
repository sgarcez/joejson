package joejson

import (
	"encoding/json"
	"fmt"
)

// GeometryTypeGeometryCollection is the value for a GeometryCollection's 'type' member.
const GeometryTypeGeometryCollection = "GeometryCollection"

// GeometryCollection is a slice of Geometries.
type GeometryCollection []GeometryCollectionMember

// AppendPoint appends a Point to the collection.
func (g GeometryCollection) AppendPoint(m Point) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// AppendMultiPoint appends a MultiPoint to the collection.
func (g GeometryCollection) AppendMultiPoint(m MultiPoint) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// AppendLineString appends a LineString to the collection.
func (g GeometryCollection) AppendLineString(m LineString) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// AppendMuliLineString appends a MultiLineString to the collection.
func (g GeometryCollection) AppendMuliLineString(m MultiLineString) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// AppendPolygon appends a Polygon to the collection.
func (g GeometryCollection) AppendPolygon(m Polygon) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// AppendMultiPolygon appends a MultiPolygon to the collection.
func (g GeometryCollection) AppendMultiPolygon(m MultiPolygon) GeometryCollection {
	return append(g, GeometryCollectionMember{m})
}

// MarshalJSON is a custom JSON marshaller.
func (g GeometryCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Geometries []GeometryCollectionMember `json:"geometries"`
		Type       string                     `json:"type"`
	}{
		g,
		GeometryTypeGeometryCollection,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (g *GeometryCollection) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Geometries []json.RawMessage `json:"geometries"`
		Type       string            `json:"type"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if tmp.Type != GeometryTypeGeometryCollection {
		return fmt.Errorf("invalid type %q, expected %q", tmp.Type, GeometryTypeGeometryCollection)
	}

	for _, geom := range tmp.Geometries {
		geom, err := unmarshalGeometry(geom)
		if err != nil {
			return err
		}
		*g = append(*g, GeometryCollectionMember{geom})
	}

	return nil
}

// GeometryCollectionMember is a Geometry belonging to a GeometryCollection.
type GeometryCollectionMember struct {
	geometry any
}

// AsPoint casts the Geometry to a Point.
func (g GeometryCollectionMember) AsPoint() (Point, bool) {
	p, ok := g.geometry.(Point)
	return p, ok
}

// AsMultiPoint casts the Geometry to a MultiPoint.
func (g GeometryCollectionMember) AsMultiPoint() (MultiPoint, bool) {
	p, ok := g.geometry.(MultiPoint)
	return p, ok
}

// AsLineString casts the Geometry to a LineString.
func (g GeometryCollectionMember) AsLineString() (LineString, bool) {
	p, ok := g.geometry.(LineString)
	return p, ok
}

// AsMultiLineString casts the Geometry to a MultiLineString.
func (g GeometryCollectionMember) AsMultiLineString() (MultiLineString, bool) {
	p, ok := g.geometry.(MultiLineString)
	return p, ok
}

// AsPolygon casts the Geometry to a Polygon.
func (g *GeometryCollectionMember) AsPolygon() (Polygon, bool) {
	p, ok := g.geometry.(Polygon)
	return p, ok
}

// AsMultiPolygon casts the Geometry to a MultiPolygon.
func (g *GeometryCollectionMember) AsMultiPolygon() (MultiPolygon, bool) {
	p, ok := g.geometry.(MultiPolygon)
	return p, ok
}

// MarshalJSON is a custom JSON marshaller.
func (g GeometryCollectionMember) MarshalJSON() ([]byte, error) {
	switch v := g.geometry.(type) {
	default:
		return json.Marshal(v)
	}
}

// Type is the type of the Geometry.
func (g GeometryCollectionMember) Type() string {
	switch g.geometry.(type) {
	case Point:
		return GeometryTypePoint
	case MultiPoint:
		return GeometryTypeMultiPoint
	case LineString:
		return GeometryTypeLineString
	case MultiLineString:
		return GeometryTypeMultiLineString
	case Polygon:
		return GeometryTypePolygon
	case MultiPolygon:
		return GeometryTypeMultiPolygon
	default:
		return ""
	}
}
