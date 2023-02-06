package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/src/models"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestDeleteCollinearSegments(t *testing.T) {
	polygons := mock_polygons.AllPolygons
	result_polygons := mock_polygons.AllNCPolygons

	for i, polygon := range polygons {
		polygon.DeleteCollinearSegments()
		var result_polygon *models.Polygon = result_polygons[i]

		if !(polygon.Equal(result_polygon, true)) {
			t.Errorf("Error in deleting collinear segments in polygon #%v! Expected %v. Got instead %v.", i, result_polygon.ToString(), polygon.ToString())
		}
	}
}
