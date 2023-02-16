package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	Origins_test2                []*models.Point = []*models.Point{B, B, S1, D, A, S13, S4, S3, S9, B}
	OrthogonalIntersection_test2 []*models.Point = []*models.Point{
		models.NewPoint(3, 3),
		models.NewPoint(3, 3),
		models.NewPoint(6, 8),
		models.NewPoint(3, 3),
		models.NewPoint(3, 3),
		models.NewPoint(5, 5),
		models.NewPoint(5.25, 6.75),
		models.NewPoint(1, 3),
		models.NewPoint(-1.29230769, 7.26153846),
		models.NewPoint(5.557377049, 7.13114754),
	}
)
