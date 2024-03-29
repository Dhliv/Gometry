package models

import (
	"fmt"
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

/*
? Separate the concept of a Line of the concept of Segment?
*/
type Line struct {
	A *Point
	B *Point
}

func NewLine(A, B *Point) *Line {
	return &Line{A: A.Copy(), B: B.Copy()}
}

func (L Line) String() string {
	return fmt.Sprintf("(%v - %v)", L.A, L.B)
}

/*
Returns the slope of the line
? Report when the slope is infinite?
*/
func (L *Line) Slope() float64 {
	return (L.A.Y - L.B.Y) / (L.A.X - L.B.X)
}

/*
Returns the inverted slope of the line, that means ortogonal slope

? Report when the inverted slope is infinite?
*/
func (L *Line) InvSlope() float64 {
	return -1 * (L.A.X - L.B.X) / (L.A.Y - L.B.Y)
}

func (L *Line) SlopeVector() *Point {
	var X float64 = L.B.X - L.A.X
	var Y float64 = L.B.Y - L.A.Y
	return NewPoint(X, Y)
}

/*
Calculates the point inside line by specifying 'y' coordinate. It assumes that Line L is not horizontal (L.Slope() == 0)
*/
func (L *Line) GetPointInsideLineByYCoordinate(y float64) *Point {
	var hLine *Line = &Line{NewPoint(0, y), NewPoint(1, y)}

	return L.IntersectionPointOnALine(hLine)
}

/*
Return a Point D like Line CD is orthogonal to Line AB. That means Line(A,B) * Line(C,D) == 0
*/
func (L *Line) OrthogonalLinePoint(C *Point) *Point {
	var x, y float64

	if math.Abs(L.A.X-L.B.X) < EPSILON {
		return NewPoint(L.A.X, C.Y)
	}
	if math.Abs(L.A.Y-L.B.Y) < EPSILON {
		return NewPoint(C.X, L.A.Y)
	}

	x = (L.A.Y - C.Y + L.InvSlope()*C.X - L.Slope()*L.A.X) / (L.InvSlope() - L.Slope())
	y = C.Y + L.InvSlope()*(x-C.X)

	return NewPoint(x, y)
}

/*
Determines whether 'C' is in the line AB or not, returing true or false respectly.
*/
func (L *Line) PointInLine(C *Point) bool {
	var a, b, c, result float64
	a = L.B.Y - L.A.Y
	b = L.A.X - L.B.X
	c = L.A.Y*L.B.X - L.B.Y*L.A.X
	result = a*C.X + b*C.Y + c

	return math.Abs(result) < EPSILON
}

/*
Given a Line L and a Point P, return a Line R that is parallel to L and P is in that line.
*/
func (L *Line) ParallelLineOfALineAndAPoint(P *Point) *Line {
	var slope *Point = L.SlopeVector()
	var p1 *Point = NewPoint(P.X-slope.X*10, P.Y-slope.Y*10)
	var p2 *Point = NewPoint(P.X+slope.X*10, P.Y+slope.Y*10)
	return &Line{p1, p2}
}

/*
Calculate and return the intersection of line AB and line CD. The function works on the premise
that AB and CD are not collinear.

? REPORT IF ITS COLLINEAR ? Throw an exception?
? REPORT IF BOTH ARE PARALLEL ? Throw an exception?
*/
func (AB *Line) IntersectionPointOnALine(CD *Line) *Point {
	// Line AB represented as a1x + b1y = c1
	var a1, b1, c1, a2, b2, c2, determinant, x, y float64
	a1 = AB.B.Y - AB.A.Y
	b1 = AB.A.X - AB.B.X
	c1 = a1*(AB.A.X) + b1*(AB.A.Y)

	// Line CD represented as a2x + b2y = c2
	a2 = CD.B.Y - CD.A.Y
	b2 = CD.A.X - CD.B.X
	c2 = a2*(CD.A.X) + b2*(CD.A.Y)

	determinant = a1*b2 - a2*b1
	x = (b2*c1 - b1*c2) / determinant
	y = (a1*c2 - a2*c1) / determinant
	return NewPoint(x, y)
}

/*
	Determines whether or not the point 'P' is on the segment defined by AB.
*/

func (L *Line) PointOnSegment(P *Point) bool {
	var a, b, c, x_min, x_max, y_min, y_max float64
	a = L.A.Y - L.B.Y
	b = L.B.X - L.A.X
	c = L.A.X*L.B.Y - L.A.Y*L.B.X

	if math.Abs(a*P.X+b*P.Y+c) < EPSILON {
		x_min = math.Min(L.A.X, L.B.X)
		x_max = math.Max(L.A.X, L.B.X)
		y_min = math.Min(L.A.Y, L.B.Y)
		y_max = math.Max(L.A.Y, L.B.Y)

		return x_min < (P.X+EPSILON) && x_max > (P.X-EPSILON) && y_min < (P.Y+EPSILON) && y_max > (P.Y-EPSILON)
	}
	return false
}

/*
Determines whether or not lines AB (this line) and CD has intersection.

Returns true if they intersect, false otherwise.
*/
func (AB *Line) HasIntersection(CD *Line) bool {
	var m1, m2 float64

	m1 = (AB.A.Y - AB.B.Y) * (CD.A.X - CD.B.X)
	m2 = (CD.A.Y - CD.B.Y) * (AB.A.X - AB.B.X)

	return math.Abs(m1-m2) >= EPSILON
}

/*
Determines wheter segment AB (this) and PQ (other line) have an intersection, but doesn't overlap.
* Maybe mix this and does_segments overlap ? And call it somelike "PointsIntersecting", 0 if doesnt overlap nor intersect, 1 if insersect, 2 if overlap

TODO return intersection as well.
*/
func (AB *Line) DoesSegmentsIntersect(PQ *Line) bool {
	var o1, o2, o3, o4 int8
	var has_intersection bool
	var p_intersection *Point
	o1 = GetOrientation(AB.A, AB.B, PQ.A)
	o2 = GetOrientation(AB.A, AB.B, PQ.B)
	o3 = GetOrientation(PQ.A, PQ.B, AB.A)
	o4 = GetOrientation(PQ.A, PQ.B, AB.B)

	if o1+o2+o3+o4 == 0 { // Segments AB and PQ overlap.
		return false
	}

	has_intersection = AB.HasIntersection(PQ)
	if !has_intersection {
		return false
	}

	p_intersection = AB.IntersectionPointOnALine(PQ)
	return AB.PointOnSegment(p_intersection) && PQ.PointOnSegment(p_intersection)
}

/*
	Determines wheter segment AB (this) and PQ (other segment) overlaps.
*/

func (AB *Line) DoesSegmentsOverlap(PQ *Line) bool {
	var o1, o2, o3, o4 int8
	o1 = GetOrientation(AB.A, AB.B, PQ.A)
	o2 = GetOrientation(AB.A, AB.B, PQ.B)
	o3 = GetOrientation(PQ.A, PQ.B, AB.A)
	o4 = GetOrientation(PQ.A, PQ.B, AB.B)

	if o1+o2+o3+o4 == 0 && (PQ.PointOnSegment(AB.A) || PQ.PointOnSegment(AB.B) || AB.PointOnSegment(PQ.A) || AB.PointOnSegment(PQ.B)) {
		return true
	}
	return false
}

/*
Calculates the angle between segment AP and segment PB.
*/
func (AP *Line) AngleBetweenSegments(PB *Line) float64 {
	var A_seen_by_P, B_seen_by_P *Point
	var norm_AB, norm_PQ, dot_product, result float64
	A_seen_by_P = NewPoint(AP.A.X-AP.B.X, AP.A.Y-AP.B.Y)
	B_seen_by_P = NewPoint(PB.B.X-PB.A.X, PB.B.Y-PB.A.Y)
	norm_AB = VectorNorm(AP.A, AP.B)
	norm_PQ = VectorNorm(PB.B, PB.A)
	dot_product = A_seen_by_P.DotProduct(B_seen_by_P)

	result = dot_product / (norm_AB * norm_PQ)
	return math.Acos(result)
}
