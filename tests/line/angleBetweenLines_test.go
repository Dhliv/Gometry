package line_test

import (
	"math"
	"testing"

	. "github.com/Dhliv/Gometry/src/utils"
	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

func TestAngleOfTwoLinesThatSharesOnePointIntersect(t *testing.T) {
	lines := mock_lines.LinesFormedBy3Points
	angles := mock_lines.AngleBetweemLinesFormedBy3Points
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		result := L.AngleBetweenSegments(&R)
		if math.Abs(result-angles[i]) > EPSILON {
			t.Errorf("Error in test #%v.1! The angle is %v, obtained %v", i, angles[i], result)
		}

	}

}
