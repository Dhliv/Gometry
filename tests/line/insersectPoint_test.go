package line_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:

Comparision with the data.
The intersection point, named C, is in the line L and R.
*/
func TestIntersectPointOfTwoLines(t *testing.T) {
	lines := mock_lines.PairOfLines_test5
	points := mock_lines.IntersectionPoints
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		expected := points[i]
		result := L.IntersectionPointOnALine(&R)
		if !expected.Equal(result) {
			t.Errorf("Error in point calculation #%v.1! Expected %v. Got instead %v.", i, expected, result)
		}

		// if !L.PointInLine(result) {
		// 	t.Errorf("Error in test #%v.2! The point %v has to be in line L.", i, expected)
		// }
		// if !R.PointInLine(result) {
		// 	t.Errorf("Error in test #%v.2! The point %v has to be in line R.", i, expected)
		// }
	}

}

/*
The test will has two lines of the form AB and BC, for this reason, the answer HAS TO BE B.

Criteria of the test:
The point R is equal to the B point of the line L and the A point of the line R.
*/
func TestIntersectPointOfTwoLinesThatSharesOnePoint(t *testing.T) {
	lines := mock_lines.LinesFormedBy3Points
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		result := L.IntersectionPointOnALine(&R)
		if !result.Equal(L.B) || !result.Equal(R.A) {
			t.Errorf("Error in test #%v.1! The point %v has to be the same of L.B = %v and R.A = %v.", i, result, L.B, R.A)
		}

		// if !L.PointInLine(result) {
		// 	t.Errorf("Error in test #%v.2! The point %v has to be in line L.", i, expected)
		// }
		// if !R.PointInLine(result) {
		// 	t.Errorf("Error in test #%v.2! The point %v has to be in line R.", i, expected)
		// }
	}

}
