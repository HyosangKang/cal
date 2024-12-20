package cal_test

import (
	"math"
	"recal/cal"
	"testing"
)

func TestFunction_Differential_rate(*testing.T) {
	sin := cal.Function{
		Interval: [2]float64{0, cal.Pi / 2},
		Formula:  math.Sin,
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-.1, cal.Pi/2 + .1}
	xy.Ylim = [2]float64{-.1, 1.1}
	xy.Draw(sin.Graph())
	xy.Line(cal.Point{0, sin.Formula(0)}, cal.Point{1, sin.Formula(1)})
	xy.Save()
}

func TestFunction_Differential_abs(*testing.T) {
	abs := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  math.Abs,
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-.6, 1.6}
	xy.Draw(abs.Graph())
	xy.Line(cal.Point{0, abs.Formula(0)}, cal.Point{1, abs.Formula(1)})
	xy.Save()
}

func TestFunction_Differential_x2sin(*testing.T) {
	x2sin := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			if cal.IsZero(x) {
				return 0
			}
			return x * x * cal.Sin(1/x)
		},
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(x2sin.Graph())
	g := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			return x * x
		},
	}
	xy.Draw(g.Graph())
	g.Formula = func(x float64) float64 {
		return -x * x
	}
	xy.Draw(g.Graph())
	xy.Line(cal.Point{-1, 0}, cal.Point{1, 0})
	xy.Save()
}

func TestFunction_Differential_xsin(*testing.T) {
	x2sin := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			if cal.IsZero(x) {
				return 0
			}
			return x * cal.Sin(1/x)
		},
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(x2sin.Graph())
	g := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula: func(x float64) float64 {
			return x
		},
	}
	xy.Draw(g.Graph())
	g.Formula = func(x float64) float64 {
		return -x
	}
	xy.Draw(g.Graph())
	xy.Line(cal.Point{-1, 0}, cal.Point{1, 0})
	xy.Save()
}
