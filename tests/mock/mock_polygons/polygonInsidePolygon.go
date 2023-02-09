package mock_polygons

// Weird case with Polygon 11 inside Polygon 4.

var (
	PolygonsContainPolygon     []bool = []bool{true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon1           = []bool{false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}
	PolygonsContainPolygon2           = []bool{false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}
	PolygonsContainPolygon3           = []bool{false, false, false, true, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon4           = []bool{false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon5           = []bool{true, false, false, false, false, true, true, false, true, false, false, false, false, true, false, true, true, true, true, false, false, false, false, false, false, false}
	PolygonsContainPolygon6           = []bool{false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon7           = []bool{false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon8           = []bool{false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon9           = []bool{false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon10          = []bool{true, false, false, false, true, false, true, false, true, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, true, false}
	PolygonsContainPolygon11          = []bool{false, false, false, true, true, false, true, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon12          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon13          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon14          = []bool{false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon15          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, true, false, false, false, false, false, false, false}
	PolygonsContainPolygon16          = []bool{false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon17          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false}
	PolygonsContainPolygon18          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false}
	PolygonsContainPolygon19          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false}
	PolygonsContainPolygon20          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false}
	PolygonsContainPolygon21          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false}
	PolygonsContainPolygon22          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false}
	PolygonsContainPolygon23          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false}
	PolygonsContainPolygon24          = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false}
	PolygonsContainPolygon_ccw        = []bool{false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}

	AllPolygonsContainPolygons [][]bool = [][]bool{PolygonsContainPolygon, PolygonsContainPolygon1, PolygonsContainPolygon2, PolygonsContainPolygon3, PolygonsContainPolygon4, PolygonsContainPolygon5, PolygonsContainPolygon6, PolygonsContainPolygon7, PolygonsContainPolygon8, PolygonsContainPolygon9, PolygonsContainPolygon10, PolygonsContainPolygon11, PolygonsContainPolygon12, PolygonsContainPolygon13, PolygonsContainPolygon14, PolygonsContainPolygon15, PolygonsContainPolygon16, PolygonsContainPolygon17, PolygonsContainPolygon18, PolygonsContainPolygon19, PolygonsContainPolygon20, PolygonsContainPolygon21, PolygonsContainPolygon22, PolygonsContainPolygon23, PolygonsContainPolygon24, PolygonsContainPolygon_ccw}
)
