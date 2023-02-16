package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	PointsInLine []*models.Point = []*models.Point{
		S16, S17, S23, S11, S11, S21, S13, S22, models.NewPoint(5.8779601, 11.358834),
		models.NewPoint(-1.229508, 1.47540983),
	}
	PointsOutsideLine []*models.Point = []*models.Point{
		S2, S13, S11, S22, D, S11, S4, S18, S16, S13,
	}
)
