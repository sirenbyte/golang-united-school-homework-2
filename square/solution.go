package square

import "math"

type IntCustomType int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum IntCustomType) float64 {
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
