package cal

import "fmt"

func ExampleIsZero() {
	var eps = Eps / 5
	fmt.Println(IsZero(Eps))
	fmt.Println(IsZero(5 * eps))
	// Output:
}
