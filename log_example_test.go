package cal

import (
	"fmt"
	"math"
)

func ExampleLogBisect() {
	fmt.Printf("Log(2) ~ %v \t (Bisect)\n", LogByBisect(2))
	fmt.Printf("Log(2) ~ %v \t (math.Log)\n", math.Log(2))
	// Output:
}

func ExampleLogSeries() {
	fmt.Printf("Log(2) ~ %v \t (Series)\n", Log(2))
	fmt.Printf("Log(2) ~ %v \t (math.Log)\n", math.Log(2))
	// Output:
}
