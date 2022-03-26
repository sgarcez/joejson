# JoeJSON

A GeoJSON (RFC 7946) Go implementation.

## Features

- [x] JSON Marshalling / Unmarshalling
- [x] FeatureCollection
- [x] Feature
  - [x] ID
  - [x] Properties
  - [x] Bbox
    - [ ] Validation (axes order, etc)
  - [x] Geometry
    - [x] Point
    - [x] MultiPoint
    - [x] LineString
    - [x] Polygon
    - [x] MultiPolygon
    - [x] GeometryCollection
    - [ ] Validation (antimeridian crossing, right hand rule winding, etc)