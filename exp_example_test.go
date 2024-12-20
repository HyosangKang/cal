package cal

import (
	"fmt"
)

func ExampleExpSequence() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%.5f\n", ExpSequence(-5, i))
	}
	// Output:
	// 1.00000
	// -4.00000
	// 2.25000
	// -0.29630
	// 0.00391
	// 0.00000
	// 0.00002
	// 0.00016
	// 0.00039
	// 0.00068
}

func ExampleExpSeries() {
	for i := 5; i < 10; i++ {
		fmt.Println(Exp(2))
	}
	// Output:
}
