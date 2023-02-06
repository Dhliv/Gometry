package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

// Polygons without collinear segments
var (
	NCPolygon     *models.Polygon = models.NewPolygon(A, B, C, D)                         // Same
	NCPolygon1                    = models.NewPolygon(S1, S7, S6, S5, S4, S3, S2)         // Same
	NCPolygon2                    = models.NewPolygon(S1, S7, S6, S3, S2)                 // Same
	NCPolygon3                    = models.NewPolygon(S8, S1, S2, S6)                     // Same
	NCPolygon4                    = models.NewPolygon(S9, S2, S3, S6, S8)                 // Same
	NCPolygon5                    = models.NewPolygon(S3, S5, S11)                        // Same
	NCPolygon6                    = models.NewPolygon(S10, S1, S2, S6, S8)                // Same
	NCPolygon7                    = models.NewPolygon(S10, S6, S8)                        // Same
	NCPolygon8                    = models.NewPolygon(S11, S12, S6, S8)                   // Same
	NCPolygon9                    = models.NewPolygon(S12, S13, S6, S14, S15, S8, S3)     // Same
	NCPolygon10                   = models.NewPolygon(S12, B, S5, S3)                     // Not Same
	NCPolygon11                   = models.NewPolygon(S12, S2, S6, B)                     // Not Same
	NCPolygon12                   = models.NewPolygon(S2, S15, B, S14, S5, S4, S13, S12)  // Same
	NCPolygon13                   = models.NewPolygon(S17, S16, S15, S5, S8)              // Same
	NCPolygon14                   = models.NewPolygon(S6, S8, D)                          // Same
	NCPolygon15                   = models.NewPolygon(S15, S8, D, A)                      // Same
	NCPolygon16                   = models.NewPolygon(S18, S9, S13, B, D)                 // Same
	NCPolygon17                   = models.NewPolygon(S1, S6, S13, S14, S15, S8, S11, S3) // Same
	NCPolygon18                   = models.NewPolygon(S10, S19, S15, S5, S8)              // Same
	NCPolygon19                   = models.NewPolygon(S20, S11, S5, S14)                  // Same
	NCPolygon20                   = models.NewPolygon(S9, S7, S12)                        // Same
	NCPolygon21                   = models.NewPolygon(S24, S23, S10, S21, S22)            // Same
	NCPolygon22                   = models.NewPolygon(S18, S8, S17)                       // Same
	NCPolygon23                   = models.NewPolygon(S17, S20, S23)                      // Same
	NCPolygon24                   = models.NewPolygon(S20, S6, S12, S3, S13, S14, S5, B)  // Same
	NCPolygon_ccw                 = models.NewPolygon(S2, S3, S4, S5, S6, S7, S1)         // Same

	AllNCPolygons []*models.Polygon = []*models.Polygon{NCPolygon, NCPolygon1, NCPolygon2, NCPolygon3, NCPolygon4, NCPolygon5, NCPolygon6, NCPolygon7, NCPolygon8, NCPolygon9, NCPolygon10, NCPolygon11, NCPolygon12, NCPolygon13, NCPolygon14, NCPolygon15, NCPolygon16, NCPolygon17, NCPolygon18, NCPolygon19, NCPolygon20, NCPolygon21, NCPolygon22, NCPolygon23, NCPolygon24, NCPolygon_ccw}
)
