package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	PointsInSegment []*models.Point = []*models.Point{
		S3, S11, S4,
		S11, S11, S9,
		S14, S22, models.NewPoint(1.866863816, 9.066779323), models.NewPoint(4.232700635, 6.027250529),
	}
	PointsOutsideSegment []*models.Point = []*models.Point{
		S2, S13, S11,
		S22, D, S11,
		S4, S18, S16, S13,
	}
)
