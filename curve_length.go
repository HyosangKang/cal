package cal

func (c Curve) Length() float64 {
	x := Function{
		Interval: c.Interval,
		Formula:  c.X,
	}
	y := Function{
		Interval: c.Interval,
		Formula:  c.Y,
	}
	f := Function{
		Interval: c.Interval,
		Formula: func(t float64) float64 {
			xp := x.Differential(t)
			yp := y.Differential(t)
			return Sqrt(xp*xp + yp*yp)
		},
	}
	return f.Integral()
}
