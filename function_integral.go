package cal

const NsubInt = 50000

func (fn Function) Integral() float64 {
	a, b := fn.Interval[0], fn.Interval[1]
	xarr := Linspace(a, b, NsubInt)
	var sum float64
	f := fn.Formula
	for i := 0; i < NsubInt; i++ {
		a, b := xarr[i], xarr[i+1]
		sum += (f(a) + 4*f((a+b)/2) + f(b))
	}
	dx := (b - a) / NsubInt
	return sum * dx / 6
}
