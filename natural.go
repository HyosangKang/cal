package cal

func IterativeSquare(a float64, n int) float64 {
	var x float64 = a
	for i := 0; i < n; i++ {
		x *= x
	}
	return x
}

func NaturalSequence(n int) float64 {
	var e, f float64 = 1, 1 + 1/float64(n)
	d := BaseN(n, 2)
	for i, k := range d {
		if k == 1 {
			e *= IterativeSquare(f, i)
		}
	}
	return e
}

func NaturalSeries(n int) float64 {
	var f, p float64 = 1, 1
	for i := 0; i < n; i++ {
		p /= float64(i + 1)
		f += p
	}
	return f
}
