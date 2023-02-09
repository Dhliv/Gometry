package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/src/models"
	"github.com/Dhliv/Gometry/tests/mock/mock_points"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestPointInPolygonPerimeter(t *testing.T) {
	polygons := mock_polygons.AllPolygons
	resultPointsInPolygon := mock_polygons.AllPointsOnBorderPolygons
	points := mock_points.AllPoints
	var A *models.Point

	for i, polygon := range polygons {

		for j := 0; j < len(points); j++ {
			A = points[j]
			res := polygon.PointInPolygonPerimeter(A)
			if res != resultPointsInPolygon[i][j] {
				t.Errorf("Error in point on polygon perimeter at polygon #%v and point #%v! Expected %v. Got instead %v.", i, j, resultPointsInPolygon[i][j], res)
			}
		}
	}
}
