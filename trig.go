package cal

const Pi = 3.14159265358979323846

func Mod2Pi(x float64) float64 {
	for x < -Pi {
		x += 2 * Pi
	}
	for x > Pi {
		x -= 2 * Pi
	}
	return x
}

func Cos(x float64) float64 {
	x = Mod2Pi(x)
	f, p := 0.0, 1.0
	for i := 0; !IsZero(p); i++ {
		f += p
		p *= -x * x / float64((2*i+1)*(2*i+2))
	}
	return f
}

func Sin(x float64) float64 {
	x = Mod2Pi(x)
	f, p := 0.0, x
	for i := 0; !IsZero(p); i++ {
		f += p
		p *= -x * x / float64((2*i+2)*(2*i+3))
	}
	return f
}
