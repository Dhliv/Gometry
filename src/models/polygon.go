package models

type Polygon struct {
	Points []Point
}

func NewPolygon(points ...Point) *Polygon {
	return &Polygon{Points: points}
}
