package cal

func LogByBisect(x float64) float64 {
	var a, b, c float64 = -10, 10, 0
	f := func(t float64) float64 { return Exp(t) - x }
	for !IsZero(b - a) {
		c = (a + b) / 2
		if f(c) > 0 {
			b = c
		} else if f(c) < 0 {
			a = c
		} else {
			return c
		}
	}
	return c
}

func log(t float64) float64 {
	if t < -1 || t > 1 {
		panic("log: invalid input")
	}
	var s, p float64 = 1, t
	for i := 1; !IsZero(p); i++ {
		s += p
		p *= t * float64(i) / float64(i+1)
	}
	return s
}

func Log(x float64) float64 {
	if x < 0 || IsZero(x) {
		panic("log: invalid input")
	}
	t := (x - 1) / (x + 1)
	return log(t) - log(-t)
}
