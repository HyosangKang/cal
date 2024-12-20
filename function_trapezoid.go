package cal

func (fn Function) Trapezoid(sub int) float64 {
	xarr := Linspace(fn.Interval[0], fn.Interval[1], sub)
	var sum float64
	for i := 0; i <= sub; i++ {
		sum += fn.Formula(xarr[i])
		if 0 < i && i < sub {
			sum += fn.Formula(xarr[i])
		}
	}
	dx := (fn.Interval[1] - fn.Interval[0]) / float64(sub)
	return sum * dx / 2
}
