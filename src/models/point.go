package models

import (
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

// Point instance ind 2d coordinate system.
type Point struct {
	x float64
	y float64
}

/*
Point costructor
*/
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Rotates the actual point by 'beta' degrees in counter-clocwise.
func (P *Point) Rotate(beta float64) {
	x := math.Cos(beta)*P.x - math.Sin(beta)*P.y
	y := math.Sin(beta)*P.x + math.Cos(beta)*P.y
	P.x = x
	P.y = y
}

// Translate the origin point by 'translation_point'.
func (P *Point) Translate(translationPoint *Point) {
	P.x += translationPoint.x
	P.y += translationPoint.y
}

// Recieves a Point and return true if both are equal, false otherwise
func (Pe *Point) Equal(P *Point) bool {
	if math.Abs(Pe.x-P.x) > EPSILON || math.Abs(Pe.y-P.y) > EPSILON {
		return false
	}
	return true
}

func (Pe *Point) DotProduct(P *Point) float64 {
	return Pe.x*P.x + Pe.y*P.y
}
