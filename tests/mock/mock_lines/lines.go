package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	VerticalLine1   models.Line   = models.Line{A: A, B: D}
	VerticalLine2                 = models.Line{A: D, B: S3}
	VerticalLine3                 = models.Line{A: S4, B: S5}
	HorizontalLine1               = models.Line{A: B, B: C}
	HorizontalLine2               = models.Line{A: C, B: B}
	HorizontalLine3               = models.Line{A: S14, B: S9}
	Line1                         = models.Line{A: A, B: B}
	Line2                         = models.Line{A: C, B: D}
	Line3                         = models.Line{A: S1, B: S2}
	Line4                         = models.Line{A: S3, B: S4}
	AllLine         []models.Line = []models.Line{
		VerticalLine1, VerticalLine2, VerticalLine3,
		HorizontalLine1, HorizontalLine2, HorizontalLine3,
		Line1, Line2, Line3, Line4,
	}
)
