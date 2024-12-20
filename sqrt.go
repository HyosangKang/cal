package cal

func Sqrt(x float64) float64 {
	if x < 0 {
		panic("Invalid input")
	}
	var n int
	for x >= 4 {
		x /= 4
		n++
	}
	var y, p float64 = 0, 1
	for !IsZero(p) {
		if (y+p)*(y+p) < x {
			y += p
		}
		p /= 2
	}
	for n > 0 {
		y *= 2
		n--
	}
	return y
}
