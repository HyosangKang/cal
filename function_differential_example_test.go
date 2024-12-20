package cal

import "fmt"

func ExampleFunction_Differential() {
	f := Function{
		Interval: [2]float64{-1., 1.},
		Formula:  Cos,
	}
	fmt.Println(f.Differential(0.0))
	// Output: 0
}
