package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	RepeatedPointsPolygon  *models.Polygon = models.NewPolygon(A, B, D, C)
	RepeatedPointsPolygon1                 = models.NewPolygon(S1, S12, A, S12, S15, S15, S15, S17, S22, S22, S1)
	RepeatedPointsPolygon2                 = models.NewPolygon(S17, S17, D, C, S12, A, S12, S13, S24, S24, S14, B, B, S17, S17)

	NotRepeatedPointsPolygon  = models.NewPolygon(A, B, D, C)
	NotRepeatedPointsPolygon1 = models.NewPolygon(S1, S12, S15, S17, S22)
	NotRepeatedPointsPolygon2 = models.NewPolygon(S17, D, C, S12, S13, S24, S14, B)

	AllRepeatedPointsPolygons    []*models.Polygon = []*models.Polygon{RepeatedPointsPolygon, RepeatedPointsPolygon1, RepeatedPointsPolygon2}
	AllNotRepeatedPointsPolygons []*models.Polygon = []*models.Polygon{NotRepeatedPointsPolygon, NotRepeatedPointsPolygon1, NotRepeatedPointsPolygon2}
)
