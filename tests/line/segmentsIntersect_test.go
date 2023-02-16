package line_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:

Comparision with the data.
All the segments intersect.
*/
func TestTwoSegmentIntersect(t *testing.T) {
	lines := mock_lines.PairOfSegments_test7
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		result := L.DoesSegmentsIntersect(&R)
		if !result {
			t.Errorf("Error in test #%v.1! The segments intersect.", i)
		}
	}

}

/*
The test will has two lines of the form AB and BC, for this reason, the answer HAS TO BE 1.

Criteria of the test:
The point R is equal to the B point of the line L and the A point of the line R.
*/
func TestTwoSegmentsThatSharesOnePointIntersect(t *testing.T) {
	lines := mock_lines.LinesFormedBy3Points
	for i := 0; i < len(lines); i++ {
		L := lines[i][0]
		R := lines[i][1]
		result := L.DoesSegmentsIntersect(&R)
		if !result {
			t.Errorf("Error in test #%v.1! The segments intersect.", i)
		}

	}

}
