package models

import (
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

type Polygon struct {
	Points *[]*Point
}

func NewPolygon(points ...*Point) *Polygon {
	return &Polygon{Points: &points}
}

// Transform polygon to string literal in form [(x1, y1), ..., (xn, yn)].
func (P *Polygon) ToString() string {
	var s string = "["

	for i := 0; i < len(*P.Points); i++ {
		if i > 0 {
			s += ", "
		}
		s += (*P.Points)[i].ToString()
	}

	s += "]"
	return s
}

/*
Evaluates wheter polygon 'P' its equal to polygon 'P2'. This functions takes note of polygon
rotations.
*/
func (P *Polygon) Equal(P2 *Polygon, ignoreRotation bool) bool {
	if len(*P.Points) != len(*P2.Points) {
		return false
	}

	var n int = len(*P.Points)
	var pos int = 0
	for pos < n && ignoreRotation { // P2 could be P rotated by 'pos' positions.
		if (*P.Points)[0].Equal((*P2.Points)[pos]) {
			break
		}
		pos++
	}

	// Evaluate if every single point are equal.
	for i := 0; i < n; i++ {
		j := (pos + i) % n
		if !(*P.Points)[i].Equal((*P2.Points)[j]) {
			return false
		}
	}

	return true
}

/*
Given 2 contiguous segments, if they are collinear the 2 segments will be joined in only one.
Namely, given 2 segments AB and BC, if they are collinear those 2 segments now will be just one: AC.
*/
func (P *Polygon) DeleteCollinearSegments() {
	var pos int = 0
	var n int = len(*P.Points)
	var A, B, C *Point

	if n <= 2 {
		// P is not a polygon, its a line.
		return
	}

	for i := 0; i < n; i++ {
		A = (*P.Points)[(i-1+n)%n]
		B = (*P.Points)[i]
		C = (*P.Points)[(i+1)%n]

		if GetOrientation(A, B, C) == COLLINEAR_ORIENTATION {
			continue
		}

		pos = i
		break
	}

	var neoPolygon []*Point = make([]*Point, 1)
	neoPolygon[0] = (*P.Points)[pos]

	for i := pos + 1; i < n; i++ {
		A = neoPolygon[len(neoPolygon)-1]
		B = (*P.Points)[i]
		C = (*P.Points)[(i+1)%n]

		if GetOrientation(A, B, C) == COLLINEAR_ORIENTATION {
			continue
		}

		neoPolygon = append(neoPolygon, B)
	}

	P.Points = &neoPolygon
}

// Does a cyclic rotation on 'Points' array by 'times' position to left.
func (P *Polygon) CyclicRotation(times int) {
	var neoPolygon []*Point = make([]*Point, 0)
	if times == 0 {
		return
	}

	var n int = len(*P.Points)
	for i := 0; i < n; i++ {
		j := (i + times) % n
		neoPolygon = append(neoPolygon, (*P.Points)[j])
	}

	P.Points = &neoPolygon
}

// Determines wheter point 'P' is on the perimeter of 'polygon'.
func (Pol *Polygon) PointInPolygonPerimeter(P *Point) bool {
	var line *Line
	var A, B *Point
	var n int = len(*Pol.Points)

	for i := 0; i < n; i++ {
		A = (*Pol.Points)[i]
		B = (*Pol.Points)[(i+1)%n]
		line = NewLine(A, B)

		if line.PointOnSegment(P) {
			return true
		}
	}

	return false
}

/*
Calculates the signed area of 'polygon'. Area would be < 0 when polygon is
given in clockwise order. We could make this function ignore that order
to always return the positive area of polygon.
*/
func (P *Polygon) Area(ignoreClockwiseOrder bool) float64 {
	var area float64 = 0
	var n int = len(*P.Points)
	var A, B *Point

	for i := 0; i < n; i++ {
		A = (*P.Points)[i]
		B = (*P.Points)[(i+1)%n]

		area += (A.X*B.Y - B.X*A.Y) / float64(2)
	}

	if ignoreClockwiseOrder {
		area = math.Abs(area)
	}

	return area
}

// Sorts 'polygon' in clockwise order.
func (P *Polygon) SortClockwise() {
	var area float64 = P.Area(false)

	if area < 0.0 { // Polygon is already sorted clockwise.
		return
	}

	var sortedPolygon []*Point = make([]*Point, 0)
	for i := len(*P.Points) - 1; i >= 0; i-- {
		sortedPolygon = append(sortedPolygon, (*P.Points)[i])
	}

	P.Points = &sortedPolygon
}

// Deletes contiguos points that are equal.
func (P *Polygon) DeleteRepeatedPoints() {
	var neoPolygon []*Point = make([]*Point, 0)
	var A, B *Point
	var n int = len(*P.Points)

	for i := 0; i < n; i++ {
		A = (*P.Points)[i]
		B = (*P.Points)[(i+1)%n]

		if A.Equal(B) {
			continue
		}

		neoPolygon = append(neoPolygon, A)
	}

	P.Points = &neoPolygon
}

// Determines whether or not some point 'P' is inside the polygon.
func (Pol *Polygon) PointInPolygon(P *Point) bool {
	var angle_sum float64 = 0
	var A, B *Point
	var n int = len(*Pol.Points)
	var line, AP, PB *Line
	var orientation int8

	for i := 0; i < n; i++ {
		A = (*Pol.Points)[i]
		B = (*Pol.Points)[(i+1)%n]
		line = NewLine(A, B)
		if line.PointOnSegment(P) {
			return true
		}

		orientation = GetOrientation(A, P, B)
		AP, PB = NewLine(A, P), NewLine(P, B)
		if orientation == COUNTERCLOCKWISE_ORIENTATION {
			angle_sum += AP.AngleBetweenSegments(PB)
		} else {
			angle_sum -= AP.AngleBetweenSegments(PB)
		}
	}

	return math.Abs(angle_sum) > math.Pi
}

// Returns wheter current polygon is inside 'PObjetive' polygon.
//
// 🚧 Take note that you should check if polygons have intersection before using this method, 'cause this method only works
// if 'PObjetive' is convex or if there is not an intersection between current and 'PObjetive' polygons. 🚧
func (P *Polygon) InsidePolygon(PObjetive *Polygon) bool {
	var inside bool = true

	for _, p := range *P.Points {
		if !PObjetive.PointInPolygon(p) {
			inside = false
			break
		}
	}

	return inside
}

// Returns a copy of polygon.
func (P *Polygon) Copy() *Polygon {
	var neoPoints []*Point

	for _, p := range *P.Points {
		neoPoints = append(neoPoints, p.Copy())
	}

	return NewPolygon(neoPoints...)
}
