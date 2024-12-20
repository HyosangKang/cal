package cal

import "fmt"

func ExampleFunction_LeftSum() {
	f := Function{
		Interval: [2]float64{0., 1.},
		Formula:  func(x float64) float64 { return x * x },
	}
	sub := 10
	for i := 0; i < 3; i++ {
		fmt.Printf("%.5f\n", f.LeftSum(sub))
		sub *= 10
	}
	// Output:
	// 0.28500
	// 0.32835
	// 0.33283
}
