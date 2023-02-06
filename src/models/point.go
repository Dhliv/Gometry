package models

import (
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

// Point instance ind 2d coordinate system.
type Point struct {
	X float64
	Y float64
}

/*
Point costructor
*/
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// Rotates the actual point by 'beta' degrees in counter-clocwise.
func (P *Point) Rotate(beta float64) {
	x := math.Cos(beta)*P.X - math.Sin(beta)*P.Y
	y := math.Sin(beta)*P.X + math.Cos(beta)*P.Y
	P.X = x
	P.Y = y
}

// Translate the origin point by 'translation_point'.
func (P *Point) Translate(translationPoint *Point) {
	P.X += translationPoint.Y
	P.Y += translationPoint.Y
}

// Recieves a Point and return true if both are equal, false otherwise
func (Pe *Point) Equal(P *Point) bool {
	if math.Abs(Pe.X-P.X) > EPSILON || math.Abs(Pe.Y-P.Y) > EPSILON {
		return false
	}
	return true
}

func (Pe *Point) DotProduct(P *Point) float64 {
	return Pe.X*P.X + Pe.Y*P.Y
}

// Calculates the norm of 'vector' to 'origin'.
func VectorNorm(vector, origin Point) float64 {
	x := math.Pow((vector.X - origin.X), 2)
	y := math.Pow((vector.Y - origin.Y), 2)

	return math.Sqrt(x + y)
}

/*
Determines the orientation of the points in the order A->B->C, returning one of the
following results:

  - 0 : the points were collinear.
  - 1 : the points were clockwise.
  - 2 : the points were counterclockwise.
*/
func GetOrientation(A, B, C *Point) int8 {
	var val float64 = (B.Y-C.Y)*(B.X-A.X) - (A.Y-B.Y)*(C.X-B.X)

	if math.Abs(val) < EPSILON {
		return COLLINEAR_ORIENTATION
	}

	if val > float64(0.0) {
		return CLOCKWISE_ORIENTATION
	}

	return COUNTERCLOCKWISE_ORIENTATION
}
