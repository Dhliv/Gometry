package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	Lines_test4 []models.Line = []models.Line{ //We are not using vertical lines because its slope is infitite (div by 0)
		HorizontalLine1, HorizontalLine2, HorizontalLine3,
		Line1, Line2, Line3, Line4,
	}
	Points_test4 []*models.Point = []*models.Point{
		S3, S5, S23,
		S4, S3, S9, S11,
	}
)
