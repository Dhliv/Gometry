package logic

import (
	"fmt"
	"sort"

	"github.com/Dhliv/Gometry/src/models"
)

type WeilerAtherton struct {
	ClippedPolygon, ClippingPolygon         *models.Polygon
	ClippedPolygonList, ClippingPolygonList []*models.PointEntering
	ClippingPointsInClippedPerimeter        []bool
	BorderIntersections                     [][]*models.BorderPoint
}

/*
Constructor of WeilerAtherton. Relies on Clipping and Clipped Polygons. List of these polygons are empty by default. Take note that ClippedPolygon should be convex.
*/
func NewWeilerAtherton(clipping, clipped *models.Polygon) *WeilerAtherton {
	return &WeilerAtherton{
		ClippedPolygon: clipped, ClippingPolygon: clipping,
		ClippedPolygonList: []*models.PointEntering{}, ClippingPolygonList: []*models.PointEntering{},
		ClippingPointsInClippedPerimeter: []bool{},
		BorderIntersections:              [][]*models.BorderPoint{},
	}
}

// Creates a list of size len(ClippingPolygon) that stands True if the i-th point in 'ClippingPolygon' is
// on the perimeter of 'ClippedPolygon'.
func (wa *WeilerAtherton) CheckPointsInPerimeter() {
	wa.ClippingPointsInClippedPerimeter = make([]bool, len(*wa.ClippingPolygon.Points))

	for i, P := range *wa.ClippingPolygon.Points {
		wa.ClippingPointsInClippedPerimeter[i] = wa.ClippedPolygon.PointInPolygonPerimeter(P)
	}
}

// Sorts every intersection point inside a segment by its vectorial norm.
func (wa *WeilerAtherton) SortClippingPolygonIntersections() {
	var A, P, Q *models.Point

	for i := 0; i < len(*wa.ClippingPolygon.Points); i++ {
		A = (*wa.ClippingPolygon.Points)[i]

		sort.Slice(wa.BorderIntersections[i], func(x, y int) bool {
			var a, b float64
			P, Q = wa.BorderIntersections[i][x].P, wa.BorderIntersections[i][y].P
			a, b = models.VectorNorm(P, A), models.VectorNorm(Q, A)
			return a < b
		})
	}
}

// Handles the overlapping of 2 segments and decides which points will be considered to add to BorderIntersections.
func (wa *WeilerAtherton) handleOverlappingSegments(A, B, P, Q *models.Point, i, j int) {
	var AB, PQ *models.Line = models.NewLine(A, B), models.NewLine(P, Q)

	if PQ.PointOnSegment(A) {
		wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(A, j, true, true, false))
		wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Q, j, true, true, false))
		return
	}

	if PQ.PointOnSegment(B) {
		if AB.PointOnSegment(P) {
			wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(P, j, true, true, false))
			wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(B, j, true, true, true))
			return
		}

		wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Q, j, true, true, false))
		wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(B, j, true, true, true))
		return
	}

	wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(P, j, true, true, false))
	wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Q, j, true, true, true))
}

func (wa *WeilerAtherton) handleIntersectingSegments(A, B, P, Q *models.Point, i, j int) {
	var n int = len(*wa.ClippingPolygon.Points)
	var AB, PQ *models.Line = models.NewLine(A, B), models.NewLine(P, Q)
	var Intersection *models.Point = AB.IntersectionPointOnALine(PQ)
	var previousIndex, nextIndex, nextNextIndex = (i - 1 + n) % n, (i + 1) % n, (i + 2) % n
	// var currentPoint, nextPoint bool

	// Point can't appear twice on list
	if P.Equal(Intersection) {
		return
	}

	// A or B are border points
	if Intersection.Equal(A) || Intersection.Equal(B) {
		if Intersection.Equal(A) && wa.ClippingPointsInClippedPerimeter[previousIndex] {
			// intersection is end of border points segments.
			wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Intersection, j, true, true, true))
			return
		}

		if Intersection.Equal(A) {
			// intersection is only one border point.
			wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Intersection, j, true, false, false))
			return
		}

		if wa.ClippingPointsInClippedPerimeter[nextNextIndex] {
			// intersection point is the begining of border points segment.
			wa.BorderIntersections[nextIndex] = append(wa.BorderIntersections[nextIndex], models.NewBorderPoint(Intersection, j, true, true, false))
		}

		return
	}

	// Checking the validity of intersection point
	if Intersection.Equal(Q) {
		wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewBorderPoint(Intersection, j, true, false, false))
		return
	}

	wa.BorderIntersections[i] = append(wa.BorderIntersections[i], models.NewDefaultBorderPoint(Intersection, j))
}

/*
Given 2 polygons 'ClippedPolygon' and 'ClippingPolygon' determines which segments from 'ClippingPolygon' intersects with segments of 'ClippedPolygon'. Creates an saves a list of borderPoints that contains information about the intersection between a pair of segments from 'ClippedPolygon' and 'ClippingPolygon'.
*/
func (wa *WeilerAtherton) GetSegmentIntersections() {
	var n int = len(*wa.ClippingPolygon.Points)
	var m int = len(*wa.ClippedPolygon.Points)
	var nextIndex int
	var A, B, P, Q *models.Point
	var L1, L2 *models.Line
	wa.BorderIntersections = make([][]*models.BorderPoint, n)

	for i := 0; i < n; i++ {
		nextIndex = (i + 1) % n

		if wa.ClippingPointsInClippedPerimeter[i] && wa.ClippingPointsInClippedPerimeter[nextIndex] {
			continue
		}

		A = (*wa.ClippingPolygon.Points)[i]
		B = (*wa.ClippingPolygon.Points)[nextIndex]
		L1 = models.NewLine(A, B)

		for j := 0; j < m; j++ {
			P = (*wa.ClippedPolygon.Points)[j]
			Q = (*wa.ClippedPolygon.Points)[(j+1)%m]
			L2 = models.NewLine(P, Q)

			if L1.DoesSegmentsOverlap(L2) {
				wa.handleOverlappingSegments(A, B, P, Q, i, j)
				continue
			}

			if L1.DoesSegmentsIntersect(L2) {
				fmt.Printf("Segments %v and %v have an intersection\n", i, j)
				wa.handleIntersectingSegments(A, B, P, Q, i, j)
				continue
			}
		}
	}
}

func (wa *WeilerAtherton) BuildClippingAndClippedPolygonLists() {
	var BP *models.BorderPoint
	var P, A, B *models.Point
	var PE *models.PointEntering
	var isIncomplete, isSingle, isEntering, lastPoint, nextPoint bool
	var clippedIntersections [][]*models.PointEntering = make([][]*models.PointEntering, len(*wa.ClippedPolygon.Points))
	var nextIndex, previousIndex, n int
	wa.CheckPointsInPerimeter()
	wa.GetSegmentIntersections()
	wa.SortClippingPolygonIntersections()

	// Building preliminar clipping_polygon_list and clipped_polygon_list.
	for i := 0; i < len(*wa.ClippingPolygon.Points); i++ {
		wa.ClippingPolygonList = append(wa.ClippingPolygonList, models.NewDefaultPointEntering((*wa.ClippingPolygon.Points)[i], false))

		for j := 0; j < len(wa.BorderIntersections[i]); j++ {
			BP = wa.BorderIntersections[i][j]
			P = BP.P
			isIncomplete, isEntering, isSingle = true, false, false

			if !BP.IsBorder || (BP.IsSegment && !BP.IsEnd) {
				// Point doesn't rely on following points to know wheter is entering or not to clippedPolygon
				isIncomplete = false
			}

			if BP.IsBorder && !BP.IsSegment {
				// It's just one point that lies on the border
				isSingle = true
			}

			if !isIncomplete {
				// we could determine wheter this point it's entering or not to clippedPolygon
				var aux int = len(wa.ClippingPolygonList) - 1
				for wa.ClippingPolygonList[aux].P.Equal(P) {
					// We check the first point that its different from current one. The first point on the list is outside the clippedPolygon, so the loop will always stop.
					aux--
				}

				// Entering is false in case the last point different from current one lies inside the polygon, true otherwise.
				isEntering = !wa.ClippedPolygon.PointInPolygon(wa.ClippingPolygonList[aux].P)
			}

			PE = models.NewPointEntering(P, true, isEntering, isIncomplete, isSingle)
			wa.ClippingPolygonList = append(wa.ClippingPolygonList, PE)
			clippedIntersections[BP.Index] = append(clippedIntersections[BP.Index], PE)
		}
	}

	n = len(wa.ClippingPolygonList)
	// Completing ClippingPolygonList
	for i := 0; i < len(wa.ClippingPolygonList); i++ {
		PE = wa.ClippingPolygonList[i]

		if !PE.IsIncomplete {
			continue
		}

		PE.IsIncomplete = false
		nextIndex = (i + 1) % n
		for wa.ClippingPolygonList[nextIndex].P.Equal(PE.P) {
			// Grabs the first following point that it's different from current one.
			nextIndex = (nextIndex + 1) % n
		}

		previousIndex = (i - 1 + n) % n
		for wa.ClippingPolygonList[previousIndex].P.Equal(PE.P) {
			// Grabs the first previous point that it's different from current one.
			previousIndex = (previousIndex - 1 + n) % n
		}

		if PE.IsSingle {
			lastPoint = wa.ClippedPolygon.PointInPolygon(wa.ClippingPolygonList[previousIndex].P)
			nextPoint = wa.ClippedPolygon.PointInPolygon(wa.ClippingPolygonList[nextIndex].P)

			if lastPoint != nextPoint {
				// Point is indeed an intersection
				wa.ClippingPolygonList[i].IsEntering = !lastPoint
				continue
			}

			// Point should not be considered as intersection
			wa.ClippingPolygonList[i].Intersection = false
			continue
		}

		nextPoint = wa.ClippedPolygon.PointInPolygon(wa.ClippingPolygonList[nextIndex].P)
		wa.ClippingPolygonList[i].IsEntering = nextPoint
	}

	// We sort clippedPolygonList
	for i := 0; i < len(*wa.ClippedPolygon.Points); i++ {
		P = (*wa.ClippedPolygon.Points)[i]
		PE = models.NewDefaultPointEntering(P, false)
		wa.ClippedPolygonList = append(wa.ClippedPolygonList, PE)

		sort.Slice(clippedIntersections[i], func(x, y int) bool {
			var a, b float64
			A, B = clippedIntersections[i][x].P, clippedIntersections[i][y].P
			a, b = models.VectorNorm(A, P), models.VectorNorm(B, P)
			return a < b
		})

		wa.ClippedPolygonList = append(wa.ClippedPolygonList, clippedIntersections[i]...)
	}
}

/*
Returns every (if any) intersection polygon in 'ClippedPolygon' and 'ClippingPolygon'. 'ClippedPolygon' should
always be a convex polygon, i.e. the camera_vision.
*/
func (wa *WeilerAtherton) GetIntersectionPolygons() []*models.Polygon {
	var idx, n, m int
	var turn int8
	var seenPoints []bool
	var initialPoint *models.PointEntering
	var intersectionPolygon *models.Polygon
	var intersectionPolygons []*models.Polygon
	var notExit bool
	wa.ClippingPolygon.SortClockwise()

	if wa.ClippingPolygon.InsidePolygon(wa.ClippedPolygon) {
		wa.ClippingPolygon.DeleteCollinearSegments()
		return []*models.Polygon{wa.ClippingPolygon}
	}

	wa.ClippingPolygon.DeleteCollinearSegments()

	// We will rotate clipping_polygon till the first point is not inside the polygon.
	// We could actually do this 'cause, as a previous step, we did verify that the polygon its not contained.
	// Thats the reason why it has at least one point outside, and its the one of interest.
	idx = 0
	for wa.ClippedPolygon.PointInPolygon((*wa.ClippingPolygon.Points)[idx]) {
		idx++
	}
	wa.ClippingPolygon.CyclicRotation(idx)

	wa.BuildClippingAndClippedPolygonLists()
	n, m = len(wa.ClippingPolygonList), len(wa.ClippedPolygonList)
	seenPoints = make([]bool, n)

	for i := 0; i < n; i++ {
		if !wa.ClippingPolygonList[i].Intersection || !wa.ClippingPolygonList[i].IsEntering || seenPoints[i] {
			continue
		}

		initialPoint = wa.ClippingPolygonList[i]
		intersectionPolygon = models.NewPolygon(initialPoint.P)
		seenPoints[i], notExit = true, true
		turn = 1
		idx = (i + 1) % n

		for notExit {
			if turn == 1 {
				if seenPoints[idx] {
					break
				}

				if !(wa.ClippingPolygonList[idx].Intersection && !wa.ClippingPolygonList[idx].IsEntering) {
					// We didn't find exit from ClippingPolygon list.
					// Current point of ClippingPolygon list should be added to intersection polygon.
					*intersectionPolygon.Points = append(*intersectionPolygon.Points, wa.ClippingPolygonList[idx].P)
					seenPoints[idx] = true
					idx = (idx + 1) % n
					continue
				}

				// We found exit from ClippingPolygon list into ClippedPolygon list.
				for j := 0; j < m; j++ {
					if !wa.ClippedPolygonList[j].Equal(wa.ClippingPolygonList[idx]) {
						continue
					}

					// We found exit point inside ClippedPolygon list.
					*intersectionPolygon.Points = append(*intersectionPolygon.Points, wa.ClippedPolygonList[j].P)
					idx = (j + 1) % m
					turn = 1 - turn
					break
				}
			} else {
				// Found initial point.
				if initialPoint.Equal(wa.ClippedPolygonList[idx]) {
					break
				}

				*intersectionPolygon.Points = append(*intersectionPolygon.Points, wa.ClippedPolygonList[idx].P)
				if !(wa.ClippedPolygonList[idx].Intersection && wa.ClippedPolygonList[idx].IsEntering) {
					// We didn't find exit point.
					idx = (idx + 1) % m
					continue
				}

				for j := 0; j < n; j++ {
					if !wa.ClippingPolygonList[j].Equal(wa.ClippedPolygonList[idx]) {
						continue
					}

					if seenPoints[j] {
						// It's a loop of points.
						notExit = false
						break
					}

					seenPoints[j] = true
					turn = 1 - turn
					idx = (j + 1) % n
					break
				}
			}
		}

		if len(*intersectionPolygon.Points) == 1 {
			// This could happen due to Breaking Loop causing some special handling polygon.
			continue
		}
		intersectionPolygon.DeleteRepeatedPoints()
		intersectionPolygon.DeleteCollinearSegments()
		intersectionPolygons = append(intersectionPolygons, intersectionPolygon)
	}

	if len(intersectionPolygons) == 0 {
		// It's possible that ClippedPolygon it's inside clipping_polygon.
		if wa.ClippedPolygon.InsidePolygon(wa.ClippingPolygon) {
			return []*models.Polygon{wa.ClippedPolygon}
		}
	}

	return intersectionPolygons
}

// Calculates the polygon intersections (if any) between camera footprint of 'camera' and some 'polygon' defined as
// interest zone. Make sure to pass a copy of Polygon to this function.
func CalculatePolygonIntersection(camera *models.Camera, polygon *models.Polygon) (bool, *[]*models.Polygon) {
	var cameraFootprint *models.Polygon = camera.CameraFootprint.ToWeilerAthertonRepresentation()
	var copyPolygon *models.Polygon = polygon.Copy()
	var wa *WeilerAtherton = NewWeilerAtherton(copyPolygon, cameraFootprint)

	var intersectionPolygons []*models.Polygon = wa.GetIntersectionPolygons()

	if len(intersectionPolygons) == 0 {
		return false, nil
	}

	return true, &intersectionPolygons
}
