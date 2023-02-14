package models

/*
Defines wheter or not a point 'P' is a intersection point between a pair of segments, furthermore, in case its
an intersection the direction of the vector is defined by 'entering'. Aditionally, the intersection point's
direction may not be defined at the beggining, thats where 'incomplete' parameter comes in. 'isSingle' is related
to this reason as well, because the direction of entering could rely on next points that are not known.
*/
type PointEntering struct {
	P                                                *Point
	Intersection, IsEntering, IsIncomplete, IsSingle bool
}

/*
Constructor of PointEntering
*/
func NewPointEntering(P *Point, IsIntersection, IsEntering, Incomplete, IsSingle bool) *PointEntering {
	return &PointEntering{
		P:            P,
		Intersection: IsIntersection,
		IsEntering:   IsEntering,
		IsIncomplete: Incomplete,
		IsSingle:     IsSingle,
	}
}

/*
Constructor with default parameteres

Variables IsEntering, IsIncomplete and IsSingle are initialized to false
*/
func NewDefaultPointEntering(P *Point, Intersection bool) *PointEntering {
	return NewPointEntering(P, Intersection, false, false, false)
}

/*
Recieves a PointEntering and return true if both are equal, false otherwise
*/
func (Pe *PointEntering) Equal(Po *PointEntering) bool {
	return Pe.P.Equal(Po.P) && Pe.IsEntering == Po.IsEntering && Pe.Intersection == Po.Intersection
}
