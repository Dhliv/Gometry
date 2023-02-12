package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestPolygonInsidePolygon(t *testing.T) {
	polygons := mock_polygons.AllPolygons
	resultsPolygonInPolygon := mock_polygons.AllPolygonsContainPolygons

	for i, polygon := range polygons {

		for j, polygon2 := range polygons {
			res := polygon.InsidePolygon(&polygon2)
			if res != resultsPolygonInPolygon[i][j] {
				t.Errorf("Error in point on polygon perimeter at polygon #%v and point #%v! Expected %v. Got instead %v.", i, j, resultsPolygonInPolygon[i][j], res)
			}
		}
	}
}
