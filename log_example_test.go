package cal

import (
	"fmt"
	"math"
)

func ExampleLogBisect() {
	fmt.Printf("Log(2) ~ %v \t (Bisect)\n", LogByBisect(2))
	fmt.Printf("Log(2) ~ %v \t (math.Log)\n", math.Log(2))
	// Output:
	// 	Log(2) ~ 0.6931471804592442 	 (Series)
	// Log(2) ~ 0.6931471805599453 	 (math.Log)
}

func ExampleLogSeries() {
	fmt.Printf("Log(2) ~ %v \t (Series)\n", Log(2))
	fmt.Printf("Log(2) ~ %v \t (math.Log)\n", math.Log(2))
	// Output:
	// 	Log(2) ~ 0.6931471804592442 	 (Series)
	// Log(2) ~ 0.6931471805599453 	 (math.Log)
}
