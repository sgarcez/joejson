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
			ft:   Feature{}.WithPoint(&Point{-170.0, 40.0}),
			json: `{"type":"Feature","geometry":{"coordinates":[-170,40],"type":"Point"}}`,
			raw:  []float64{-170, 40},
		},
		"Point with Properties": {
			ft: Feature{
				ID: "abc",
				Properties: map[string]any{
					"foo": "bar",
				},
			}.WithPoint(&Point{-170.0, 40.0}),
			json: `{"id":"abc","type":"Feature","geometry":{"coordinates":[-170,40],"type":"Point"},"properties":{"foo":"bar"}}`,
		},
		"MultiPoint": {
			ft: Feature{}.WithMultiPoint(
				&MultiPoint{
					{-170.0, 40.0},
				}),
			raw:  [][]float64{{-170, 40}},
			json: `{"type":"Feature","geometry":{"coordinates":[[-170,40]],"type":"MultiPoint"}}`,
		},
		"LineString": {
			ft: Feature{}.WithLineString(
				&LineString{
					{-170.0, 40.0},
				}),
			raw:  [][]float64{{-170, 40}},
			json: `{"type":"Feature","geometry":{"coordinates":[[-170,40]],"type":"LineString"}}`,
		},
		"MultiLineString": {
			ft: Feature{}.WithMultiLineString(
				&MultiLineString{
					{
						{-170.0, 40.0},
					},
				}),
			raw:  [][][]float64{{{-170, 40}}},
			json: `{"type":"Feature","geometry":{"coordinates":[[[-170,40]]],"type":"MultiLineString"}}`,
		},
		"Polygon": {
			ft: Feature{}.WithPolygon(
				&Polygon{
					{
						{-170.0, 40.0},
					},
				}),
			raw:  [][][]float64{{{-170, 40}}},
			json: `{"type":"Feature","geometry":{"coordinates":[[[-170,40]]],"type":"Polygon"}}`,
		},
		"MultiPolygon": {
			ft: Feature{}.WithMultiPolygon(
				&MultiPolygon{
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
					&Point{-170.0, 40.0},
				).AppendPolygon(
					&Polygon{
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
				assert.IsType(t, &Point{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiPoint:
				got, ok := tt.ft.AsMultiPoint()
				assert.True(t, ok)
				assert.IsType(t, &MultiPoint{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeLineString:
				got, ok := tt.ft.AsLineString()
				assert.True(t, ok)
				assert.IsType(t, &LineString{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiLineString:
				got, ok := tt.ft.AsMultiLineString()
				assert.True(t, ok)
				assert.IsType(t, &MultiLineString{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypePolygon:
				got, ok := tt.ft.AsPolygon()
				assert.True(t, ok)
				assert.IsType(t, &Polygon{}, got)
				if tt.raw != nil {
					assert.Equal(t, tt.raw, got.Raw())
				}
			case GeometryTypeMultiPolygon:
				got, ok := tt.ft.AsMultiPolygon()
				assert.True(t, ok)
				assert.IsType(t, &MultiPolygon{}, got)
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

func TestFeatureIDJSONMarshall(t *testing.T) {
	testCases := map[string]struct {
		ft   Feature
		json string
		err  string
	}{
		"string": {
			ft:   Feature{ID: "1"}.WithPoint(Point{0, 0}),
			json: `{"id":"1","type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"uint8": {
			ft:   Feature{ID: uint8(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"uint16": {
			ft:   Feature{ID: uint16(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"uint32": {
			ft:   Feature{ID: uint32(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"uint64": {
			ft:   Feature{ID: uint64(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"int8": {
			ft:   Feature{ID: int8(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"int16": {
			ft:   Feature{ID: int16(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"int32": {
			ft:   Feature{ID: int32(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"int64": {
			ft:   Feature{ID: int64(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"float32": {
			ft:   Feature{ID: float32(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"float64": {
			ft:   Feature{ID: float64(1)}.WithPoint(Point{0, 0}),
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
		},
		"other type than int or string": {
			ft:  Feature{ID: true}.WithPoint(Point{0, 0}),
			err: `json: error calling MarshalJSON for type joejson.Feature: invalid type "bool" for id, expected string or numeric`,
		},
	}

	t.Parallel()
	for name, tt := range testCases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			bs, err := json.Marshal(tt.ft)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.json, string(bs))
			}
		})
	}
}

func TestFeatureIDJSONUnmarshall(t *testing.T) {
	testCases := map[string]struct {
		json string
		ft   Feature
		err  string
	}{
		"string": {
			json: `{"id":"1","type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
			ft:   Feature{ID: "1"}.WithPoint(Point{0, 0}),
		},
		"integer": {
			json: `{"id":1,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
			ft:   Feature{ID: float64(1)}.WithPoint(Point{0, 0}),
		},
		"decimal": {
			json: `{"id":1.5,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
			ft:   Feature{ID: float64(1.5)}.WithPoint(Point{0, 0}),
		},
		"other type than numeric or string": {
			json: `{"id":true,"type":"Feature","geometry":{"coordinates":[0,0],"type":"Point"}}`,
			err:  `invalid type "bool" for id, expected string or numeric`,
		},
	}

	t.Parallel()
	for name, tt := range testCases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var ft Feature
			err := json.Unmarshal([]byte(tt.json), &ft)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.ft, ft)
			}
		})
	}
}
