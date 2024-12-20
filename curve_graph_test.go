package cal_test

import (
	"recal/cal"
	"testing"
)

func TestCurve_Graph_cycloid(t *testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 6.9}
	xy.Ylim = [2]float64{-1.1, 6.9}
	xy.Draw(cal.Cycloid.Graph())
	t0, t1, t2 := 0., 2.8, 5.8
	xy.Draw(cal.CycloidCircle(t0).Graph())
	xy.Draw(cal.CycloidCircle(t1).Graph())
	xy.Draw(cal.CycloidCircle(t2).Graph())
	xy.Point(cal.Cycloid.X(t0), cal.Cycloid.Y(t0))
	xy.Point(cal.Cycloid.X(t1), cal.Cycloid.Y(t1))
	xy.Point(cal.Cycloid.X(t2), cal.Cycloid.Y(t2))
	xy.Save()
}

func TestCurve_Graph(*testing.T) {
	end := 3.0
	c := cal.Curve{
		Interval: [2]float64{0, end},
		X: func(t float64) float64 {
			return t
		},
		Y: func(t float64) float64 {
			return cal.Exp(-t) + t
		},
	}
	graph := c.Graph()
	graph.Lines = append(graph.Lines, cal.Line{
		From: cal.Point{X: 0, Y: 0},
		To:   cal.Point{X: 3, Y: 3}})
	xy := cal.NewXY()
	xy.Xlim = [2]float64{0. - .1, 3.2}
	xy.Ylim = [2]float64{0. - .1, 3.2}
	xy.Draw(graph)
	xy.Save()
}

func TestCurve_Graph_circle(*testing.T) {
	xy := cal.NewXY()
	xy.Xlim = [2]float64{-1.1, 1.1}
	xy.Ylim = [2]float64{-1.1, 1.1}
	xy.Draw(cal.UnitCircle.Graph())
	var n int = 25
	for i := 0; i <= n; i++ {
		t := -1 + 2*float64(i)/float64(n)
		xy.Line(cal.Point{X: t, Y: -1}, cal.Point{X: t, Y: 1})
		xy.Line(cal.Point{X: -1, Y: t}, cal.Point{X: 1, Y: t})
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := -1. + 2*float64(i)/float64(n)
			y := -1. + 2*float64(j)/float64(n)
			if x*x+y*y <= 1 {
				xy.Point(x, y)
			}
		}
	}
	xy.Save()
}
