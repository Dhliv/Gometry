package models

import (
	"fmt"
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

type Camera struct {
	X, Y, Z         float64
	Yaw, Pitch      float64
	lens            *CMOS_Lens
	CameraFootprint *Quadrilateral
}

/*
Camera constructor. Relies on the position of the camera, its Yaw and Pitch angles, the format of the lens and FOV of the camera. Take note that 'pitch' angle should ∈ [-90, 0].

Returns:

  - Camera struct and nil if 'format' of lens was available
  - nil and pitch angle error if pitch ∉ [-90, 0]
  - nil and sensor not found error otherwise
*/
func NewCamera(x, y, z, yaw, pitch float64, format string, FOV float64) (*Camera, error) {
	if pitch < float64(-90) || pitch > float64(0) {
		// TODO handle exception
		return nil, fmt.Errorf("pitch angle should ∈ [-90, 0], got %v instead", pitch)
	}

	camera := &Camera{
		X: x, Y: y, Z: z,
		Yaw: yaw, Pitch: pitch,
	}

	lens, err := NewCMOS_Lens(format, FOV)

	if err != nil {
		// TODO handle the exception
		return nil, err
	}

	camera.lens = lens
	camera.getCameraFootprint()
	return camera, nil
}

/*
By using the angles 'pitch', 'HFOV' and 'VFOV' and the height
from the drone to the ground, we project the area that the drone sees in the drone's coordinate
system, taking z=0 in all the quadrilateral's points. This projection is called Camera Footprint.

Take note of angles, here they are taken in counter-clockwise, starting from x-axis for rotations
made in z-axis and starting from x-axis for rotations made in y-axis. Look Yaw, Pitch, Roll for reference.
*/
func (Cam *Camera) getCameraFootprint() {
	var minY, minX, maxY, maxX float64
	var maxθY float64
	var A, B, C, D *Point

	minY = -Cam.Z / math.Tan(COMPLETE_CIRCUNFERENCE+ONE_GRADE_IN_RADIANS*Cam.Pitch-ONE_GRADE_IN_RADIANS*Cam.lens.VFOV/2)

	maxθY = COMPLETE_CIRCUNFERENCE + ONE_GRADE_IN_RADIANS*Cam.Pitch + ONE_GRADE_IN_RADIANS*Cam.lens.VFOV/2
	if maxθY >= COMPLETE_CIRCUNFERENCE {
		// ! Angle cant be greater or equal than 360°
		maxθY = COMPLETE_CIRCUNFERENCE - 0.01
	}

	maxY = -Cam.Z / math.Tan(maxθY)
	if maxY >= (minY + float64(MAX_DISTANCE)) {
		// ! Maximum distance can't be greater that MAX_DISTANCE meters from the min-y measurement.
		maxY = (minY + float64(MAX_DISTANCE))
	}

	minX = math.Abs(minY) * math.Tan(ONE_GRADE_IN_RADIANS*Cam.lens.HFOV/2)
	maxX = maxY * math.Tan(ONE_GRADE_IN_RADIANS*Cam.lens.HFOV/2) // ! maxY cant be negative by restrictions of Pitch

	A = NewPoint(-maxX, maxY)
	B = NewPoint(maxX, maxY)
	C = NewPoint(-minX, minY)
	D = NewPoint(minX, minY)

	cameraFootprint := NewQuadrilateral(A, B, C, D)
	cameraFootprint.Translate(NewPoint(Cam.X, Cam.Y))
	cameraFootprint.Rotate(Cam.Yaw, true)

	Cam.CameraFootprint = cameraFootprint
}
