package models

import "fmt"

/*
Contains information about point 'P' like if it's on border of polygon, its index on another polygon in case is border,
if it belongs to a segment of border points and in case it belongs to some segment of border points it tells wheter
it's the end of said segment.
*/
type BorderPoint struct {
	P                          *Point
	Index                      int
	IsBorder, IsSegment, IsEnd bool
}

/*
Default constructor for BorderPoint. This constructor takes point 'P' as a simple intersection point that was calculated.
*/
func NewDefaultBorderPoint(P *Point, idx int) *BorderPoint {
	return &BorderPoint{
		P:        P,
		Index:    idx,
		IsBorder: false, IsSegment: false, IsEnd: false,
	}
}

/*
Constructor of BorderPoint. Relies on:
  - P: current point that lies directly on polygon perimeter and was part of another polygon.
  - idx: index of current point in the polygon list in which it was checked.
  - isBorder: tells wheter current point lies on the border (perimeter) from some polygon.
  - isSegment: tells wheter current point is part of segment of border points.
  - isEnd: tells wheter current point is the end of segment of border poinst (in case that is segment).
*/
func NewBorderPoint(P *Point, idx int, isBorder, isSegment, isEnd bool) *BorderPoint {
	return &BorderPoint{
		P:        P,
		Index:    idx,
		IsBorder: isBorder, IsSegment: isSegment, IsEnd: isEnd,
	}
}

func (bp *BorderPoint) String() string {
	return fmt.Sprintf("P: %v, Index: %v, IsBorder: %v, IsSegment: %v, IsEnd: %v", bp.P.ToString(), bp.Index, bp.IsBorder, bp.IsSegment, bp.IsEnd)
}
