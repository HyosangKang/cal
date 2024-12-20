package cal_test

import (
	"recal/cal"
	"testing"
)

var discFn = cal.Function{
	Interval: [2]float64{-1., 1.},
	Formula: func(x float64) float64 {
		if x < 0 {
			return x*x - 1
		} else if x > 0 {
			return x*x + 1
		} else {
			return 0
		}
	}}

func TestFunction_GraphContinuous(t *testing.T) {
	g := discFn.GraphContinuous(.1)
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 2.1}
	xy.Draw(g)
	xy.Save()
}

func TestFunction(t *testing.T) {
	a, b := discFn.Interval[0], discFn.Interval[1]
	var n int = 9
	xarr := Linspace(a, b, n)
	var g cal.Graph
	for i := 0; i < n; i++ {
		x0, x1 := xarr[i], xarr[i+1]
		y0, y1 := discFn.Formula(x0), discFn.Formula(x1)
		g.Lines = append(g.Lines, cal.Line{
			From: cal.Point{X: x0, Y: y0},
			To:   cal.Point{X: x1, Y: y1},
		})
		g.Points = append(g.Points, cal.Point{X: x0, Y: y0})
		if i == n-1 {
			g.Points = append(g.Points, cal.Point{X: x1, Y: y1})
		}
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 2.1}
	xy.Draw(g)
	xy.Save()
}

func TestFunction_SinxX(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-10.5, 10.5}
	xy.Ylim = [2]float64{-.3, 1.1}
	xy.Draw(cal.SinxX.Graph())
	xy.Save()
}

func TestFunction_Exotic(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-2.1, 2.1}
	xy.Ylim = [2]float64{-.1, 1.1}
	xy.Draw(cal.Exotic.Graph())
	xy.Save()
}

func TestFunction_X2Sin(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(cal.X2Sin.Graph())
	xy.Save()
}

func TestFunction_X3Sin(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(cal.X3Sin.Graph())
	xy.Save()
}

// var itv = [2]float64{.1, 3.9}

// var (
// 	exp = cal.Function{
// 		Interval: itv,
// 		Formula:  math.Exp,
// 	}
// 	log = cal.Function{
// 		Interval: itv,
// 		Formula:  math.Log,
// 	}
// 	cube = cal.Function{
// 		Interval: itv,
// 		Formula:  func(x float64) float64 { return x * x * x },
// 	}
// )

// var (
// 	zero = cal.Function{
// 		Interval: [2]float64{-2, 2},
// 		Formula:  func(x float64) float64 { return 0 },
// 	}
// )
