package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
)

var (
	Y_points_test1 []float64 = []float64{
		30, 40, 60,
		1000000, 0.0001, 30.8723, 10.23245}
	Lines_test1 []models.Line = []models.Line{
		VerticalLine1, VerticalLine2, VerticalLine3,
		Line1, Line2, Line3, Line4,
	}
	Points_test1 []*models.Point = []*models.Point{
		models.NewPoint(3, 30),
		models.NewPoint(3, 40),
		models.NewPoint(6, 60),
		models.NewPoint(-999988, 1000000),
		models.NewPoint(3.9999, 0.0001),
		models.NewPoint(40.026525, 30.8723),
		models.NewPoint(9.27894, 10.23245),
	}
)
