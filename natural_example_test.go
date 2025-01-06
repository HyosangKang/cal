package cal

import (
	"fmt"
)

func ExampleNaturalSequence() {
	var n int = 1
	for i := 0; i < 8; i++ {
		a := NaturalSequence(n)
		fmt.Printf("%.3f\t %.4f\t (n = %.e)\n",
			a, a, float64(n))
		n *= 10
	}
	// Output:
	// 2.000	 2.0000	 (n = 1e+00)
	// 2.594	 2.5937	 (n = 1e+01)
	// 2.705	 2.7048	 (n = 1e+02)
	// 2.717	 2.7169	 (n = 1e+03)
	// 2.718	 2.7181	 (n = 1e+04)
	// 2.718	 2.7183	 (n = 1e+05)
	// 2.718	 2.7183	 (n = 1e+06)
	// 2.718	 2.7183	 (n = 1e+07)
}

func ExampleNaturalSeries() {
	fmt.Printf("%.6f (sequence)\n", NaturalSequence(10))
	fmt.Printf("%.6f (series)\n", NaturalSeries(10))
	// Output:
	// 2.593742 (sequence)
	// 2.718282 (series)
}
