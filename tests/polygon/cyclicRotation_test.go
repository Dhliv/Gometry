package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestCyclicRotation(t *testing.T) {
	polygons := mock_polygons.AllCyclicPolygons
	result_polygons := mock_polygons.AllCyclicPolygonsResults

	for i := 0; i < len(polygons); i++ {
		polygon := (polygons[i].Polygon)
		polygon.CyclicRotation(polygons[i].Times)

		if !polygon.Equal(result_polygons[i], false) {
			t.Errorf("Error in cyclic rotation of polygon #%v! Expected %v. Got instead %v.", i, result_polygons[i].ToString(), polygon.ToString())
		}
	}
}
