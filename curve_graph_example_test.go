package cal

import "fmt"

func ExampleCurve_UnitCircle() {
	var n, c int = 100, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := -1. + 2*float64(i)/float64(n)
			y := -1. + 2*float64(j)/float64(n)
			if x*x+y*y <= 1 || IsZero(x*x+y*y-1) {
				c++
			}
		}
	}
	fmt.Println(c)
	fmt.Println(float64(c) / float64(n*n) * 4)
	// Output: 625
}
