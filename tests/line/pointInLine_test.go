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
func TestSamePointsOfTheLine(t *testing.T) {
	lines := mock_lines.AllLine

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if !line.PointInLine(line.A) || !line.PointInLine(line.B) {
			t.Errorf("Error on test #%v! The points belongs to the line because are literally the atributes A and B.", i)
		}
	}

}

/*
This is going to test with 2 different points that the struct Line has. One is on the line, the another one not.
*/
func TestDifferentPointsOfTheLine(t *testing.T) {
	lines := mock_lines.AllLine
	pointsInLine := mock_lines.PointsInLine
	pointsOutsideLine := mock_lines.PointsOutsideLine
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		inLine := pointsInLine[i]
		outLine := pointsOutsideLine[i]
		if !line.PointInLine(inLine) {
			t.Errorf("Error on test #%v.1! The point %v belongs to the line %v.", i, inLine, line)
		}
		if line.PointInLine(outLine) {
			t.Errorf("Error on test #%v.2! The point %v doesn't belong to the line %v.", i, outLine, line)
		}
	}
}
