package cal

func (fn Function) LeftSum(sub int) float64 {
	a, b := fn.Interval[0], fn.Interval[1]
	xarr := Linspace(a, b, sub)
	var sum float64
	for i := 0; i < sub; i++ {
		sum += fn.Formula(xarr[i])
	}
	dx := (b - a) / float64(sub)
	return sum * dx
}
