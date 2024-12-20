package cal

func (fn Function) Differential(c float64) float64 {
	var h float64 = Eps
	l := (fn.Formula(c+h) - fn.Formula(c)) / h
	r := (fn.Formula(c) - fn.Formula(c-h)) / h
	for !IsZero(l - r) {
		h /= 2
		l = (fn.Formula(c+h) - fn.Formula(c)) / h
		r = (fn.Formula(c) - fn.Formula(c-h)) / h
	}
	return (l + r) / 2
}
