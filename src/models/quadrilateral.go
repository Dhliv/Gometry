package models

import . "github.com/Dhliv/Gometry/src/utils"

/*
This class contains 4 points, being (x, y). This beforementioned points represents
the Camera Footprint of drone's camera.
*/
type Quadrilateral struct {
	A, B, C, D *Point
}

/*
Initialize projected quadrilateral on the form:

..A----------B..

.....C----D.....
*/
func NewQuadrilateral(A, B, C, D *Point) *Quadrilateral {
	return &Quadrilateral{A: A, B: B, C: C, D: D}
}

func (Q *Quadrilateral) Rotate(rotationAngle float64, isYaw bool) {
	var beta = rotationAngle
	if isYaw {
		/*
			This is because we want to rotate the vectors counter-clockwise so the y-axis
			points directly to the north, but the rotation_angle is measured clockwise, so
			we do the equivalence of this angle by doing the rotation in the opposite sense.
		*/
		beta = COMPLETE_CIRCUNFERENCE - rotationAngle*ONE_GRADE_IN_RADIANS
	}

	Q.A.Rotate(beta)
	Q.B.Rotate(beta)
	Q.C.Rotate(beta)
	Q.D.Rotate(beta)
}

// Translate the quadrilateral by 'originPoint'. This is vectorial sum.
func (Q *Quadrilateral) Translate(originPoint *Point) {
	Q.A.Translate(originPoint)
	Q.B.Translate(originPoint)
	Q.C.Translate(originPoint)
	Q.D.Translate(originPoint)
}

// Represent the quadrilateral as Polygon
func (Q *Quadrilateral) ToWeilerAthertonRepresentation() *Polygon {
	return NewPolygon(Q.A, Q.B, Q.D, Q.C)
}
