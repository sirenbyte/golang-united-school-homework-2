package golang_united_school_homework_2

import "math"

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	if sidesNum == SidesCircle {
		return 2 * sideLen * math.Pi
	}
	if sidesNum == SidesSquare {
		return sideLen * sideLen
	}
	if sidesNum == SidesTriangle {
		return ((sideLen * sideLen) * math.Sqrt(3)) / 4
	}
	return 0
}
