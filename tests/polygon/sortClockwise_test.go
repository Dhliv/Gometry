package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/src/models"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestSortClockwise(t *testing.T) {
	polygons := mock_polygons.AllPolygons
	result_polygons := mock_polygons.AllClockwisePolygons

	for i, polygon := range polygons {
		polygon.SortClockwise()
		var result_polygon *models.Polygon = result_polygons[i]

		if !(polygon.Equal(result_polygon, true)) {
			t.Errorf("Error in clockwise sorting in polygon #%v! Expected %v. Got instead %v.", i, result_polygon.ToString(), polygon.ToString())
		}
	}
}
