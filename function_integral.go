package cal

const NsubInt = 1000

func (fn Function) Integral() float64 {
	xarr := Linspace(fn.Interval[0], fn.Interval[1], NsubInt)
	var sum float64
	for i := 0; i < NsubInt; i++ {
		a, b := xarr[i], xarr[i+1]
		sum += (fn.Formula(a) + 4*fn.Formula((a+b)/2) + fn.Formula(b))
	}
	dx := (fn.Interval[1] - fn.Interval[0]) / NsubInt
	return sum * dx / 6
}
