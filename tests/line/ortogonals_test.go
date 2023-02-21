package line_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_lines"
)

/*
Criteria of test:
*/
func TestOrthogonal(t *testing.T) {
	lines := mock_lines.AllLine
	origins := mock_lines.Origins_test2
	orthogonal := mock_lines.OrthogonalIntersection_test2
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		origin := origins[i]
		expected := orthogonal[i]
		result := line.OrthogonalLinePoint(origin)
		if !expected.Equal(result) {
			t.Errorf("Error in point calculation #%v! Expected %v. Got instead %v.", i, expected, result)
		}
	}

}
