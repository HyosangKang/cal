package cal

import (
	"fmt"
	"math"
)

var Cycloid = Curve{
	Interval: [2]float64{0, 2 * math.Pi},
	X: func(t float64) float64 {
		return t - Sin(t)
	},
	Y: func(t float64) float64 {
		return 1 - Cos(t)
	},
}

func ExampleCurve_Length() {
	fmt.Printf("%.4f\n", Cycloid.Length())
	// Output: 7.4256
}
