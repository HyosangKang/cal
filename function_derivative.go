package cal

func (fn Function) Derivative() Function {
	return Function{
		Interval: fn.Interval,
		Formula:  fn.Differential,
	}
}
