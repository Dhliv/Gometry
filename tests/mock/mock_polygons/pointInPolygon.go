package mock_polygons

var (
	PointsInPolygon     []bool = []bool{true, true, true, true, false, false, true, false, true, false, false, false, false, false, true, true, true, true, false, false, false, false, false, false, false, true, true, false}
	PointsInPolygon1           = []bool{false, false, false, false, true, true, true, true, true, true, true, false, false, false, false, false, true, false, false, true, false, false, false, false, false, false, false, true}
	PointsInPolygon2           = []bool{false, false, false, false, true, true, true, false, false, true, true, false, false, false, false, false, true, false, false, true, false, false, false, false, false, false, false, false}
	PointsInPolygon3           = []bool{true, true, false, false, true, true, false, true, false, true, false, true, false, false, false, true, true, true, false, false, false, false, false, false, false, false, false, true}
	PointsInPolygon4           = []bool{true, true, false, false, false, true, true, true, true, true, false, true, true, false, false, true, true, true, false, false, false, false, false, false, false, false, false, true}
	PointsInPolygon5           = []bool{false, false, false, false, false, false, true, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false}
	PointsInPolygon6           = []bool{true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, false, false, false, true, false, false, false, true, true, true}
	PointsInPolygon7           = []bool{false, true, false, true, false, false, false, false, true, true, false, true, false, true, true, false, false, true, false, false, false, false, false, false, false, true, true, true}
	PointsInPolygon8           = []bool{true, true, false, false, false, false, true, true, true, true, false, true, false, false, true, true, true, true, false, false, false, false, false, false, false, false, false, true}
	PointsInPolygon9           = []bool{true, true, false, false, false, false, true, true, true, true, false, true, false, false, false, true, true, true, true, false, false, false, false, false, false, false, false, true}
	PointsInPolygon10          = []bool{true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, true, true, true, false, false, false, false, false, false, false, false, false, false}
	PointsInPolygon11          = []bool{true, true, false, false, false, true, false, true, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false, false, false, false, true}
	PointsInPolygon12          = []bool{true, true, false, false, false, true, false, true, true, false, false, false, false, false, false, true, true, true, true, false, false, false, false, false, false, false, false, false}
	PointsInPolygon13          = []bool{true, false, false, true, false, false, true, true, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, false, false, false, true, false}
	PointsInPolygon14          = []bool{false, true, false, true, false, false, false, false, true, true, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, true, true}
	PointsInPolygon15          = []bool{true, true, false, true, false, false, true, false, true, false, false, true, false, false, true, true, true, true, true, false, false, false, false, false, false, false, true, false}
	PointsInPolygon16          = []bool{false, true, true, true, false, false, true, false, true, false, false, false, true, false, true, false, true, true, false, false, false, true, false, false, false, true, true, false}
	PointsInPolygon17          = []bool{true, true, false, false, true, false, true, false, true, true, false, true, false, false, true, true, true, true, true, false, false, false, false, false, false, false, false, false}
	PointsInPolygon18          = []bool{true, false, true, true, false, true, true, true, true, false, false, true, false, true, true, true, true, true, true, false, false, false, true, false, false, true, true, false}
	PointsInPolygon19          = []bool{false, true, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, true, false, false, false, false, false, true, false, false, false, false}
	PointsInPolygon20          = []bool{true, false, false, false, false, false, false, false, false, false, true, false, true, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false}
	PointsInPolygon21          = []bool{false, false, true, true, false, false, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, true, false, false, true, true, true, true}
	PointsInPolygon22          = []bool{false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false}
	PointsInPolygon23          = []bool{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, true, false, false, true, false, false, true, false}
	PointsInPolygon24          = []bool{true, true, false, false, false, false, true, true, true, true, false, false, false, false, false, true, true, true, false, false, false, false, false, true, false, false, false, true}
	PointsInPolygon_ccw        = []bool{false, false, false, false, true, true, true, true, true, true, true, false, false, false, false, false, true, false, false, true, false, false, false, false, false, false, false, true}

	AllPointsInPolygons [][]bool = [][]bool{PointsInPolygon, PointsInPolygon1, PointsInPolygon2, PointsInPolygon3, PointsInPolygon4, PointsInPolygon5, PointsInPolygon6, PointsInPolygon7, PointsInPolygon8, PointsInPolygon9, PointsInPolygon10, PointsInPolygon11, PointsInPolygon12, PointsInPolygon13, PointsInPolygon14, PointsInPolygon15, PointsInPolygon16, PointsInPolygon17, PointsInPolygon18, PointsInPolygon19, PointsInPolygon20, PointsInPolygon21, PointsInPolygon22, PointsInPolygon23, PointsInPolygon24, PointsInPolygon_ccw}
)
