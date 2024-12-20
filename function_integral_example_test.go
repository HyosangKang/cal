package cal

import "fmt"

func ExampleFunction_Integral() {
	f := Function{
		Interval: [2]float64{-1 / Sqrt(2), 1 / Sqrt(2)},
		Formula: func(x float64) float64 {
			return 1 / Sqrt(1-x*x)
		},
	}
	fmt.Printf("%.4f\n", 4*f.Integral())
	// Output: 6.2832
}

func ExampleFunction_Integeral_area() {
	f := Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			return 2 * Sqrt(1-x*x)
		},
	}
	fmt.Printf("%.4f\n", f.Integral())
	// Output: 3.1416
}
