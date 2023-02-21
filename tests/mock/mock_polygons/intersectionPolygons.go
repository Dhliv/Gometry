package mock_polygons

import (
	"github.com/Dhliv/Gometry/src/models"
	. "github.com/Dhliv/Gometry/tests/mock/mock_points"
)

var (
	IntersectionPolygon     []*models.Polygon = []*models.Polygon{models.NewPolygon(A, B, D, C)}
	IntersectionPolygon1                      = []*models.Polygon{models.NewPolygon(S3, models.NewPoint(3.3636363636364, 8.6363636363636), models.NewPoint(5.1818181818182, 6.8181818181818)), models.NewPolygon(S5, models.NewPoint(6, 6), models.NewPoint(6.8, 5.2))}
	IntersectionPolygon2                      = []*models.Polygon{models.NewPolygon(S3, models.NewPoint(3.3636363636364, 8.6363636363636), S13)}
	IntersectionPolygon3                      = []*models.Polygon{models.NewPolygon(A, B, models.NewPoint(7.741935483871, 2.5806451612903), models.NewPoint(2.1621621621622, 6.4864864864865))}
	IntersectionPolygon4                      = []*models.Polygon{models.NewPolygon(S3, S13, B, models.NewPoint(6.8181818181818, 2.2727272727273), models.NewPoint(1.4705882352941, 4.4117647058824), A, models.NewPoint(3.3636363636364, 8.6363636363636))}
	IntersectionPolygon5                      = []*models.Polygon{Polygon5.Copy()}
	IntersectionPolygon6                      = []*models.Polygon{Polygon.Copy()}
	IntersectionPolygon7                      = []*models.Polygon{models.NewPolygon(models.NewPoint(5.4545454545455, 6.5454545454545), B, D, models.NewPoint(1.8181818181818, 2.1818181818182))}
	IntersectionPolygon8                      = []*models.Polygon{models.NewPolygon(S11, A, B, models.NewPoint(6.2307692307692, 2.0769230769231))}
	IntersectionPolygon9                      = []*models.Polygon{models.NewPolygon(S3, A, B, models.NewPoint(7.4210526315789, 2.4736842105263))}
	IntersectionPolygon10                     = []*models.Polygon{models.NewPolygon(A, B, S5, S3)}
	IntersectionPolygon11                     = []*models.Polygon{models.NewPolygon(B, A)}
	IntersectionPolygon12                     = []*models.Polygon{models.NewPolygon(S14, S5, models.NewPoint(6, 6)), models.NewPolygon(A, S13)}
	IntersectionPolygon13                     = []*models.Polygon{models.NewPolygon(D, A, models.NewPoint(8.2727272727273, 3.7272727272727), S5, models.NewPoint(7.5652173913043, 2.5217391304348))}
	IntersectionPolygon14                     = []*models.Polygon{models.NewPolygon(models.NewPoint(6.1111111111111, 5.8888888888889), B, D)}
	IntersectionPolygon15                     = []*models.Polygon{models.NewPolygon(A, B, D)}
	IntersectionPolygon16                     = []*models.Polygon{models.NewPolygon(models.NewPoint(1.9230769230769, 5.7692307692308), S13, B, D, C)}
	IntersectionPolygon17                     = []*models.Polygon{models.NewPolygon(models.NewPoint(6.2307692307692, 2.0769230769231), S11, S3, models.NewPoint(2, 6), A, B)}
	IntersectionPolygon18                     = []*models.Polygon{models.NewPolygon(models.NewPoint(8.2727272727273, 3.7272727272727), S5, models.NewPoint(7.5652173913043, 2.5217391304348), D, C, A)}
	IntersectionPolygon19                     = []*models.Polygon{models.NewPolygon(models.NewPoint(6.4285714285714, 2.1428571428571), S11, S5, S14, B)}
	IntersectionPolygon20                     = []*models.Polygon{}
	IntersectionPolygon21                     = []*models.Polygon{models.NewPolygon(S22, models.NewPoint(6, 6), models.NewPoint(7, 5), S23, D)}
	IntersectionPolygon22                     = []*models.Polygon{}
	IntersectionPolygon23                     = []*models.Polygon{}
	IntersectionPolygon24                     = []*models.Polygon{models.NewPolygon(S3, A, S13), models.NewPolygon(B, S5, S14)}
	IntersectionPolygon_ccw                   = []*models.Polygon{models.NewPolygon(S3, models.NewPoint(3.3636363636364, 8.6363636363636), models.NewPoint(5.1818181818182, 6.8181818181818)), models.NewPolygon(S5, models.NewPoint(6, 6), models.NewPoint(6.8, 5.2))}

	AllIntersectionPolygons [][]*models.Polygon = [][]*models.Polygon{IntersectionPolygon, IntersectionPolygon1, IntersectionPolygon2, IntersectionPolygon3, IntersectionPolygon4, IntersectionPolygon5, IntersectionPolygon6, IntersectionPolygon7, IntersectionPolygon8, IntersectionPolygon9, IntersectionPolygon10, IntersectionPolygon11, IntersectionPolygon12, IntersectionPolygon13, IntersectionPolygon14, IntersectionPolygon15, IntersectionPolygon16, IntersectionPolygon17, IntersectionPolygon18, IntersectionPolygon19, IntersectionPolygon20, IntersectionPolygon21, IntersectionPolygon22, IntersectionPolygon23, IntersectionPolygon24, IntersectionPolygon_ccw}
)
