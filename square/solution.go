package square

import "math"

func CalcSquare(sideLen float64, sidesNum string) float64 {
	if sidesNum == "SidesCircle" {
		return 2 * sideLen * math.Pi
	}
	if sidesNum == "SidesSquare" {
		return sideLen * sideLen
	}
	if sidesNum == "SidesTriangle" {
		return ((sideLen * sideLen) * math.Sqrt(3)) / 4
	}
	return 0
}
