package mock_lines

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	PairOfLines_test5 [][]models.Line = [][]models.Line{
		{HorizontalLine1, VerticalLine1},
		{HorizontalLine2, VerticalLine2},
		{HorizontalLine3, VerticalLine3},
		{Line1, Line3},
		{Line3, Line4},
		{Line1, Line4},
	}
	IntersectionPoints []*models.Point = []*models.Point{
		models.NewPoint(3, 3),
		models.NewPoint(3, 3),
		models.NewPoint(6, 5),
		models.NewPoint(2.54545454, 9.454545454),
		models.NewPoint(21, 20),
		models.NewPoint(5.181818181, 6.818181818),
	}
	LinesFormedBy3Points [][]models.Line = [][]models.Line{
		{models.Line{A: A, B: D}, models.Line{A: D, B: B}},
		{models.Line{A: C, B: D}, models.Line{A: D, B: B}},
		{models.Line{A: C, B: B}, models.Line{A: B, B: D}},
		{models.Line{A: S8, B: S20}, models.Line{A: S20, B: S14}},
		{models.Line{A: S3, B: S9}, models.Line{A: S9, B: S13}},
	}
)
