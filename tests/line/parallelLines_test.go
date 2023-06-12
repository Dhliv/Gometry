package line_test

import (
	"math"
	"testing"

	. "github.com/Dhliv/Gometry/src/utils"
	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:

# Slope of Line R is equal to Line L

Point P belongs to the line R
*/
func TestParallel(t *testing.T) {
	lines := mock_lines.Lines_test4
	points := mock_lines.Points_test4
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		point := points[i]
		result := line.ParallelLineOfALineAndAPoint(point)
		if math.Abs(line.Slope()-result.Slope()) > EPSILON {
			t.Errorf("Error on test #%v! The slopes of the lines has to be the same. L has slope %v and R has slope %v", i, line.Slope(), result.Slope())
		}
		// if result.PointInLine(point) {
		// 	t.Errorf("Error on test #%v! The point %v has to be in R.", i, point)
		// }
	}
}
