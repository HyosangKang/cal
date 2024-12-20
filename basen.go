package cal

func BaseN(b, n int) []int {
	var d []int
	for b > 0 {
		d = append(d, b%n)
		b /= n
	}
	return d
}
