package cal

import "math"

type Curve struct {
	Interval [2]float64
	X, Y     func(float64) float64
}

var UnitCircle = Curve{
	Interval: [2]float64{0, 2 * math.Pi},
	X: func(t float64) float64 {
		return Cos(t)
	},
	Y: func(t float64) float64 {
		return Sin(t)
	},
}
