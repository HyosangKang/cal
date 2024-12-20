package cal

import "math"

var CycloidCircle = func(p float64) Curve {
	return Curve{
		Interval: [2]float64{0, 2 * math.Pi},
		X: func(t float64) float64 {
			return p + Cos(t)
		},
		Y: func(t float64) float64 {
			return Sin(t) + 1
		},
	}
}
