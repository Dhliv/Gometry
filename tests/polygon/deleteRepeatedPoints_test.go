package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/src/models"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestDeleteRepeatedPoints(t *testing.T) {
	polygons := mock_polygons.AllRepeatedPointsPolygons
	result_polygons := mock_polygons.AllNotRepeatedPointsPolygons

	for i, polygon := range polygons {
		polygon.DeleteRepeatedPoints()
		var result_polygon *models.Polygon = result_polygons[i]

		if !(polygon.Equal(result_polygon, true)) {
			t.Errorf("Error in deleting repeated points in polygon #%v! Expected %v. Got instead %v.", i, result_polygon.ToString(), polygon.ToString())
		}
	}
}
