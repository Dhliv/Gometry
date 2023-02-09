package models

import (
	"fmt"
	"math"

	. "github.com/Dhliv/Gometry/src/utils"
)

type SensorDimensions struct {
	Width, Height, Diagonal float64
}

var formatMap map[string]SensorDimensions = map[string]SensorDimensions{
	"1/4":   {Diagonal: 5, Width: 3.6, Height: 2.7},
	"1/3.6": {Diagonal: 5, Width: 4, Height: 3},
	"1/3.2": {Diagonal: 5.68, Width: 4.536, Height: 3.416},
	"1/3":   {Diagonal: 6, Width: 4.8, Height: 3.6},
	"1/2.7": {Diagonal: 6.721, Width: 5.371, Height: 4.035},
	"1/2.5": {Diagonal: 7.182, Width: 5.76, Height: 4.29},
	"1/2":   {Diagonal: 8, Width: 6.4, Height: 4.8},
	"1/1.8": {Diagonal: 8.933, Width: 8.5, Height: 6.8},
	"1/1.7": {Diagonal: 9.5, Width: 7.6, Height: 5.7},
	"2/3":   {Diagonal: 11, Width: 8.8, Height: 6.6},
	"1":     {Diagonal: 16, Width: 12.8, Height: 9.6},
}

type CMOS_Lens struct {
	Format          string
	FOV, HFOV, VFOV float64
}

// Calculates the vertical and horizontal FOV from camera sensor specifications.
//
// Returns the vertical and horizontal field of view respectively.
func calcVFOV_HFOV(FOV float64, sensor *SensorDimensions) (VFOV float64, HFOV float64) {
	var focalLenght float64 = (sensor.Diagonal / (2 * math.Tan(ONE_GRADE_IN_RADIANS*FOV/2))) * ONE_RADIAN_IN_GRADES
	HFOV = 2 * math.Atan(sensor.Width/focalLenght) * ONE_RADIAN_IN_GRADES
	VFOV = 2 * math.Atan(sensor.Height/focalLenght) * ONE_RADIAN_IN_GRADES
	return
}

/*
Constructor for CMOS_lens. Relies on format to get an specified sensor and FOV to calculate HFOV and VFOV.

Returns: CMOS_Lens and nil if sensor has been found, nil and sensor not found error otherwise.
*/
func NewCMOS_Lens(format string, FOV float64) (*CMOS_Lens, error) {
	sensor, ok := formatMap[format]

	if !ok {
		return nil, fmt.Errorf("sensor with format (%v) not found", format)
	}

	VFOV, HFOV := calcVFOV_HFOV(FOV, &sensor)

	return &CMOS_Lens{
		FOV: FOV, VFOV: VFOV, HFOV: HFOV,
		Format: format,
	}, nil
}
