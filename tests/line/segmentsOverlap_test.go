package line_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:

Comparision with the data.
All the segments overlap.
*/
func TestTwoSegmentOverlap(t *testing.T) {
	lines := mock_lines.PairOfOverlappingSegments_test8
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		result := L.DoesSegmentsOverlap(&R)
		if !result {
			t.Errorf("Error in test #%v.1! The segments intersect.", i)
		}
	}

}
