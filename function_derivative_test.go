package cal_test

import (
	"fmt"
	"math"
	"recal/cal"
	"testing"
)

func TestFunction_Derivative_x2(*testing.T) {
	x2 := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  func(x float64) float64 { return x * x },
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-.1, 1.1}
	xy.Draw(x2.Graph())
	xy.Line(cal.Point{-1, 0}, cal.Point{1, 0})
	xy.Save()
}

func TestFunction_Derivative_inv(*testing.T) {
	f := cal.Function{
		Interval: [2]float64{-1, 1},
		Formula:  func(x float64) float64 { return x*x + 1 },
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.2, 2.2}
	xy.Ylim = [2]float64{-1.2, 2.2}
	grp := f.Graph()
	xy.Draw(grp)
	for _, l := range grp.Lines {
		xy.Line(cal.Point{l.From.Y, l.From.X},
			cal.Point{l.To.Y, l.To.X})
	}
	g := cal.Function{
		Interval: [2]float64{-1, 2},
		Formula:  func(x float64) float64 { return x },
	}
	xy.Draw(g.Graph())
	xy.Point(1, 0)
	xy.Save()
}

func TestDerivative_mvt(*testing.T) {
	g := cal.Function{
		Interval: [2]float64{-1.03, 1.22},
		Formula: func(x float64) float64 {
			return math.Pow(x, 3) - x + 1.5
		},
	}
	dg := (g.Formula(1.2) - g.Formula(-1)) / 2.2
	c := -math.Sqrt((1 + dg) / 3)

	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.2, 1.5}
	xy.Ylim = [2]float64{1, 2.3}
	f := cal.Function{
		Interval: [2]float64{c - .2, c + .2},
		Formula: func(x float64) float64 {
			return (3*c*c-1)*(x-c) + g.Formula(c)
		},
	}
	xy.Draw(f.Graph())
	c = math.Sqrt((1 + dg) / 3)
	fmt.Println(c)
	f = cal.Function{
		Interval: [2]float64{c - .2, c + .2},
		Formula: func(x float64) float64 {
			return (3*c*c-1)*(x-c) + g.Formula(c)
		},
	}
	xy.Draw(f.Graph())
	xy.Draw(g.Graph())
	xy.Line(cal.Point{-1, g.Formula(-1)},
		cal.Point{1.2, g.Formula(1.2)})
	xy.Save()
}

func TestFunction_Derivative_cauchy(*testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	intv := [2]float64{-cal.Pi/4 - .1, 3*cal.Pi/4 + .1}
	f := cal.Curve{
		Interval: intv,
		X:        cal.Cos,
		Y:        cal.Sin,
	}
	g := cal.Function{
		Interval: [2]float64{f.X(-cal.Pi/4) - .2, f.X(-cal.Pi/4) + .2},
		Formula: func(x float64) float64 {
			return -(x - f.X(cal.Pi/4)) + f.Y(cal.Pi/4)
		},
	}
	xy.Draw(f.Graph())
	xy.Draw(g.Graph())
	xy.Point(f.X(cal.Pi/4), f.Y(cal.Pi/4))
	xy.Line(cal.Point{f.X(-cal.Pi / 4), f.Y(-cal.Pi / 4)},
		cal.Point{f.X(3 * cal.Pi / 4), f.Y(3 * cal.Pi / 4)})
	xy.Save()
}

func TestFunction_Derivative_sinxx(*testing.T) {
	f := cal.Function{
		Interval: [2]float64{-10, 10},
		Formula: func(x float64) float64 {
			if cal.IsZero(x) {
				return 1
			}
			return cal.Sin(x) / x
		},
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-10.1, 10.1}
	xy.Ylim = [2]float64{-.3, 1.1}
	xy.Draw(f.Graph())
	xy.Save()
}
