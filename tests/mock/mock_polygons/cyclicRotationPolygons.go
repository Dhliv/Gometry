package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

type cyclicPolygon struct {
	Polygon *models.Polygon
	Times   int
}

var (
	AllCyclicPolygons []*cyclicPolygon = []*cyclicPolygon{{Polygon: Polygon, Times: 0}, {Polygon: Polygon1, Times: 1}, {Polygon: Polygon2, Times: 2}, {Polygon: Polygon3, Times: 3}, {Polygon: Polygon4, Times: 4}, {Polygon: Polygon5, Times: 5}, {Polygon: Polygon6, Times: 6}, {Polygon: Polygon7, Times: 7}, {Polygon: Polygon8, Times: 8}, {Polygon: Polygon9, Times: 9}, {Polygon: Polygon10, Times: 10}, {Polygon: Polygon11, Times: 11}, {Polygon: Polygon12, Times: 12}, {Polygon: Polygon13, Times: 13}, {Polygon: Polygon14, Times: 14}, {Polygon: Polygon15, Times: 15}, {Polygon: Polygon16, Times: 16}, {Polygon: Polygon17, Times: 17}, {Polygon: Polygon18, Times: 18}, {Polygon: Polygon19, Times: 19}, {Polygon: Polygon20, Times: 20}, {Polygon: Polygon21, Times: 21}, {Polygon: Polygon22, Times: 22}, {Polygon: Polygon23, Times: 23}, {Polygon: Polygon24, Times: 24}, {Polygon: Polygon_ccw, Times: 25}}
)

var (
	CyclicPolygon     *models.Polygon = models.NewPolygon(A, B, D, C)
	CyclicPolygon1                    = models.NewPolygon(S7, S6, S5, S4, S3, S2, S1)
	CyclicPolygon2                    = models.NewPolygon(S6, S3, S2, S1, S7)
	CyclicPolygon3                    = models.NewPolygon(S6, S8, S1, S2)
	CyclicPolygon4                    = models.NewPolygon(S8, S9, S2, S3, S6)
	CyclicPolygon5                    = models.NewPolygon(S11, S3, S5)
	CyclicPolygon6                    = models.NewPolygon(S1, S2, S6, S8, S10)
	CyclicPolygon7                    = models.NewPolygon(S6, S8, S10)
	CyclicPolygon8                    = models.NewPolygon(S11, S12, S6, S8)
	CyclicPolygon9                    = models.NewPolygon(S6, S14, S15, S8, S3, S12, S13)
	CyclicPolygon10                   = models.NewPolygon(S5, S3, S12, S13, S14, B)
	CyclicPolygon11                   = models.NewPolygon(B, S14, S13, S12, S2, S6)
	CyclicPolygon12                   = models.NewPolygon(S5, S4, S13, S12, S2, S15, B, S14)
	CyclicPolygon13                   = models.NewPolygon(S5, S8, S17, S16, S15)
	CyclicPolygon14                   = models.NewPolygon(D, S6, S8)
	CyclicPolygon15                   = models.NewPolygon(A, S15, S8, D)
	CyclicPolygon16                   = models.NewPolygon(S9, S13, B, D, S18)
	CyclicPolygon17                   = models.NewPolygon(S6, S13, S14, S15, S8, S11, S3, S1)
	CyclicPolygon18                   = models.NewPolygon(S5, S8, S10, S19, S15)
	CyclicPolygon19                   = models.NewPolygon(S14, S20, S11, S5)
	CyclicPolygon20                   = models.NewPolygon(S12, S9, S7)
	CyclicPolygon21                   = models.NewPolygon(S23, S10, S21, S22, S24)
	CyclicPolygon22                   = models.NewPolygon(S8, S17, S18)
	CyclicPolygon23                   = models.NewPolygon(S23, S17, S20)
	CyclicPolygon24                   = models.NewPolygon(S20, S6, S12, S3, S13, S14, S5, B)
	CyclicPolygon_ccw                 = models.NewPolygon(S6, S7, S1, S2, S3, S4, S5)

	AllCyclicPolygonsResults []*models.Polygon = []*models.Polygon{CyclicPolygon, CyclicPolygon1, CyclicPolygon2, CyclicPolygon3, CyclicPolygon4, CyclicPolygon5, CyclicPolygon6, CyclicPolygon7, CyclicPolygon8, CyclicPolygon9, CyclicPolygon10, CyclicPolygon11, CyclicPolygon12, CyclicPolygon13, CyclicPolygon14, CyclicPolygon15, CyclicPolygon16, CyclicPolygon17, CyclicPolygon18, CyclicPolygon19, CyclicPolygon20, CyclicPolygon21, CyclicPolygon22, CyclicPolygon23, CyclicPolygon24, CyclicPolygon_ccw}
)
