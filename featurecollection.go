package joejson

import (
	"encoding/json"
	"fmt"
)

// TypeFeatureCollection is the value for a FeatureCollection's 'type' member.
const TypeFeatureCollection string = "FeatureCollection"

// FeatureCollection is a collection of Features.
type FeatureCollection struct {
	Features []Feature
	Bbox     []Position
}

// MarshalJSON is a custom JSON marshaller.
func (f FeatureCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     string     `json:"type"`
		Features []Feature  `json:"features"`
		BBox     []Position `json:"bbox,omitempty"`
	}{
		TypeFeatureCollection,
		f.Features,
		f.Bbox,
	})
}

// UnmarshalJSON is a custom JSON unmarshaller.
func (f *FeatureCollection) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Type     string     `json:"type"`
		Features []Feature  `json:"features"`
		Bbox     []Position `json:"bbox"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if tmp.Type != TypeFeatureCollection {
		return fmt.Errorf("invalid type %q, expected %q", tmp.Type, TypeFeatureCollection)
	}

	f.Features = tmp.Features
	f.Bbox = tmp.Bbox
	return nil
}
