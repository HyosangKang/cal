package cal

func ExpSequence(x float64, n int) float64 {
	if n < 0 {
		return 0
	}
	d := BaseN(n, 2)
	var e, f float64 = 1, 1 + x/float64(n)
	for i, k := range d {
		if k == 1 {
			e *= IterativeSquare(f, i)
		}
	}
	return e
}

func Exp(x float64) float64 {
	var f, p float64 = 1, 1
	for i := 0; !IsZero(p); i++ {
		p *= x / float64(i+1)
		f += p
	}
	return f
}
