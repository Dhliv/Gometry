package line_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:
*/

/*
This is going to test with the same points that the struct Line has.
*/
func TestSamePointsOfTheSegment(t *testing.T) {
	lines := mock_lines.AllLine

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if !line.PointOnSegment(line.A) || !line.PointOnSegment(line.B) {
			t.Errorf("Error on test #%v! The points belongs to the line because are literally the atributes A and B.", i)
		}
	}

}

/*
This is going to test with 2 different points that the struct Line has. One is on the segment, the another one not.
*/
func TestDifferentPointsOfTheSegment(t *testing.T) {
	lines := mock_lines.AllLine
	pointsInLine := mock_lines.PointsInSegment
	pointsOutsideLine := mock_lines.PointsOutsideSegment
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		inLine := pointsInLine[i]
		outLine := pointsOutsideLine[i]
		if !line.PointOnSegment(inLine) {
			t.Errorf("Error on test #%v.1! The point %v belongs to the line %v.", i, inLine, line)
		}
		if line.PointOnSegment(outLine) {
			t.Errorf("Error on test #%v.2! The point %v doesn't belong to the line %v.", i, outLine, line)
		}
	}
}
