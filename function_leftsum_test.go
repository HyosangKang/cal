package cal_test

import (
	"recal/cal"
	"testing"
)

func TestFunction_LeftSum(*testing.T) {
	var n int = 5
	var a, b float64 = 0, 1
	f := cal.Function{
		Interval: [2]float64{a, b},
		Formula: func(x float64) float64 {
			return x * x
		},
	}
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-.1, 1.1}
	xy.Ylim = [2]float64{-.1, 1.1}
	xy.Draw(f.Graph())
	// draw x-axis
	xy.Line(cal.Point{X: a, Y: 0}, cal.Point{X: b, Y: 0})
	// split unequal intervals
	xarr := []float64{a}
	var d float64 = b - a
	for i := 0; i < n; i++ {
		d /= 2
		xarr = append(xarr, xarr[len(xarr)-1]+d)
	}
	for i := 1; i <= n; i++ {
		x := xarr[i]
		y := f.Formula(x)
		// draw vertical line
		xy.Line(cal.Point{X: x, Y: 0}, cal.Point{X: x, Y: y})
		// draw horizontal line
		xy.Line(cal.Point{X: xarr[i-1], Y: y}, cal.Point{X: x, Y: y})
		// draw left vertical line
		xy.Line(cal.Point{X: xarr[i-1], Y: 0}, cal.Point{X: xarr[i-1], Y: y})
	}
	xy.Save()
}
