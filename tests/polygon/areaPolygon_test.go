package polygon_test

import (
	"math"
	"testing"

	. "github.com/Dhliv/Gometry/src/utils"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

func TestAreaPolygon(t *testing.T) {
	polygons := mock_polygons.AllPolygons
	result_areas := mock_polygons.AllAreaPolygons
	var area float64

	for i := 0; i < len(polygons); i++ {
		polygon := polygons[i]
		area = polygon.Area(true)

		if math.Abs(area-result_areas[i]) > EPSILON {
			t.Errorf("Error in area calculation of polygon #%v! Expected %v. Got instead %v.", i, result_areas[i], area)
		}
	}
}
