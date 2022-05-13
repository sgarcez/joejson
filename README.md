# JoeJSON

[![CI](https://github.com/sgarcez/joejson/actions/workflows/ci.yml/badge.svg)](https://github.com/sgarcez/joejson/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgarcez/joejson)](https://goreportcard.com/report/github.com/sgarcez/joejson)
[![Go Reference](https://pkg.go.dev/badge/github.com/sgarcez/joejson.svg)](https://pkg.go.dev/github.com/sgarcez/joejson)

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
