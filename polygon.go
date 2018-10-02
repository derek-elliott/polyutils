package polyutils

type Point struct {
	X float64
	Y float64
}

type BoundingBox struct {
	Max Point
	Min Point
}

type Polygon struct {
	Points      []Point
	BoundingBox BoundingBox
	XVerts      []float64
	YVerts      []float64
}

// NewPolygon returns a pointer to a polygon given a set of points.
func NewPolygon(points []Point) *Polygon {
	xVerts, yVerts := []float64{}, []float64{}
	for _, point := range points {
		xVerts = append(xVerts, point.X)
		yVerts = append(yVerts, point.Y)
	}
	xMin, xMax := getMinMax(xVerts)
	yMin, yMax := getMinMax(yVerts)
	boundingBox := BoundingBox{Point{xMax, yMax}, Point{xMin, yMin}}
	return &Polygon{points, boundingBox, xVerts, yVerts}
}

// ArrayToPoints converts an array of float64 into a list of points
func ArrayToPoints(rawPoints [][]float64) *[]Point {
	var points []Point
	for _, item := range rawPoints {
		points = append(points, Point{item[0], item[1]})
	}
	return &points
}

// Helper function to get the min and max from an array of numbers
func getMinMax(data []float64) (min float64, max float64) {
	min, max = data[0], data[0]
	for _, i := range data {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}
	return
}

// Contains returns true if the given point is contained in the polygon
func (p *Polygon) Contains(point Point) bool {
	if !p.BoundingBox.Contains(point) {
		return false
	}
	numVert := len(p.Points)
	contains := false
	for i, j := 0, numVert-1; i < numVert; j, i = i, i+1 {
		if (p.YVerts[i] > point.Y) != (p.YVerts[j] > point.Y) && (point.X < (p.XVerts[j]-p.XVerts[i])*(point.Y-p.YVerts[i])/(p.YVerts[j]-p.YVerts[i])+p.XVerts[i]) {
			contains = !contains
		}
	}
	return contains
}

// InPolygon returns true if the given polygon contains the point
func (p *Point) InPolygon(poly Polygon) bool {
	return poly.Contains(*p)
}

// Contains returns true if the given point is in the bounding box
func (bbox *BoundingBox) Contains(p Point) bool {
	if (p.X <= bbox.Max.X) && (p.X >= bbox.Min.X) && (p.Y <= bbox.Max.Y) && (p.Y >= bbox.Min.Y) {
		return true
	} else {
		return false
	}
}

// InBoundingBox returns true if the given bounding box contains the point
func (p *Point) InBoundingBox(bbox BoundingBox) bool {
	return bbox.Contains(*p)
}
