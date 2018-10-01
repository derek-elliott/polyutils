package polyutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PolygonTestCase struct {
	Polygon Polygon
	Point   Point
	Target  bool
	Message string
}

type BBoxTestCase struct {
	BBox    BoundingBox
	Point   Point
	Target  bool
	Message string
}

type PointTestCase struct {
	Point   Point
	Geo     interface{}
	Target  bool
	Message string
}

type PolygonTestCases []PolygonTestCase

type BBoxTestCases []BBoxTestCase

type PointTestCases []PointTestCase

var (
	polyCases      PolygonTestCases
	bboxCases      BBoxTestCases
	bboxPointCases PointTestCases
	polyPointCases PointTestCases
)

func init() {
	convexPolygon := NewPolygon([]Point{Point{0, 0}, Point{1, 0}, Point{2, 1}, Point{2, 2}, Point{1, 3}, Point{0, 3}, Point{-1, 2}, Point{-1, 1}, Point{0, 0}})

	concavePolygon := NewPolygon([]Point{Point{0, 0}, Point{3, 0}, Point{2, 2}, Point{3, 4}, Point{0, 4}, Point{0, 0}})

	polyCases = PolygonTestCases{
		PolygonTestCase{
			*convexPolygon,
			Point{0.5, 1.5},
			true,
			"Convex Polygon contains point",
		},
		PolygonTestCase{
			*convexPolygon,
			Point{0, 0},
			true,
			"Convex Polygon has point on edge",
		},
		PolygonTestCase{
			*convexPolygon,
			Point{0, -1},
			false,
			"Convex polygon does not contain point",
		},
		PolygonTestCase{
			*concavePolygon,
			Point{1, 2},
			true,
			"Concave polygon contains point",
		},
		PolygonTestCase{
			*concavePolygon,
			Point{0, 0},
			true,
			"Concave polygon has point on edge",
		},
		PolygonTestCase{
			*concavePolygon,
			Point{3, 2},
			false,
			"Concave polygon does not contain point (inside bounding box)",
		},
		PolygonTestCase{
			*concavePolygon,
			Point{4, 2},
			false,
			"Concave polygon does not contain point (outside bounding box)",
		},
	}

	bboxCases = BBoxTestCases{
		BBoxTestCase{
			BoundingBox{Point{1, 1}, Point{-1, -1}},
			Point{0, 0},
			true,
			"Bounding box contains point",
		},
		BBoxTestCase{
			BoundingBox{Point{1, 1}, Point{-1, -1}},
			Point{1, 0},
			true,
			"Bounding box has point on edge",
		},
		BBoxTestCase{
			BoundingBox{Point{1, 1}, Point{-1, -1}},
			Point{2, 0},
			false,
			"Bounding box doesn't contain point",
		},
	}

	bboxPointCases = PointTestCases{
		PointTestCase{
			Point{0, 0},
			BoundingBox{Point{1, 1}, Point{-1, -1}},
			true,
			"Point in bounding box",
		},
		PointTestCase{
			Point{10, 10},
			BoundingBox{Point{1, 1}, Point{-1, -1}},
			false,
			"Point not in bounding box",
		},
	}

	polyPointCases = PointTestCases{
		PointTestCase{
			Point{0, 0},
			*convexPolygon,
			true,
			"Point in polygon",
		},
		PointTestCase{
			Point{10, 10},
			*convexPolygon,
			false,
			"Point not in polygon",
		},
	}
}

func TestPolyContains(t *testing.T) {
	assert := assert.New(t)

	for _, test := range polyCases {
		result := test.Polygon.Contains(test.Point)
		assert.Equal(test.Target, result, test.Message)
	}
}

func TestBBoxContains(t *testing.T) {
	assert := assert.New(t)

	for _, test := range bboxCases {
		result := test.BBox.Contains(test.Point)
		assert.Equal(test.Target, result, test.Message)
	}
}

func TestPointInPolygon(t *testing.T) {
	assert := assert.New(t)

	for _, test := range polyPointCases {
		result := test.Point.InPolygon(test.Geo.(Polygon))
		assert.Equal(test.Target, result, test.Message)
	}
}

func TestPointInBBox(t *testing.T) {
	assert := assert.New(t)

	for _, test := range bboxPointCases {
		result := test.Point.InBoundingBox(test.Geo.(BoundingBox))
		assert.Equal(test.Target, result, test.Message)
	}
}

func TestFromGeoJSON(t *testing.T) {
	assert := assert.New(t)
	rawJSON := `{
		  "type": "Feature",
		  "geometry": {
		     "type": "Polygon",
		     "coordinates": [
		         [ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ]
		     ]
		  },
		  "properties": {
		    "name": "Dinagat Islands"
		  }
		}`

	targetPoly := NewPolygon([]Point{Point{100.0, 0.0}, Point{101.0, 0.0}, Point{101.0, 1.0}, Point{100.0, 1.0}, Point{100.0, 0.0}})
	poly, err := FromGeoJSON([]byte(rawJSON))
	assert.NoError(err)
	assert.Equal(poly, *targetPoly, "Should be equal")
}
