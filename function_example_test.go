package cal

import (
	"math"
)

var (
	SinxX = Function{
		Interval: [2]float64{-10, 10},
		Formula: func(x float64) float64 {
			if IsZero(x) {
				return 1
			}
			return Sin(x) / x
		},
	}
	Exotic = Function{
		Interval: [2]float64{-2, 2},
		Formula: func(x float64) float64 {
			return math.Exp(-1 / (x * x))
		},
	}
	X2Sin = Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			if IsZero(x) {
				return 0
			}
			return x * x * Sin(1/x)
		},
	}
	X3Sin = Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			if IsZero(x) {
				return 0
			}
			return x * x * x * Sin(1/x)
		},
	}
	Sine = Function{
		Interval: [2]float64{-3, 3},
		Formula:  Sin,
	}
)
