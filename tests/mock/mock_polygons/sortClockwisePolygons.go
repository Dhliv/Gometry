package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	ClockwisePolygon     *models.Polygon = models.NewPolygon(A, B, D, C)
	ClockwisePolygon1                    = models.NewPolygon(S1, S7, S6, S5, S4, S3, S2)
	ClockwisePolygon2                    = models.NewPolygon(S1, S7, S6, S3, S2)
	ClockwisePolygon3                    = models.NewPolygon(S8, S1, S2, S6)
	ClockwisePolygon4                    = models.NewPolygon(S9, S2, S3, S6, S8)
	ClockwisePolygon5                    = models.NewPolygon(S3, S5, S11)
	ClockwisePolygon6                    = models.NewPolygon(S10, S1, S2, S6, S8)
	ClockwisePolygon7                    = models.NewPolygon(S10, S6, S8)
	ClockwisePolygon8                    = models.NewPolygon(S11, S12, S6, S8)
	ClockwisePolygon9                    = models.NewPolygon(S12, S13, S6, S14, S15, S8, S3)
	ClockwisePolygon10                   = models.NewPolygon(S12, S13, S14, B, S5, S3)
	ClockwisePolygon11                   = models.NewPolygon(S14, S13, S12, S2, S6, B)
	ClockwisePolygon12                   = models.NewPolygon(S2, S15, B, S14, S5, S4, S13, S12)
	ClockwisePolygon13                   = models.NewPolygon(S17, S16, S15, S5, S8)
	ClockwisePolygon14                   = models.NewPolygon(S6, S8, D)
	ClockwisePolygon15                   = models.NewPolygon(S15, S8, D, A)
	ClockwisePolygon16                   = models.NewPolygon(S18, S9, S13, B, D)
	ClockwisePolygon17                   = models.NewPolygon(S1, S6, S13, S14, S15, S8, S11, S3)
	ClockwisePolygon18                   = models.NewPolygon(S10, S19, S15, S5, S8)
	ClockwisePolygon19                   = models.NewPolygon(S20, S11, S5, S14)
	ClockwisePolygon20                   = models.NewPolygon(S9, S7, S12)
	ClockwisePolygon21                   = models.NewPolygon(S24, S23, S10, S21, S22)
	ClockwisePolygon22                   = models.NewPolygon(S18, S8, S17)
	ClockwisePolygon23                   = models.NewPolygon(S23, S20, S17)                     // Changed
	ClockwisePolygon24                   = models.NewPolygon(S6, S20, B, S5, S14, S13, S3, S12) // Changed
	ClockwisePolygon_ccw                 = models.NewPolygon(S6, S5, S4, S3, S2, S1, S7)        // Changed

	AllClockwisePolygons []*models.Polygon = []*models.Polygon{ClockwisePolygon, ClockwisePolygon1, ClockwisePolygon2, ClockwisePolygon3, ClockwisePolygon4, ClockwisePolygon5, ClockwisePolygon6, ClockwisePolygon7, ClockwisePolygon8, ClockwisePolygon9, ClockwisePolygon10, ClockwisePolygon11, ClockwisePolygon12, ClockwisePolygon13, ClockwisePolygon14, ClockwisePolygon15, ClockwisePolygon16, ClockwisePolygon17, ClockwisePolygon18, ClockwisePolygon19, ClockwisePolygon20, ClockwisePolygon21, ClockwisePolygon22, ClockwisePolygon23, ClockwisePolygon24, ClockwisePolygon_ccw}
)
