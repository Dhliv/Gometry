package polygon_test

import (
	"testing"

	"github.com/Dhliv/Gometry/src/logic"
	"github.com/Dhliv/Gometry/src/models"
	"github.com/Dhliv/Gometry/tests/mock/mock_polygons"
)

var (
	polygons        []models.Polygon    = mock_polygons.AllPolygons
	result_polygons [][]*models.Polygon = mock_polygons.AllIntersectionPolygons
)

func showError(t *testing.T, i int, result []*models.Polygon) {
	var p1, p2 string = "[", "["

	for j, polygon := range result_polygons[i] {
		if j > 0 {
			p1 += ", "
		}
		p1 += polygon.ToString()
	}
	p1 += "]"

	for j, polygon := range result {
		if j > 0 {
			p2 += ", "
		}
		p2 += polygon.ToString()
	}
	p2 += "]"

	t.Errorf("Error in intersection polygons in polygon #%v! Expected %v. Got instead %v.", i, p1, p2)
}

func TestIntersectionPolygons(t *testing.T) {
	cameraFootprint := polygons[0]
	var found bool

	for i, polygon := range polygons {
		wa := logic.NewWeilerAtherton(polygon.Copy(), cameraFootprint.Copy())
		intersectionPolygons := wa.GetIntersectionPolygons()

		if len(intersectionPolygons) != len(result_polygons[i]) {
			showError(t, i, intersectionPolygons)
			continue
		}

		for _, polygon := range intersectionPolygons {
			found = false

			for _, polygon2 := range result_polygons[i] {
				if polygon.Equal(polygon2, true) {
					found = true
					break
				}
			}

			if !found {
				showError(t, i, intersectionPolygons)
				break
			}
		}
	}
}
