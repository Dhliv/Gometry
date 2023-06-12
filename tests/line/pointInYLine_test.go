package line_test

import (
	// "math"
	"testing"
	// . "github.com/Dhliv/Gometry/src/utils"
	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:
Compare the return of the method GetPointInsideLineByYCoordinate with the data extracted.
*/
func TestPointInYLine(t *testing.T) {
	lines := mock_lines.Lines_test1
	points := mock_lines.Points_test1
	ys := mock_lines.Y_points_test1
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		point := points[i]
		y := ys[i]
		result := line.GetPointInsideLineByYCoordinate(y)
		if !point.Equal(result) {
			t.Errorf("Error in point calculation #%v! Expected %v. Got instead %v.", i, point, result)
		}
	}
}
