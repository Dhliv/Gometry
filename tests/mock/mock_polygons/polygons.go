package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	Polygon     *models.Polygon = models.NewPolygon(A, B, C, D)
	Polygon1                    = models.NewPolygon(S1, S7, S6, S5, S4, S3, S2)
	Polygon2                    = models.NewPolygon(S1, S7, S6, S3, S2)
	Polygon3                    = models.NewPolygon(S8, S1, S2, S6)
	Polygon4                    = models.NewPolygon(S9, S2, S3, S6, S8)
	Polygon5                    = models.NewPolygon(S3, S5, S11)
	Polygon6                    = models.NewPolygon(S10, S1, S2, S6, S8)
	Polygon7                    = models.NewPolygon(S10, S6, S8)
	Polygon8                    = models.NewPolygon(S11, S12, S6, S8)
	Polygon9                    = models.NewPolygon(S12, S13, S6, S14, S15, S8, S3)
	Polygon10                   = models.NewPolygon(S12, S13, S14, B, S5, S3)
	Polygon11                   = models.NewPolygon(S14, S13, S12, S2, S6, B)
	Polygon12                   = models.NewPolygon(S2, S15, B, S14, S5, S4, S13, S12)
	Polygon13                   = models.NewPolygon(S17, S16, S15, S5, S8)
	Polygon14                   = models.NewPolygon(S6, S8, D)
	Polygon15                   = models.NewPolygon(S15, S8, D, A)
	Polygon16                   = models.NewPolygon(S18, S9, S13, B, D)
	Polygon17                   = models.NewPolygon(S1, S6, S13, S14, S15, S8, S11, S3)
	Polygon18                   = models.NewPolygon(S10, S19, S15, S5, S8)
	Polygon19                   = models.NewPolygon(S20, S11, S5, S14)
	Polygon20                   = models.NewPolygon(S9, S7, S12)
	Polygon21                   = models.NewPolygon(S24, S23, S10, S21, S22)
	Polygon22                   = models.NewPolygon(S18, S8, S17)
	Polygon23                   = models.NewPolygon(S17, S20, S23)
	Polygon24                   = models.NewPolygon(S20, S6, S12, S3, S13, S14, S5, B)
	Polygon_ccw                 = models.NewPolygon(S2, S3, S4, S5, S6, S7, S1)

	AllPolygons []*models.Polygon = []*models.Polygon{Polygon, Polygon1, Polygon2, Polygon3, Polygon4, Polygon5, Polygon6, Polygon7, Polygon8, Polygon9, Polygon10, Polygon11, Polygon12, Polygon13, Polygon14, Polygon15, Polygon16, Polygon17, Polygon18, Polygon19, Polygon20, Polygon21, Polygon22, Polygon23, Polygon24, Polygon_ccw}
)
