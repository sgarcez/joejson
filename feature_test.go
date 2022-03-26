package joejson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeature(t *testing.T) {
	testCases := map[string]struct {
		ft           Feature
		json         string
		unmarshalErr string
		raw          any
	}{
		"Empty": {
			ft:           Feature{},
			json:         `{"type":"Feature","geometry":null}`,
			unmarshalErr: `unknown geometry type: ""`,
		},
		"Point": {
			ft:   Feature{}.WithPoint(Point{-170.0, 40.0}),
			json: `{"type":"Feature","geometry":{"coordinates":[-170,40],"type":"Point"}}`,
			raw:  []float64{-170, 40},
		},
		"Point with Properties": {
			ft: Feature{
				ID: "abc",
				Properties: map[string]any{
					"foo": "bar",
				},
			}.WithPoint(Point{-170.0, 40.0}),
			json: `{"id":"abc","type":"Feature","geometry":{"coordinates":[-170,40],"type":"Point"},"properties":{"foo":"bar"}}`,
		},
		"MultiPoint": {
			ft: Feature{}.WithMultiPoint(
				MultiPoint{
					{-170.0, 40.0},
				}),
			raw:  [][]float64{{-170, 40}},
			json: `{"type":"Feature","geometry":{"coordinates":[[-170,40]],"type":"MultiPoint"}}`,
		},
		"LineString": {
			ft: Feature{}.WithLineString(
				LineString{
					{-170.0, 40.0},
				}),
			raw:  [][]float64{{-170, 40}},
			json: `{"type":"Feature","geometry":{"coordinates":[[-170,40]],"type":"LineString"}}`,
		},
		"MultiLineString": {
			ft: Feature{}.WithMultiLineString(
				MultiLineString{
					{
						{-170.0, 40.0},
					},
				}),
			raw:  [][][]float64{{{-170, 40}}},
			json: `{"type":"Feature","geometry":{"coordinates":[[[-170,40]]],"type":"MultiLineString"}}`,
		},
		"Polygon": {
			ft: Feature{}.WithPolygon(
				Polygon{
					{
						{-170.0, 40.0},
					},
				}),
			raw:  [][][]float64{{{-170, 40}}},
			json: `{"type":"Feature","geometry":{"coordinates":[[[-170,40]]],"type":"Polygon"}}`,
		},
		"MultiPolygon": {
			ft: Feature{}.WithMultiPolygon(
				MultiPolygon{
					{
						{
							{-170.0, 40.0},
						},
					},
				}),
			raw:  [][][][]float64{{{{-170, 40}}}},
			json: `{"type":"Feature","geometry":{"coordinates":[[[[-170,40]]]],"type":"MultiPolygon"}}`,
		},
		"GeometryCollection": {
			ft: Feature{}.WithGeometryCollection(
				GeometryCollection{}.AppendPoint(
					Point{-170.0, 40.0},
				).AppendPolygon(
					Polygon{
						{
							{-170.0, 40.0},
						},
					},
				)),
			json: `{"type":"Feature","geometry":{"geometries":[{"coordinates":[-170,40],"type":"Point"},{"coordinates":[[[-170,40]]],"type":"Polygon"}],"type":"GeometryCollection"}}`,
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			// encode
			bs, err := json.Marshal(tt.ft)
			assert.NoError(t, err)
			assert.Equal(t, tt.json, string(bs))

			// decode
			var unmarshalledFeature Feature
			err = json.Unmarshal(bs, &unmarshalledFeature)
			if tt.unmarshalErr != "" {
				assert.EqualError(t, err, tt.unmarshalErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.ft, unmarshalledFeature)
			}

			// cast
			switch tt.ft.GeometryType() {
			case "":
				assert.Nil(t, tt.ft.geometry)
			case GeometryTypePoint:
				got, ok := tt.ft.AsPoint()
				assert.True(t, ok)
				assert.IsType(t, Point{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiPoint:
				got, ok := tt.ft.AsMultiPoint()
				assert.True(t, ok)
				assert.IsType(t, MultiPoint{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeLineString:
				got, ok := tt.ft.AsLineString()
				assert.True(t, ok)
				assert.IsType(t, LineString{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiLineString:
				got, ok := tt.ft.AsMultiLineString()
				assert.True(t, ok)
				assert.IsType(t, MultiLineString{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypePolygon:
				got, ok := tt.ft.AsPolygon()
				assert.True(t, ok)
				assert.IsType(t, Polygon{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiPolygon:
				got, ok := tt.ft.AsMultiPolygon()
				assert.True(t, ok)
				assert.IsType(t, MultiPolygon{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeGeometryCollection:
				got, ok := tt.ft.AsGeometryCollection()
				assert.True(t, ok)
				assert.IsType(t, GeometryCollection{}, got)
			default:
				assert.Fail(t, "unhandled geometry type: %s", tt.ft.GeometryType())
			}
		})
	}
}
